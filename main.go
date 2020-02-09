package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	mux := SetupRoutes(CreateApp())
	if err := http.ListenAndServe("0.0.0.0:1337", mux); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}
