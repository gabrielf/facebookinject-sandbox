package main

type MockFooService struct {
	FooFunc func() interface{}
}

func (mfs MockFooService) Foo() interface{} {
	return mfs.FooFunc()
}
