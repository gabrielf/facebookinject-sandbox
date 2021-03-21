package main

import (
	"context"
)

type MockFooService struct {
	FooFunc func(context.Context, string) interface{}
}

func (mfs MockFooService) Foo(ctx context.Context, input string) interface{} {
	return mfs.FooFunc(ctx, input)
}

type InMemoryLogger struct {
	entries []entry
}

var _ = Logger(&InMemoryLogger{})

type entry struct {
	Level       string
	Msg         string
	KeysAndVals []interface{}
}

func (l *InMemoryLogger) Infow(_ context.Context, msg string, keysAndVals ...interface{}) {
	l.entries = append(l.entries, entry{
		Level:       "info",
		Msg:         msg,
		KeysAndVals: keysAndVals,
	})
}
