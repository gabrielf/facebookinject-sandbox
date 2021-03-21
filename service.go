package main

import "context"

type FooService interface {
	Foo(context.Context) interface{}
}

type FooServiceImpl struct {
	Logger Logger `inject:""`
}

func (fs FooServiceImpl) Foo(ctx context.Context) interface{} {
	fs.Logger.Infow(ctx, "Foo called", "key", "val")
	return map[string]interface{}{
		"key": "value",
	}
}
