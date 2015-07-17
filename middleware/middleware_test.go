/*
 * Copyright 2015 Xuyuan Pang
 * Author: Xuyuan Pang
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newMiddleware(name string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "%s Before -> ", name)
				next.ServeHTTP(w, r)
				fmt.Fprintf(w, " -> %s After", name)
			},
		)
	}

}

var handler = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "handler")
}

func assertEqual(t *testing.T, actual, expected interface{}) {
	if actual != expected {
		t.Errorf("Not equal. %#v expected, but %#v actual", expected, actual)
	}
}

func TestMiddleware(t *testing.T) {
	m1 := newMiddleware("m1")

	h := m1(http.HandlerFunc(handler))

	req, _ := http.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	h.ServeHTTP(resp, req)

	assertEqual(t, resp.Body.String(),
		"m1 Before -> handler -> m1 After",
	)
}

func TestCombineMiddleware(t *testing.T) {
	m1 := newMiddleware("m1")
	m2 := newMiddleware("m2")

	m3 := CombineMiddlewares(m1, m2)

	h := m3(http.HandlerFunc(handler))

	req, _ := http.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	h.ServeHTTP(resp, req)

	assertEqual(t, resp.Body.String(),
		"m1 Before -> m2 Before -> handler -> m2 After -> m1 After",
	)
}

func runRequest(b *testing.B, handler http.Handler) {
	req, _ := http.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		resp.Body.Reset()
		handler.ServeHTTP(resp, req)
	}
}

func BenchmarkMiddleware(b *testing.B) {
	m := func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				next.ServeHTTP(w, r)
			},
		)
	}

	handler := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {},
	)

	runRequest(b, m(handler))
}

func BenchmarkCombineMiddleware(b *testing.B) {
	m := func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				next.ServeHTTP(w, r)
			},
		)
	}
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		CombineMiddlewares(m, m, m, m)
	}

}

func BenchmarkHTTP(b *testing.B) {

}
