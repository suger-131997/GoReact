package contextutil

import "context"

type stateCodeKey struct{}

func WithStateCode(ctx context.Context, stateCode int) context.Context {
	return context.WithValue(ctx, stateCodeKey{}, stateCode)
}

func StateCodeFromContext(ctx context.Context) (int, bool) {
	code, ok := ctx.Value(stateCodeKey{}).(int)
	return code, ok
}
