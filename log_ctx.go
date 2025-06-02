package log

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logCtx struct {
	OrderNr     string
	MessageID   string
	HandlerName string
	MsgKey      string
}

type keyType int

const key = keyType(0)

func Fields(ctx context.Context, fields []zapcore.Field) []zapcore.Field {
	l, ok := ctx.Value(key).(logCtx)
	if !ok {
		return fields
	}

	if l.MessageID != "" {
		fields = append(fields, zap.String("message_id", l.MessageID))
	}

	if l.OrderNr != "" {
		fields = append(fields, zap.String("order_nr", l.OrderNr))
	}

	if l.HandlerName != "" {
		fields = append(fields, zap.String("handler_name", l.HandlerName))
	}

	if l.MsgKey != "" {
		fields = append(fields, zap.String("msg_key", l.MsgKey))
	}

	return fields
}

func msg(ctx context.Context, msg string) string {
	if l, ok := ctx.Value(key).(logCtx); ok {
		if l.HandlerName != "" {
			return l.HandlerName + ": " + msg
		}
	}

	return msg
}

func WithOrderNr(ctx context.Context, orderNr string) context.Context {
	if l, ok := ctx.Value(key).(logCtx); ok {
		l.OrderNr = orderNr

		return context.WithValue(ctx, key, l)
	}

	return context.WithValue(ctx, key, logCtx{OrderNr: orderNr})
}

func WithHandlerName(ctx context.Context, handlerName string) context.Context {
	if l, ok := ctx.Value(key).(logCtx); ok {
		l.HandlerName = handlerName

		return context.WithValue(ctx, key, l)
	}

	return context.WithValue(ctx, key, logCtx{HandlerName: handlerName})
}

func WithMsgKey(ctx context.Context, msgKey string) context.Context {
	if l, ok := ctx.Value(key).(logCtx); ok {
		l.MsgKey = msgKey

		return context.WithValue(ctx, key, l)
	}

	return context.WithValue(ctx, key, logCtx{MsgKey: msgKey})
}

func WithMessageID(ctx context.Context, messageID string) context.Context {
	if l, ok := ctx.Value(key).(logCtx); ok {
		l.MessageID = messageID

		return context.WithValue(ctx, key, l)
	}

	return context.WithValue(ctx, key, logCtx{MessageID: messageID})
}
