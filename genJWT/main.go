package main

import (
	"context"
	"crypto"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"sync"
	"time"
)

const defaultKey = "CG24SDVP8OHPK395GB5G"
const configKey = "XnEsT01S@"

var ErrInvalidToken = errors.New("invalid token")
var (
	ErrInvalidKey      = errors.New("key is invalid")
	ErrInvalidKeyType  = errors.New("key is of invalid type")
	ErrHashUnavailable = errors.New("the requested hash function is unavailable")

	ErrSignatureInvalid = errors.New("signature is invalid")
)

func init() {
	crypto.RegisterHash(crypto.SHA512, sha512.New)
}

type SigningMethodHMAC struct {
	Name string
	Hash crypto.Hash
}

func (m *SigningMethodHMAC) Alg() string {
	return m.Name
}

// Verify the signature of HSXXX tokens.  Returns nil if the signature is valid.
func (m *SigningMethodHMAC) Verify(signingString, signature string, key interface{}) error {
	// Verify the key is the right type
	keyBytes, ok := key.([]byte)
	if !ok {
		return ErrInvalidKeyType
	}

	// Decode signature, for comparison
	sig, err := DecodeSegment(signature)
	if err != nil {
		return err
	}

	// Can we use the specified hashing method?
	if !m.Hash.Available() {
		return ErrHashUnavailable
	}

	// This signing method is symmetric, so we validate the signature
	// by reproducing the signature from the signing string and key, then
	// comparing that against the provided signature.
	hasher := hmac.New(m.Hash.New, keyBytes)
	hasher.Write([]byte(signingString))
	if !hmac.Equal(sig, hasher.Sum(nil)) {
		return ErrSignatureInvalid
	}

	// No validation errors.  Signature is good.
	return nil
}

// Decode JWT specific base64url encoding with padding stripped
func DecodeSegment(seg string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(seg)
}

// Implements the Sign method from SigningMethod for this signing method.
// Key must be []byte
func (m *SigningMethodHMAC) Sign(signingString string, key interface{}) (string, error) {
	if keyBytes, ok := key.([]byte); ok {
		if !m.Hash.Available() {
			return "", ErrHashUnavailable
		}

		hasher := hmac.New(m.Hash.New, keyBytes)
		hasher.Write([]byte(signingString))

		return EncodeSegment(hasher.Sum(nil)), nil
	}

	return "", ErrInvalidKeyType
}

var signingMethods = map[string]func() SigningMethod{}
var signingMethodLock = new(sync.RWMutex)

// Register the "alg" name and a factory function for signing method.
// This is typically done during init() in the method's implementation
func RegisterSigningMethod(alg string, f func() SigningMethod) {
	signingMethodLock.Lock()
	defer signingMethodLock.Unlock()

	signingMethods[alg] = f
}

func main() {

	res, err := genJwt("root")
	if err != nil {
		panic(err)
	}

	log.Printf("token   %s  \n", res)
}

func genJwt(subject string) (res []byte, err error) {
	now := time.Now()
	var expiresAt int64
	expiresAt = now.Add(time.Hour * 999999).Unix()

	// HS512
	SigningMethodHS512 := &SigningMethodHMAC{"HS512", crypto.SHA512}
	RegisterSigningMethod(SigningMethodHS512.Alg(), func() SigningMethod {
		return SigningMethodHS512
	})
	token := NewWithClaims(SigningMethodHS512, &StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: expiresAt,
		NotBefore: now.Unix(),
		Subject:   subject,
	})

	tokenStr, err := token.SignedString([]byte(configKey))
	if err != nil {
		return nil, err
	}

	tokenInfo := &tokenInfo{
		ExpiresAt:   expiresAt,
		TokenType:   "Bearer",
		AccessToken: tokenStr,
	}
	return tokenInfo.EncodeToJSON()
}

type Option func(*options)

// func New(store Storer, opts ...Option) Auther {
// 	o := options{
// 		tokenType:        "Bearer",
// 		expired:          7200,
// 		mobileExpiredDay: 7200,
// 		signingMethod:    jwt.SigningMethodHS512,
// 		signingKey:       []byte(defaultKey),
// 	}

// 	for _, opt := range opts {
// 		opt(&o)
// 	}

// 	o.keyFuncs = append(o.keyFuncs, func(t *Token) (interface{}, error) {
// 		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, ErrInvalidToken
// 		}
// 		return o.signingKey, nil
// 	})

// 	if o.signingKey2 != nil {
// 		o.keyFuncs = append(o.keyFuncs, func(t *Token) (interface{}, error) {
// 			if _, ok := t.Method.(*SigningMethodHMAC); !ok {
// 				return nil, ErrInvalidToken
// 			}
// 			return o.signingKey2, nil
// 		})
// 	}

// 	return &JWTAuth{
// 		opts:  &o,
// 		store: store,
// 	}
// }

// type Auther interface {
// 	// GenerateToken Generate a JWT (JSON Web Token) with the provided subject.
// 	GenerateToken(ctx context.Context, subject string) (TokenInfo, error)
// 	// DestroyToken Invalidate a token by removing it from the token store.
// 	DestroyToken(ctx context.Context, accessToken string) error
// 	// ParseSubject Parse the subject (or user identifier) from a given access token.
// 	ParseSubject(ctx context.Context, accessToken string) (string, error)
// 	// Release any resources held by the JWTAuth instance.
// 	Release(ctx context.Context) error
// }

// type JWTAuth struct {
// 	opts  *options
// 	store Storer
// }

// func (a *JWTAuth) GenerateToken(ctx context.Context, subject string) (TokenInfo, error) {
// 	now := time.Now()
// 	var expiresAt int64

// 	token := NewWithClaims(a.opts.signingMethod, &StandardClaims{
// 		IssuedAt:  now.Unix(),
// 		ExpiresAt: expiresAt,
// 		NotBefore: now.Unix(),
// 		Subject:   subject,
// 	})

// 	tokenStr, err := token.SignedString(a.opts.signingKey)
// 	if err != nil {
// 		return nil, err
// 	}

// 	tokenInfo := &tokenInfo{
// 		ExpiresAt:   expiresAt,
// 		TokenType:   a.opts.tokenType,
// 		AccessToken: tokenStr,
// 	}
// 	return tokenInfo, nil
// }

// func (a *JWTAuth) parseToken(tokenStr string) (*StandardClaims, error) {
// 	var (
// 		token *Token
// 		err   error
// 	)

// 	for _, keyFunc := range a.opts.keyFuncs {
// 		token, err = jwt.ParseWithClaims(tokenStr, &StandardClaims{}, keyFunc)
// 		if err != nil || token == nil || !token.Valid {
// 			continue
// 		}
// 		break
// 	}

// 	if err != nil || token == nil || !token.Valid {
// 		return nil, ErrInvalidToken
// 	}

// 	return token.Claims.(*StandardClaims), nil
// }

// func (a *JWTAuth) callStore(fn func(Storer) error) error {
// 	if store := a.store; store != nil {
// 		return fn(store)
// 	}
// 	return nil
// }

// func (a *JWTAuth) DestroyToken(ctx context.Context, tokenStr string) error {
// 	claims, err := a.parseToken(tokenStr)
// 	if err != nil {
// 		return err
// 	}

// 	return a.callStore(func(store Storer) error {
// 		expired := time.Until(time.Unix(claims.ExpiresAt, 0))
// 		return store.Set(ctx, tokenStr, expired)
// 	})
// }

// func (a *JWTAuth) ParseSubject(ctx context.Context, tokenStr string) (string, error) {
// 	if tokenStr == "" {
// 		return "", ErrInvalidToken
// 	}

// 	claims, err := a.parseToken(tokenStr)
// 	if err != nil {
// 		return "", err
// 	}

// 	err = a.callStore(func(store Storer) error {
// 		if exists, err := store.Check(ctx, tokenStr); err != nil {
// 			return err
// 		} else if exists {
// 			return ErrInvalidToken
// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		return "", err
// 	}

// 	return claims.Subject, nil
// }

// func (a *JWTAuth) Release(ctx context.Context) error {
// 	return a.callStore(func(store Storer) error {
// 		return store.Close(ctx)
// 	})
// }

type options struct {
	signingMethod    SigningMethod
	signingKey       []byte
	signingKey2      []byte
	keyFuncs         []func(*Token) (interface{}, error)
	expired          int
	mobileExpiredDay int
	tokenType        string
}

// Storer is the interface that storage the token.
type Storer interface {
	Set(ctx context.Context, tokenStr string, expiration time.Duration) error
	Delete(ctx context.Context, tokenStr string) error
	Check(ctx context.Context, tokenStr string) (bool, error)
	Close(ctx context.Context) error
}

type tokenInfo struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresAt   int64  `json:"expires_at"`
}

func (t *tokenInfo) GetAccessToken() string {
	return t.AccessToken
}

func (t *tokenInfo) GetTokenType() string {
	return t.TokenType
}

func (t *tokenInfo) GetExpiresAt() int64 {
	return t.ExpiresAt
}

func (t *tokenInfo) EncodeToJSON() ([]byte, error) {
	return json.Marshal(t)
}
