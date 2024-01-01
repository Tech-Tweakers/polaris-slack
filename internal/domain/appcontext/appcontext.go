package appcontext

import (
	"context"
	"slack-messages-api/internal/infrastructure/logger/logwrapper"
)

type Context interface {
	SetLogger(logger logwrapper.Logger)
	Logger() logwrapper.Logger
}

type appContext struct {
	logger logwrapper.Logger
	ctx    context.Context
}

func New(ctx context.Context) Context {
	return &appContext{
		ctx: ctx,
	}
}

func NewBackground() Context {
	ctx := context.Background()
	return &appContext{
		ctx: ctx,
	}
}

func (appContext *appContext) SetLogger(logger logwrapper.Logger) {
	appContext.logger = logger
}

func (appContext *appContext) Logger() logwrapper.Logger {
	return appContext.logger
}
