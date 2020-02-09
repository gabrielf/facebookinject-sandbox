package main

type FooService interface {
	Foo() interface{}
}

type FooServiceImpl struct {
}

func (fs FooServiceImpl) Foo() interface{} {
	return map[string]interface{}{
		"key": "value",
	}
}
