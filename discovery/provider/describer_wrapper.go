package provider

import (
	model "github.com/opengovern/og-describer-cloudflare/discovery/pkg/models"
	"github.com/opengovern/og-util/pkg/describe/enums"
	"golang.org/x/net/context"
	"github.com/cloudflare/cloudflare-go"
	"golang.org/x/time/rate"
	"net/http"
	"time"
	"errors"

)


type CloudFlareAPIHandler struct {
	Conn         *cloudflare.API
	AccountID    string
	RateLimiter  *rate.Limiter
	Semaphore    chan struct{}
	MaxRetries   int
	RetryBackoff time.Duration
}

func NewCloudFlareAPIHandler(client *cloudflare.API, accountID string, rateLimit rate.Limit, burst int, maxConcurrency int, maxRetries int, retryBackoff time.Duration) *CloudFlareAPIHandler {
	return &CloudFlareAPIHandler{
		Conn:         client,
		AccountID:    accountID,
		RateLimiter:  rate.NewLimiter(rateLimit, burst),
		Semaphore:    make(chan struct{}, maxConcurrency),
		MaxRetries:   maxRetries,
		RetryBackoff: retryBackoff,
	}
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



var (
	triggerTypeKey string = "trigger_type"
)
func WithTriggerType(ctx context.Context, tt enums.DescribeTriggerType) context.Context {
	return context.WithValue(ctx, triggerTypeKey, tt)
}


// DescribeListByCloudFlare A wrapper to pass cloudflare authorization to describer functions
func DescribeListByCloudFlare(describe func(context.Context, *CloudFlareAPIHandler, *model.StreamSender) ([]model.Resource, error)) model.ResourceDescriber {
	return func(ctx context.Context, cfg model.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
		ctx = WithTriggerType(ctx, triggerType)

		// Create cloudflare client using token or (email, api key)
		var conn *cloudflare.API
		var err error
		// Check for the token
		if cfg.Token != "" {
			conn, err = cloudflare.NewWithAPIToken(cfg.Token)
			if err != nil {
				return nil, err
			}
		}

		cloudflareAPIHandler := NewCloudFlareAPIHandler(conn, cfg.AccountID, rate.Every(time.Second/4), 1, 10, 5, 5*time.Minute)

		// Get values from describer
		var values []model.Resource
		result, err := describe(ctx, cloudflareAPIHandler, stream)
		if err != nil {
			return nil, err
		}
		values = append(values, result...)
		return values, nil
	}
}

// DescribeSingleByCloudFlare A wrapper to pass cloudflare authorization to describer functions
func DescribeSingleByCloudFlare(describe func(context.Context, *CloudFlareAPIHandler, string) (*model.Resource, error)) model.SingleResourceDescriber {
	return func(ctx context.Context, cfg model.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, resourceID string,stram *model.StreamSender) (*model.Resource, error) {
		ctx = WithTriggerType(ctx, triggerType)

		// Create cloudflare client using token or (email, api key)
		var conn *cloudflare.API
		var err error
		// Check for the token
		if cfg.Token != "" {
			conn, err = cloudflare.NewWithAPIToken(cfg.Token)
			if err != nil {
				return nil, err
			}
		}

		cloudflareAPIHandler := NewCloudFlareAPIHandler(conn, cfg.AccountID, rate.Every(time.Second/4), 1, 10, 5, 5*time.Minute)

		// Get value from describer
		value, err := describe(ctx, cloudflareAPIHandler, resourceID)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
}
