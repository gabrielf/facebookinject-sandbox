package main

import (
	"context"
)

type MockFooService struct {
	FooFunc func(context.Context) interface{}
}

func (mfs MockFooService) Foo(ctx context.Context) interface{} {
	return mfs.FooFunc(ctx)
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
