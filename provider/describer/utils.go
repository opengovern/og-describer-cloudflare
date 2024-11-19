package describer

import (
	"context"
	"errors"
	"github.com/cloudflare/cloudflare-go"
	"time"
)

const (
	maxRetries = 3
	backoff    = 2 * time.Second
	perPage    = 50
	page       = 1
)

func getAllAccounts(ctx context.Context, conn *cloudflare.API) ([]cloudflare.Account, error) {
	pageOpts := cloudflare.PaginationOptions{
		PerPage: perPage,
		Page:    page,
	}
	accounts, _, err := conn.Accounts(ctx, pageOpts)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func retry(ctx context.Context, operation func() (interface{}, error), shouldRetry func(error) bool) (interface{}, error) {
	var result interface{}
	var err error

	for attempt := 0; attempt < maxRetries; attempt++ {
		// Call the operation
		result, err = operation()
		if err == nil {
			return result, nil
		}

		// Check if the error is retryable
		if !shouldRetry(err) {
			return nil, err
		}

		// Wait before retrying
		select {
		case <-time.After(backoff):
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}

	return nil, err
}

func shouldRetryError(err error) bool {
	var cloudflareErr *cloudflare.APIRequestError
	if errors.As(err, &cloudflareErr) {
		return cloudflareErr.ClientRateLimited()
	}
	return false
}
