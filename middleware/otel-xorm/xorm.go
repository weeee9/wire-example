package otelxorm

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"xorm.io/builder"
	"xorm.io/xorm/contexts"
)

func NewTracingHook() contexts.Hook {
	return &tracingHook{}
}

type tracingHook struct{}

var _ contexts.Hook = &tracingHook{}

const (
	tracerName = "xorm-trace"
	spanName   = "xorm-query"
)

func (hook *tracingHook) BeforeProcess(c *contexts.ContextHook) (context.Context, error) {
	return c.Ctx, nil
}

func (h *tracingHook) AfterProcess(c *contexts.ContextHook) error {
	tracer := otel.Tracer(tracerName)

	_, span := tracer.Start(c.Ctx, spanName)
	defer span.End()

	if err := c.Err; err != nil {
		span.SetAttributes(
			attribute.Key("Err").String(err.Error()),
		)
	}

	sql, _ := builder.ConvertToBoundSQL(c.SQL, c.Args)
	span.SetAttributes(
		attribute.Key("SQL").String(sql),
		attribute.Key("Args").String(fmt.Sprint(c.Args)),
		attribute.Key("execute_time").String(c.ExecuteTime.String()),
	)

	return nil
}
