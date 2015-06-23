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
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestList(t *testing.T) {
	convey.Convey("TestList", t, func() {
		convey.Convey("Test Stack", func() {
			stack := NewStack()
			vals := []string{"a", "b", "c"}
			for _, v := range vals {
				stack.Push(v)
			}
			convey.So(stack.Len(), convey.ShouldEqual, 3)
			for i := len(vals); i > 0; i-- {
				convey.So(stack.Pop(), convey.ShouldEqual, vals[i-1])
			}
			stack.Reset()
			convey.So(stack.Len(), convey.ShouldEqual, 0)
		})
		convey.Convey("Test Queue", func() {
			queue := NewQueue()
			vals := []string{"a", "b", "c"}
			for _, v := range vals {
				queue.Push(v)
			}
			convey.So(queue.Len(), convey.ShouldEqual, 3)
			for i := 0; i < len(vals); i++ {
				convey.So(queue.Pop(), convey.ShouldEqual, vals[i])
			}
			queue.Reset()
			convey.So(queue.Len(), convey.ShouldEqual, 0)
		})
	})
}
