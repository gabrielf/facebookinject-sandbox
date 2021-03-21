package main

import (
	"context"
	"encoding/json"
	"os"
)

type Logger interface {
	Infow(ctx context.Context, msg string, keysAndVals ...interface{})
}

type StdoutLogger struct {
}

var _ = Logger(StdoutLogger{})

func (l StdoutLogger) Infow(_ context.Context, msg string, keysAndVals ...interface{}) {
	entry := map[string]interface{}{
		"level": "info",
		"msg":   msg,
	}

	// Just example code, will panic on malformed keysAndVals,
	for i := 0; i < len(keysAndVals); i += 2 {
		entry[keysAndVals[i].(string)] = keysAndVals[i+1]
	}

	if err := json.NewEncoder(os.Stdout).Encode(entry); err != nil {
		panic(err.Error())
	}
}
