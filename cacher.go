package garmin

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/dghubble/oauth1"
)

var (
	ErrTokenCacheNotFound = errors.New("could not find token in cache")
	ErrTokenCacheExpired  = errors.New("token saved in cache is expired")
)

func TokenCacheOk(cache TokenCacher) bool {
	_, err := cache.GetAccessToken()
	return err == nil
}

type OAuth1Token = oauth1.Token

type TokenCacher interface {
	SaveAccessToken(at *AccessToken) error
	GetAccessToken() (*AccessToken, error)
	DelAccessToken() error
	SaveOAuth1Token(token *OAuth1Token) error
	GetOAuth1Token() (*OAuth1Token, error)
	DelOAuth1Token() error
}

type InMemTokenCacher struct {
	at *AccessToken
	ot *oauth1.Token
}

func (imtc *InMemTokenCacher) GetAccessToken() (*AccessToken, error) {
	if imtc.at == nil {
		return nil, ErrTokenCacheNotFound
	}
	return imtc.at, nil
}

func (imtc *InMemTokenCacher) GetOAuth1Token() (*OAuth1Token, error) {
	if imtc.ot == nil {
		return nil, ErrTokenCacheNotFound
	}
	return imtc.ot, nil
}

func (imtc *InMemTokenCacher) SaveAccessToken(at *AccessToken) error {
	imtc.at = at
	return nil
}

func (imtc *InMemTokenCacher) SaveOAuth1Token(token *OAuth1Token) error {
	imtc.ot = token
	return nil
}

func (imtc *InMemTokenCacher) DelAccessToken() error {
	imtc.at = nil
	return nil
}

func (imtc *InMemTokenCacher) DelOAuth1Token() error {
	imtc.ot = nil
	return nil
}

func NewFileTokenCacher(p string) *FileTokenCacher {
	return &FileTokenCacher{Path: p}
}

type FileTokenCacher struct {
	Path string
	mem  InMemTokenCacher
}

func (ftc *FileTokenCacher) SaveAccessToken(at *AccessToken) error {
	if err := ftc.save("access_token.json", at); err != nil {
		return err
	}
	return ftc.mem.SaveAccessToken(at)
}

func (ftc *FileTokenCacher) GetAccessToken() (*AccessToken, error) {
	var at AccessToken
	if t, err := ftc.mem.GetAccessToken(); err == nil && t != nil {
		return t, nil
	}
	if err := ftc.get("access_token.json", &at); err != nil {
		return &at, err
	}
	return &at, ftc.mem.SaveAccessToken(&at)
}

func (ftc *FileTokenCacher) SaveOAuth1Token(token *OAuth1Token) error {
	if err := ftc.save("oauth1_token.json", token); err != nil {
		return err
	}
	return ftc.mem.SaveOAuth1Token(token)
}

func (ftc *FileTokenCacher) GetOAuth1Token() (*OAuth1Token, error) {
	var token oauth1.Token
	if t, err := ftc.mem.GetOAuth1Token(); err == nil && t != nil {
		return t, nil
	}
	if err := ftc.get("oauth1_token.json", &token); err != nil {
		return &token, err
	}
	return &token, ftc.mem.SaveOAuth1Token(&token)
}

func (ftc *FileTokenCacher) DelAccessToken() (err error) {
	if err = ftc.mem.DelAccessToken(); err != nil {
		return
	}
	return os.Remove(filepath.Join(ftc.Path, "oauth1_token.json"))
}

func (ftc *FileTokenCacher) DelOAuth1Token() (err error) {
	if err = ftc.mem.DelOAuth1Token(); err != nil {
		return
	}
	return os.Remove(filepath.Join(ftc.Path, "oauth1_token.json"))
}

func (ftc *FileTokenCacher) save(name string, token any) error {
	_, err := os.Stat(ftc.Path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(ftc.Path, 0755)
		if err != nil {
			return err
		}
	}
	f, err := os.OpenFile(
		filepath.Join(ftc.Path, name),
		os.O_TRUNC|os.O_CREATE|os.O_WRONLY,
		0600,
	)
	if err != nil {
		return err
	}
	defer f.Close()
	b, err := json.Marshal(token)
	if err != nil {
		return err
	}
	_, err = f.Write(b)
	return err
}

func (ftc *FileTokenCacher) get(name string, token any) error {
	f, err := os.OpenFile(
		filepath.Join(ftc.Path, name),
		os.O_RDONLY,
		0600,
	)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrTokenCacheNotFound
		}
		return err
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, token)
	return err
}

func getCachedPair(cacher TokenCacher) (*oauth1.Token, *AccessToken, error) {
	ot, err := cacher.GetOAuth1Token()
	if err != nil {
		return nil, nil, err
	}
	at, err := cacher.GetAccessToken()
	if err != nil {
		return ot, nil, err
	}
	return ot, at, nil
}
