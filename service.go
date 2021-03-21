package main

import "context"

type FooService interface {
	Foo(ctx context.Context, input string) interface{}
}

type FooServiceImpl struct {
	Logger Logger `inject:""`
}

func (fs FooServiceImpl) Foo(ctx context.Context, input string) interface{} {
	fs.Logger.Infow(ctx, "Foo called", "input", input)
	return map[string]string{
		"input": input,
	}
}
