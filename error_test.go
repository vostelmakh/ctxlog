package log

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapErrorCtx(t *testing.T) {
	tests := []struct {
		name string
		ctx  context.Context
		err  error
	}{
		{
			name: "With logCtx",
			ctx:  context.WithValue(context.Background(), key, logCtx{}),
			err:  errors.New("original error"),
		},
		{
			name: "Without logCtx",
			ctx:  context.Background(),
			err:  errors.New("original error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WrapErrorCtx(tt.ctx, tt.err)

			var errorWithLogCtx *ErrorWithLogCtx
			assert.True(t, errors.As(err, &errorWithLogCtx))
		})
	}
}

func TestErrorCtx(t *testing.T) {
	ctx := context.Background()
	errCtx := &ErrorWithLogCtx{next: errors.New("original error"), ctx: logCtx{}}
	tests := []struct {
		name        string
		ctx         context.Context
		err         error
		expectedCtx context.Context
	}{
		{
			name:        "With ErrorWithLogCtx",
			ctx:         ctx,
			err:         errCtx,
			expectedCtx: context.WithValue(ctx, key, errCtx.ctx),
		},
		{
			name:        "Without ErrorWithLogCtx",
			ctx:         ctx,
			err:         errors.New("some other error"),
			expectedCtx: ctx,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, ErrorCtx(tt.ctx, tt.err), tt.expectedCtx)
		})
	}
}
