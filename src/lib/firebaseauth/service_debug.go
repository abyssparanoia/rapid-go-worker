package firebaseauth

import (
	"context"
	"fmt"

	"github.com/abyssparanoia/rapid-go-worker/src/lib/log"
)

type serviceDebug struct {
}

// Authentication ... 認証を行う
func (s *serviceDebug) Authentication(ctx context.Context, ah string) (string, *Claims, error) {
	var userID string
	claims := &Claims{}

	// ユーザーを取得できたらデバッグリクエストと判定する
	if user := getUserByAuthHeader(ah); user != "" {
		claims = NewDummyClaims()
		return user, claims, nil
	}

	// 通常の認証を行う
	token := getTokenByAuthHeader(ah)
	if token == "" {
		err := log.Warninge(ctx, "token empty error")
		return userID, claims, err
	}

	c, err := getAuthClient(ctx)
	if err != nil {
		log.Warningm(ctx, "getAuthClient", err)
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
func (s *serviceDebug) SetCustomClaims(ctx context.Context, userID string, claims *Claims) error {
	c, err := getAuthClient(ctx)
	if err != nil {
		log.Errorm(ctx, "getAuthClient", err)
		return err
	}

	// AuthorizationHeaderからUserが取得できたらデバッグリクエストと判定する
	ah := getAuthHeader(ctx)
	if getUserByAuthHeader(ah) == "" {
		err = c.SetCustomUserClaims(ctx, userID, claims.ToMap())
		if err != nil {
			log.Errorm(ctx, "c.SetCustomUserClaims", err)
			return err
		}
	}
	return nil
}

// NewDebugService ... DebugServiceを作成する
func NewDebugService() Service {
	return &serviceDebug{}
}
