package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/gomega"
)

func TestAppWithProdDeps(t *testing.T) {
	g := NewGomegaWithT(t)

	r := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()

	mux := SetupRoutes(CreateApp(ProdDeps()))
	mux.ServeHTTP(w, r)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(resp.StatusCode).To(Equal(http.StatusOK))
	g.Expect(string(body)).To(ContainSubstring(`"key":"value"`))
}

func TestAppWithDefaultMockDeps(t *testing.T) {
	g := NewGomegaWithT(t)

	r := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()

	mux := SetupRoutes(CreateApp(MockDeps()))

	// As the default mock runs a FooFunc that is not set it will panic with a nil pointer dereference
	g.Expect(func() {
		mux.ServeHTTP(w, r)
	}).To(Panic())
}

func TestAppWithMockDeps(t *testing.T) {
	g := NewGomegaWithT(t)

	r := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()

	// This is how mock dependencies for this particular test is setup
	mockDeps := MockDeps(func(deps *Deps) *Deps {
		deps.FooService = &MockFooService{
			FooFunc: func() interface{} {
				return "kaboom"
			},
		}
		return deps
	})

	mux := SetupRoutes(CreateApp(mockDeps))
	mux.ServeHTTP(w, r)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(resp.StatusCode).To(Equal(http.StatusOK))
	g.Expect(string(body)).To(ContainSubstring(`"kaboom"`))
}
