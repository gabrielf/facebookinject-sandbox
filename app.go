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
// Order deps alphabetically to avoid having to think when adding one.
type Deps struct {
	BarService BarService
	FooService FooService
	Logger     Logger
}

// CreateApp uses Deps to inject all app handlers
func CreateApp(deps Deps) App {
	var g inject.Graph
	var a App
	err := g.Provide(
		&inject.Object{Value: &a},
		&inject.Object{Value: deps.BarService},
		&inject.Object{Value: deps.FooService},
		&inject.Object{Value: deps.Logger},
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
		BarService: &BarServiceImpl{},
		FooService: &FooServiceImpl{},
		Logger:     &StdoutLogger{},
	}
}

// TestDeps return the default test deps but allows these to be overridden by
// one or more functions that get the chance to set their own dependencies.
func TestDeps(alterDeps ...func(Deps) Deps) Deps {
	testDeps := DefaultTestDeps()
	for _, alter := range alterDeps {
		testDeps = alter(testDeps)
	}
	return testDeps
}

// DefaultTestDeps returns the prod deps but tries to replace all dependencies
// that talk to external systems have been replaced with their mock equivalents.
// This includes persistence, logger, error reporter, message queues etc.
func DefaultTestDeps() Deps {
	DefaultMocks = &MockDeps{
		Logger: &InMemoryLogger{},
	}

	deps := ProdDeps()
	deps.Logger = DefaultMocks.Logger
	return deps
}

// DefaultMocks contains the mock deps as set by DefaultTestDeps. This makes
// it possible to access these mocks in tests to verify interactions.
// Unfortunately this makes parallel tests impossible as it is a global var.
var DefaultMocks *MockDeps

type MockDeps struct {
	Logger *InMemoryLogger
}

func SetupRoutes(a App) http.Handler {
	chain := noodle.New()

	mux := http.NewServeMux()
	mux.Handle("/foo", chain.Then(a.FooHandler.Foo))

	return mux
}
