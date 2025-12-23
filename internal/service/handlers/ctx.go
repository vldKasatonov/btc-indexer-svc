package handlers

import (
	"context"
	"net/http"

	"github.com/vldKasatonov/btc-indexer-svc/internal/data"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	usersQCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxUsersQ(entry data.UsersQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, usersQCtxKey, entry)
	}
}

func UsersQ(r *http.Request) data.UsersQ {
	return r.Context().Value(usersQCtxKey).(data.UsersQ).New()
}
