package auth

import "golang.org/x/net/context"

type Auth interface {
	Authorize(ctx context.Context) (*model.User, error)
	Authenticate(payload map[string]interface{}, userID string, uri string) bool
}
