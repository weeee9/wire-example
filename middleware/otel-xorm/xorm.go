package otelxorm

import (
	"context"

	"github.com/opentracing/opentracing-go"
	tracerLog "github.com/opentracing/opentracing-go/log"
	"github.com/rs/zerolog/log"
	"xorm.io/builder"
	"xorm.io/xorm/contexts"
)

type xormHookContext string

var xormHookContextKey = xormHookContext("xorm")

func NewTracingHook() contexts.Hook {
	return &tracingHook{}
}

type tracingHook struct{}

var _ contexts.Hook = &tracingHook{}

func (hook *tracingHook) BeforeProcess(ctx *contexts.ContextHook) (context.Context, error) {
	span, _ := opentracing.StartSpanFromContext(ctx.Ctx, "xorm-hook")

	ctx.Ctx = context.WithValue(ctx.Ctx, xormHookContextKey, span)

	return ctx.Ctx, nil
}

func (h *tracingHook) AfterProcess(c *contexts.ContextHook) error {
	span, ok := c.Ctx.Value(xormHookContextKey).(opentracing.Span)
	if !ok {
		log.Warn().Msg("got no span")
		return nil
	}
	defer span.Finish()

	if c.Err != nil {
		log.Error().Err(c.Err).Msg("got some error")
	}

	sql, err := builder.ConvertToBoundSQL(c.SQL, c.Args)
	if err != nil {
		log.Error().Err(c.Err).Msg("failed to get SQL")
	}

	span.LogFields(tracerLog.String("SQL", sql))
	span.LogFields(tracerLog.Object("args", c.Args))
	span.SetTag("execute_time", c.ExecuteTime)

	return nil
}
