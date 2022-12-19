package handlers

import (
	"context"
	"net/http"

	"gitlab.com/distributed_lab/acs/identity-svc/internal/data"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	usersCtxKey
	positionsCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func UsersQ(r *http.Request) data.UsersQ {
	return r.Context().Value(usersCtxKey).(data.UsersQ).New()
}

func CtxUsersQ(entry data.UsersQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, usersCtxKey, entry)
	}
}

func CtxPositions(entry []string) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, positionsCtxKey, entry)
	}
}

func Positions(r *http.Request) []string {
	return r.Context().Value(positionsCtxKey).([]string)
}
