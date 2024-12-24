//go:build functional

package garmin

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"testing"

	"github.com/dghubble/oauth1"

	"github.com/jylitalo/go-garmin/internal/rt"
)

func TestFunctional(t *testing.T) {
	tok := testToken(t)
	var claims garminClaims
	err := unmarshalClaims(&claims, tok.AccessToken)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", claims)
}

func TestFunctional_Login(t *testing.T) {
	t.Skip()
	slog.SetLogLoggerLevel(slog.LevelDebug)
	cacher.DelOAuth1Token()
	cacher.DelAccessToken()
	client := NewClient(WithCacher(cacher))
	client.prependTransport(new(rt.Debugger))
	var (
		err         error
		accessToken *AccessToken
	)
	_, accessToken, err = login(
		client,
		os.Getenv("GOGARMIN_TEST_EMAIL"),
		os.Getenv("GOGARMIN_TEST_PASSWORD"),
	)
	if errors.Is(err, ErrTokenCacheNotFound) {
		t.Errorf("token cache failed: %v", err)
	} else if errors.Is(err, ErrTokenCacheExpired) {
		t.Errorf("token cache stored an expired token")
	} else if err != nil {
		t.Fatal("failed to login:", err)
	}
	client.prependTransport(newAccessTokenInjector(accessToken, nil))
	api := NewAPI(client)
	_ = api
}

func TestFunctional_Refresher(t *testing.T) {
	t.Skip()
	cacher.DelAccessToken()
	cacher.DelOAuth1Token()
	client := NewClient(WithCacher(cacher))
	client.prependTransport(new(rt.Debugger))
	var (
		err         error
		accessToken *AccessToken
		basicToken  *oauth1.Token
	)
	basicToken, accessToken, err = login(
		client,
		os.Getenv("GOGARMIN_TEST_EMAIL"),
		os.Getenv("GOGARMIN_TEST_PASSWORD"),
	)
	if errors.Is(err, ErrTokenCacheNotFound) {
		t.Errorf("token cache failed: %v", err)
	} else if errors.Is(err, ErrTokenCacheExpired) {
		t.Errorf("token cache stored an expired token")
	} else if err != nil {
		t.Fatal("failed to login:", err)
	}
	client.prependTransport(newAccessTokenInjector(accessToken, nil))
	refresher := oauth1TokenRefresher{
		token:  basicToken,
		client: client,
	}
	newToken, err := refresher.Refresh(accessToken)
	if err != nil {
		t.Fatal(err)
	}
	if err = cacher.SaveAccessToken(newToken); err != nil {
		t.Fatal(err)
	}
	if accessToken.AccessToken == newToken.AccessToken {
		t.Error("access tokens should be different")
	}
	if accessToken.ExpiresAt().After(newToken.ExpiresAt()) {
		t.Error("old token should not expire after new token")
	}
	var claims, newClaims garminClaims
	err = unmarshalClaims(&claims, accessToken.AccessToken)
	if err != nil {
		t.Fatal(err)
	}
	err = unmarshalClaims(&newClaims, newToken.AccessToken)
	if err != nil {
		t.Fatal(err)
	}
	if claims.Exp > newClaims.Exp {
		t.Error("old claims exp should not be larger than the new claims exp")
	}
	if claims.GarminGUID != newClaims.GarminGUID {
		t.Error("the GUID in both sets of claims should be the same")
	}
	if claims.Iss != newClaims.Iss {
		t.Error("the iss in both sets of claims should be the same")
	}
}
