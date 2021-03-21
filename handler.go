package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type FooHandler struct {
	FooService FooService `inject:""`
}

func (fh FooHandler) Foo(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return json.NewEncoder(w).Encode(fh.FooService.Foo(ctx, r.FormValue("input")))
}
