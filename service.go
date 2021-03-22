package main

import (
	"context"
	"fmt"
)

type FooService interface {
	Foo(ctx context.Context, input string) interface{}
}

type FooServiceImpl struct {
	Logger        Logger        `inject:""`
	ErrorReporter ErrorReporter `inject:""`
}

func (fs FooServiceImpl) Foo(ctx context.Context, input string) interface{} {
	fs.Logger.Infow(ctx, "Foo called", "input", input)
	if input == "" {
		fs.ErrorReporter.ReportError(fmt.Errorf("empty input"))
	}
	return map[string]string{
		"input": input,
	}
}
