package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	app := CreateApp()
	SetupRoutes(app)
	if err := http.ListenAndServe("0.0.0.0:1337", http.DefaultServeMux); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}
