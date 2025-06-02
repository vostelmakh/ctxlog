package log

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestLogCtx(t *testing.T) {
	tests := []struct {
		name        string
		withFunc    func(ctx context.Context) context.Context
		extraFields []zapcore.Field
		resFields   []zapcore.Field
	}{
		{
			name: "handler_name",
			withFunc: func(ctx context.Context) context.Context {
				return WithHandlerName(ctx, "handler")
			},
			extraFields: nil,
			resFields:   []zapcore.Field{zap.String("handler_name", "handler")},
		},
		{
			name: "order_nr",
			withFunc: func(ctx context.Context) context.Context {
				return WithOrderNr(ctx, "ORD-123")
			},
			extraFields: nil,
			resFields:   []zapcore.Field{zap.String("order_nr", "ORD-123")},
		},
		{
			name: "msg_key",
			withFunc: func(ctx context.Context) context.Context {
				newCtx := WithMsgKey(ctx, "key")

				return newCtx
			},
			extraFields: nil,
			resFields:   []zapcore.Field{zap.String("msg_key", "key")},
		},
	}

	for _, tt := range tests {
		ctx := context.Background()
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.resFields, Fields(tt.withFunc(ctx), tt.extraFields))
		})
	}
}
