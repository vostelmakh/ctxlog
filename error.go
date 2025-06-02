package log

import (
	"context"
	"errors"
)

type ErrorWithLogCtx struct {
	next error
	ctx  logCtx
}

func WrapErrorCtx(ctx context.Context, err error) error {
	lCtx, _ := ctx.Value(key).(logCtx)

	return &ErrorWithLogCtx{
		next: err,
		ctx:  lCtx,
	}
}

func ErrorCtx(ctx context.Context, err error) context.Context {
	var errorWithLCtx *ErrorWithLogCtx
	
	if errors.As(err, &errorWithLCtx) {
		ctx = context.WithValue(ctx, key, errorWithLCtx.ctx)
	}

	return ctx
}

func (e *ErrorWithLogCtx) Error() string {
	if e.next != nil {
		return e.next.Error()
	}

	return ""
}

func (e *ErrorWithLogCtx) Unwrap() error {
	if e.next != nil {
		return e.next
	}

	return nil
}
