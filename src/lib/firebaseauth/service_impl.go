package firebaseauth

import (
	"context"
	"fmt"

	"github.com/abyssparanoia/rapid-go-woker/src/lib/log"
)

type service struct {
}

// Authentication ... 認証を行う
func (s *service) Authentication(ctx context.Context, ah string) (string, *Claims, error) {
	var userID string
	claims := &Claims{}

	c, err := getAuthClient(ctx)
	if err != nil {
		log.Warningm(ctx, "getAuthClient", err)
		return userID, claims, err
	}

	token := getTokenByAuthHeader(ah)
	if token == "" {
		err := log.Warninge(ctx, "token empty error")
		return userID, claims, err
	}

	t, err := c.VerifyIDToken(ctx, token)
	if err != nil {
		msg := fmt.Sprintf("c.VerifyIDToken: %s", token)
		log.Warningm(ctx, msg, err)
		return userID, claims, err
	}

	userID = t.UID
	claims.SetMap(t.Claims)

	return userID, claims, nil
}

// SetCustomClaims ... カスタムClaimsを設定
func (s *service) SetCustomClaims(ctx context.Context, userID string, claims *Claims) error {
	c, err := getAuthClient(ctx)
	if err != nil {
		log.Errorm(ctx, "getAuthClient", err)
		return err
	}

	err = c.SetCustomUserClaims(ctx, userID, claims.ToMap())
	if err != nil {
		log.Errorm(ctx, "c.SetCustomUserClaims", err)
		return err
	}
	return nil
}

// NewService ... Serviceを作成する
func NewService() Service {
	return &service{}
}
