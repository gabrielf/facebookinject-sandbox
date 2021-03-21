package main

import (
	"context"
)

type MockFooService struct {
	FooFunc func() interface{}
}

func (mfs MockFooService) Foo(context.Context) interface{} {
	return mfs.FooFunc()
}

type InMemoryLogger struct {
	entries []entry
}

var _ = Logger(InMemoryLogger{})

type entry struct {
	level       string
	msg         string
	keysAndVals []interface{}
}

func (l InMemoryLogger) Infow(_ context.Context, msg string, keysAndVals ...interface{}) {
	l.entries = append(l.entries, entry{
		level:       "info",
		msg:         msg,
		keysAndVals: keysAndVals,
	})
}
