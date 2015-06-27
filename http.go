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

import "net/http"

// Filter processes request before handler.
type Filter interface {
	Filter(http.ResponseWriter, *http.Request, http.Handler)
}

// FilterFunc is a function as a Filter.
type FilterFunc func(http.ResponseWriter, *http.Request, http.Handler)

// Filter implements Filter interface.
func (ff FilterFunc) Filter(w http.ResponseWriter, r *http.Request, next http.Handler) {
	ff(w, r, next)
}

// FilterHandler returns a new Handler which filters request before the Handler.
func FilterHandler(filter Filter, handler http.Handler) http.Handler {
	if filter == nil {
		return handler
	}
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			filter.Filter(w, req, handler)
		})
}

type filterHandler struct {
	filter  Filter
	handler http.Handler
}

func (fh *filterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if fh.filter == nil {
		fh.handler.ServeHTTP(w, r)
		return
	}
	fh.filter.Filter(w, r, fh.handler)
}

// Combine2Filters combines 2 Filters into a single Filter.
func Combine2Filters(first Filter, second Filter) Filter {
	// return the other if one is nil
	// return nil if both are nil
	if first == nil {
		return second
	}
	if second == nil {
		return first
	}

	nnext := &filterHandler{filter: second}
	return FilterFunc(
		func(w http.ResponseWriter, req *http.Request, next http.Handler) {
			// nnext := FilterHandler(second, next)
			nnext.handler = next
			first.Filter(w, req, nnext)
		})
}

// CombineFilters combines multi Filters into a single Filter.
func CombineFilters(filters ...Filter) Filter {
	// if no filters provided, return nil
	if filters == nil || len(filters) == 0 {
		return nil
	}
	// first Filter
	first := filters[0]
	// combine all others as the second Filter
	second := CombineFilters(filters[1:]...)
	// combine the two Filters
	return Combine2Filters(first, second)
}

// Handler combine all Filters and the Handler into a new Handler which
// contains these Filters and calls Handler finally.
func Handler(handler http.Handler, filters ...Filter) http.Handler {
	filter := CombineFilters(filters...)
	return &filterHandler{filter: filter, handler: handler}
}

// Handler combine all Filters and the HandlerFunc into a new Handler which
// contains these Filters and calls Handler finally.
func HandlerFunc(hf func(http.ResponseWriter, *http.Request), filters ...Filter) http.Handler {
	return Handler(http.HandlerFunc(hf), filters...)
}

// Handle is a easy way to call http.Handle function.
func Handle(pattern string, handler http.Handler, filters ...Filter) {
	http.Handle(pattern, Handler(handler, filters...))
}

// HandleFunc is a easy way to call http.HandleFunc function.
func HandleFunc(pattern string, hf func(http.ResponseWriter, *http.Request), filters ...Filter) {
	http.Handle(pattern, HandlerFunc(hf, filters...))
}
