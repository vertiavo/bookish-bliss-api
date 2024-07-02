package context

import (
	"context"
	"log"

	"github.com/vertiavo/bookish-bliss-api/internal/config"
)

type key int

const (
	configKey key = iota
	loggerKey
)

// WithConfig returns a new context with the given config.
func WithConfig(ctx context.Context, cfg *config.Config) context.Context {
	return context.WithValue(ctx, configKey, cfg)
}

// WithLogger returns a new context with the given logger.
func WithLogger(ctx context.Context, logger *log.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

// GetConfig returns the config from the context.
func GetConfig(ctx context.Context) *config.Config {
	return ctx.Value(configKey).(*config.Config)
}

// GetLogger returns the logger from the context.
func GetLogger(ctx context.Context) *log.Logger {
	return ctx.Value(loggerKey).(*log.Logger)
}
