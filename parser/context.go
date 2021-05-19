package parser

import (
	"bytes"
	"context"
	"errors"
)

type Key int

// These are keys for the context that floats around.
const (
	KeyHyphaName Key = iota
	KeyInputBuffer
	KeyRecursionLevel
)

// ParsingDone is returned by Context when the parsing is done because there is no more inputFrom.
var ParsingDone = errors.New("parsing done")

// ContextFromStringInput returns the context for the given inputFrom.
func ContextFromStringInput(hyphaName, input string) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(
		context.WithValue(
			context.WithValue(
				context.WithValue(
					context.Background(),
					KeyHyphaName,
					hyphaName),
				KeyInputBuffer,
				bytes.NewBufferString(input),
			),
			KeyRecursionLevel,
			0,
		),
	)
	return ctx, cancel
}

func hyphaNameFrom(ctx context.Context) string {
	return ctx.Value(KeyHyphaName).(string)
}

func inputFrom(ctx context.Context) *bytes.Buffer {
	return ctx.Value(KeyInputBuffer).(*bytes.Buffer)
}
