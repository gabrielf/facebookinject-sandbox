package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

func TestAppWithProdDeps(t *testing.T) {
	g := NewGomegaWithT(t)

	r := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()

	mux := SetupRoutes(CreateApp(ProdDeps()))
	mux.ServeHTTP(w, r)

	g.Expect(w.Code).To(Equal(http.StatusOK))
	g.Expect(w.Body).To(ContainSubstring(`"key":"value"`))
}

func TestAppWithDefaultTestDeps(t *testing.T) {
	g := NewGomegaWithT(t)

	r := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()

	mux := SetupRoutes(CreateApp(TestDeps()))
	mux.ServeHTTP(w, r)

	// Access the default mocks via the global var DefaultMocks
	g.Expect(DefaultMocks.InMemoryLogger.entries).To(HaveLen(1))
	g.Expect(DefaultMocks.InMemoryLogger.entries).To(ContainElement(MatchFields(IgnoreExtras, Fields{
		"Level": Equal("info"),
	})))
}

func TestAppWithOverriddenTestDeps(t *testing.T) {
	g := NewGomegaWithT(t)

	r := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()

	// This is how test dependencies are overridden
	testDeps := TestDeps(func(deps Deps) Deps {
		deps.FooService = &MockFooService{
			FooFunc: func(context.Context) interface{} {
				return "hello world"
			},
		}
		return deps
	})

	mux := SetupRoutes(CreateApp(testDeps))
	mux.ServeHTTP(w, r)

	g.Expect(w.Code).To(Equal(http.StatusOK))
	g.Expect(w.Body).To(ContainSubstring(`"hello world"`))
}
