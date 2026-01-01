package handlers

import (
	"context"
	"net/http"

	"github.com/vldKasatonov/btc-indexer-svc/internal/config"
	"github.com/vldKasatonov/btc-indexer-svc/internal/data"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	signerConfigKey
	usersQCtxKey
	userIdKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxSigner(entry *config.SignerConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, signerConfigKey, entry)
	}
}

func Signer(r *http.Request) *config.SignerConfig {
	return r.Context().Value(signerConfigKey).(*config.SignerConfig)
}

func CtxUsersQ(entry data.UsersQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, usersQCtxKey, entry)
	}
}

func UsersQ(r *http.Request) data.UsersQ {
	return r.Context().Value(usersQCtxKey).(data.UsersQ).New()
}

func CtxUserId(userId int64) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, userIdKey, userId)
	}
}

func UserId(r *http.Request) int64 {
	return r.Context().Value(userIdKey).(int64)
}
