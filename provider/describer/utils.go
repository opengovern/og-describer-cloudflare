package describer

import (
	"context"
	"errors"
	"github.com/cloudflare/cloudflare-go"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

const (
	perPage = 100
	page    = 1
)

type CloudFlareAPIHandler struct {
	Conn         *cloudflare.API
	RateLimiter  *rate.Limiter
	Semaphore    chan struct{}
	MaxRetries   int
	RetryBackoff time.Duration
}

func NewCloudFlareAPIHandler(client *cloudflare.API, rateLimit rate.Limit, burst int, maxConcurrency int, maxRetries int, retryBackoff time.Duration) *CloudFlareAPIHandler {
	return &CloudFlareAPIHandler{
		Conn:         client,
		RateLimiter:  rate.NewLimiter(rateLimit, burst),
		Semaphore:    make(chan struct{}, maxConcurrency),
		MaxRetries:   maxRetries,
		RetryBackoff: retryBackoff,
	}
}

func getAccount(ctx context.Context, handler *CloudFlareAPIHandler) (*cloudflare.Account, error) {
	var accounts []cloudflare.Account
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		pageOpts := cloudflare.PaginationOptions{
			PerPage: perPage,
			Page:    page,
		}
		accounts, _, e = handler.Conn.Accounts(ctx, pageOpts)
		if e != nil {
			var httpErr *cloudflare.APIRequestError
			if errors.As(e, &httpErr) {
				statusCode = &httpErr.StatusCode
			}
		}
		return statusCode, e
	}
	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return nil, err
	}
	return &accounts[0], nil
}

func getApplications(ctx context.Context, handler *CloudFlareAPIHandler, accountID string) ([]cloudflare.AccessApplication, error) {
	var accessApps []cloudflare.AccessApplication
	var pageAccessApps []cloudflare.AccessApplication
	var pageData cloudflare.ResultInfo
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		opts := cloudflare.PaginationOptions{
			PerPage: perPage,
			Page:    page,
		}
		for {
			pageAccessApps, pageData, e = handler.Conn.AccessApplications(ctx, accountID, opts)
			if e != nil {
				var httpErr *cloudflare.APIRequestError
				if errors.As(e, &httpErr) {
					statusCode = &httpErr.StatusCode
				}
			}
			accessApps = append(accessApps, pageAccessApps...)
			if pageData.Page >= pageData.TotalPages {
				break
			}
			opts.Page = opts.Page + 1
		}
		return statusCode, e
	}
	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return nil, err
	}
	return accessApps, nil
}

func getZones(ctx context.Context, handler *CloudFlareAPIHandler) ([]cloudflare.Zone, error) {
	var zones []cloudflare.Zone
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		zones, e = handler.Conn.ListZones(ctx)
		if e != nil {
			var httpErr *cloudflare.APIRequestError
			if errors.As(e, &httpErr) {
				statusCode = &httpErr.StatusCode
			}
		}
		return statusCode, e
	}
	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return nil, err
	}
	return zones, nil
}

// DoRequest executes the Cloudflare API request with rate limiting, retries, and concurrency control.
func (h *CloudFlareAPIHandler) DoRequest(ctx context.Context, requestFunc func() (*int, error)) error {
	h.Semaphore <- struct{}{}
	defer func() { <-h.Semaphore }()
	var statusCode *int
	var err error
	for attempt := 0; attempt <= h.MaxRetries; attempt++ {
		// Wait based on rate limiter
		if err = h.RateLimiter.Wait(ctx); err != nil {
			return err
		}
		// Execute the request function
		statusCode, err = requestFunc()
		if err == nil {
			return nil
		}
		// Handle rate limit errors
		if statusCode != nil && *statusCode == http.StatusTooManyRequests {
			backoff := h.RetryBackoff * (1 << attempt)
			time.Sleep(backoff)
			continue
		}
		// Handle temporary network errors
		if isTemporary(err) {
			backoff := h.RetryBackoff * (1 << attempt)
			time.Sleep(backoff)
			continue
		}
		break
	}
	return err
}

// isTemporary checks if an error is temporary.
func isTemporary(err error) bool {
	if err == nil {
		return false
	}
	var netErr interface{ Temporary() bool }
	if errors.As(err, &netErr) {
		return netErr.Temporary()
	}
	return false
}
