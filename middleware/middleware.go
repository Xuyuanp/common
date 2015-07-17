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

import "net/http"

// Middleware is a HTTP middleware.
type Middleware func(next http.Handler) http.Handler

// CombineMiddlewares combine multi Middleware into a single Middleware.
func CombineMiddlewares(ms ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		count := len(ms)
		for i := count - 1; i >= 0; i-- {
			m := ms[i]
			next = m(next)
		}
		return next
	}
}

// Handler returns a new http.Handler contains the original handler and Middlewares.
func Handler(handler http.Handler, ms ...Middleware) http.Handler {
	return CombineMiddlewares(ms...)(handler)
}

// HandlerFunc returns a new http.Handler contains the original handler function and Middlewares.
func HandlerFunc(fn func(http.ResponseWriter, *http.Request), ms ...Middleware) http.Handler {
	return Handler(http.HandlerFunc(fn), ms...)
}
