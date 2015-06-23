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

package common

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestFilter(t *testing.T) {
	f1 := FilterFunc(
		func(w http.ResponseWriter, r *http.Request, next http.Handler) {
			fmt.Fprintf(w, "f1Before -> ")
			next.ServeHTTP(w, r)
			fmt.Fprintf(w, " -> f1After")
		})
	f2 := FilterFunc(
		func(w http.ResponseWriter, r *http.Request, next http.Handler) {
			fmt.Fprintf(w, "f2Before -> ")
			next.ServeHTTP(w, r)
			fmt.Fprintf(w, " -> f2After")
		})
	f3 := FilterFunc(
		func(w http.ResponseWriter, r *http.Request, next http.Handler) {
			fmt.Fprintf(w, "f3Before -> ")
			next.ServeHTTP(w, r)
			fmt.Fprintf(w, " -> f3After")
		})

	h := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "handler")
		})

	convey.Convey("Test Filter", t, func() {
		convey.Convey("Test FilterHandler with Filter", func() {
			handler := FilterHandler(f1, h)
			req, _ := http.NewRequest("GET", "/", nil)
			resp := httptest.NewRecorder()
			handler.ServeHTTP(resp, req)
			convey.So(resp.Body.String(), convey.ShouldEqual, "f1Before -> handler -> f1After")
		})
		convey.Convey("Test FilterHandler with nil Filter", func() {
			f := FilterHandler(nil, h)
			req, _ := http.NewRequest("GET", "/", nil)
			resp := httptest.NewRecorder()
			f.ServeHTTP(resp, req)
			convey.So(resp.Body.String(), convey.ShouldEqual, "handler")
		})
		convey.Convey("Test Combine2Filters with the first nil Filter", func() {
			f := Combine2Filters(nil, f1)
			handler := FilterHandler(f, h)
			req, _ := http.NewRequest("GET", "/", nil)
			resp := httptest.NewRecorder()
			handler.ServeHTTP(resp, req)
			convey.So(resp.Body.String(), convey.ShouldEqual, "f1Before -> handler -> f1After")
		})
		convey.Convey("Test Combine2Filters with the second nil Filter", func() {
			f := Combine2Filters(f1, nil)
			handler := FilterHandler(f, h)
			req, _ := http.NewRequest("GET", "/", nil)
			resp := httptest.NewRecorder()
			handler.ServeHTTP(resp, req)
			convey.So(resp.Body.String(), convey.ShouldEqual, "f1Before -> handler -> f1After")
		})
		convey.Convey("Test Combine2Filters", func() {
			f := Combine2Filters(f1, f2)
			handler := FilterHandler(f, h)
			req, _ := http.NewRequest("GET", "/", nil)
			resp := httptest.NewRecorder()
			handler.ServeHTTP(resp, req)
			convey.So(resp.Body.String(), convey.ShouldEqual, "f1Before -> f2Before -> handler -> f2After -> f1After")
		})
		convey.Convey("Test CombineFilters without Filter", func() {
			f := CombineFilters()
			convey.So(f, convey.ShouldBeNil)
		})
		convey.Convey("Test CombineFilters", func() {
			f := CombineFilters(f1, f2, f3)
			handler := FilterHandler(f, h)
			req, _ := http.NewRequest("GET", "/", nil)
			resp := httptest.NewRecorder()
			handler.ServeHTTP(resp, req)
			convey.So(resp.Body.String(), convey.ShouldEqual, "f1Before -> f2Before -> f3Before -> handler -> f3After -> f2After -> f1After")
		})
		convey.Convey("Test HandlerFunc", func() {
			handler := HandlerFunc(h, f1, f2)
			req, _ := http.NewRequest("GET", "/", nil)
			resp := httptest.NewRecorder()
			handler.ServeHTTP(resp, req)
			convey.So(resp.Body.String(), convey.ShouldEqual, "f1Before -> f2Before -> handler -> f2After -> f1After")
		})
		convey.Convey("Test Handle", func() {
			Handle("/foo", h, f1, f2)
			req, _ := http.NewRequest("GET", "/foo", nil)
			resp := httptest.NewRecorder()
			http.DefaultServeMux.ServeHTTP(resp, req)
			convey.So(resp.Body.String(), convey.ShouldEqual, "f1Before -> f2Before -> handler -> f2After -> f1After")
		})
		convey.Convey("Test HandleFunc", func() {
			HandleFunc("/bar", h, f1, f2)
			req, _ := http.NewRequest("GET", "/bar", nil)
			resp := httptest.NewRecorder()
			http.DefaultServeMux.ServeHTTP(resp, req)
			convey.So(resp.Body.String(), convey.ShouldEqual, "f1Before -> f2Before -> handler -> f2After -> f1After")
		})
	})
}
