package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/andviro/noodle"
	"github.com/facebookgo/inject"
)

type App struct {
	FooHandler *FooHandler `inject:""`
}

func CreateApp() App {
	var g inject.Graph
	var a App
	var fsh FooHandler
	var fss FooServiceImpl
	err := g.Provide(
		&inject.Object{Value: &a},
		&inject.Object{Value: &fsh},
		&inject.Object{Value: &fss},
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

func SetupRoutes(a App) {
	chain := noodle.New()
	http.DefaultServeMux.Handle("/foo", chain.Then(a.FooHandler.Foo))
}
