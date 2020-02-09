package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/andviro/noodle"
	"github.com/facebookgo/inject"
)

// App contain all the handlers
type App struct {
	FooHandler *FooHandler `inject:""`
}

// Deps contain all the services and other dependencies such as logger, stores etc
type Deps struct {
	FooService FooService
}

// CreateApp uses Deps to inject all app handlers
func CreateApp(deps Deps) App {
	var g inject.Graph
	var a App
	err := g.Provide(
		&inject.Object{Value: &a},
		&inject.Object{Value: deps.FooService},
	)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := g.Populate(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return a
}

func ProdDeps() Deps {
	return Deps{
		FooService: &FooServiceImpl{},
	}
}

// MockDeps return the default mock deps but these can be overridden by
// providing one or more functions that get the chance to set their own
// mocks.
func MockDeps(alterDeps ...func(Deps) Deps) Deps {
	mockDeps := Deps{
		FooService: &MockFooService{},
	}
	for _, alter := range alterDeps {
		mockDeps = alter(mockDeps)
	}
	return mockDeps
}

func SetupRoutes(a App) http.Handler {
	chain := noodle.New()

	mux := http.NewServeMux()
	mux.Handle("/foo", chain.Then(a.FooHandler.Foo))

	return mux
}
