package main

import (
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
)

func TestHelloWorld(t *testing.T) {
	apitest.New().
		Handler(router).
		Get(path).
		Expect(t).
		// Body(`{"message": "hello world 1"}`).
		Assert(jsonpath.)
		Status(http.StatusOK).
		End()
}
