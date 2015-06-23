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
	"errors"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestAssert(t *testing.T) {
	convey.Convey("TestAssert", t, func() {
		convey.Convey("Test Assert", func() {
			convey.Convey("test false", func() {
				val := false
				err := errors.New("error")
				defer func() {
					if e := recover(); e != nil {
						convey.So(e, convey.ShouldEqual, err)
					} else {
						convey.So(e, convey.ShouldNotBeNil)
					}
				}()
				Assert(val, err)
			})
			convey.Convey("test true", func() {
				val := true
				err := errors.New("error")
				defer func() {
					if e := recover(); e != nil {
						convey.So(e, convey.ShouldBeNil)
					} else {
						convey.So(e, convey.ShouldBeNil)
					}
				}()
				Assert(val, err)
			})
		})
		convey.Convey("Test AssertFunc", func() {
			convey.Convey("test false", func() {
				val := false
				res := 1
				AssertFunc(val, func() {
					res = 2
				})
				convey.So(res, convey.ShouldEqual, 2)
			})
			convey.Convey("test true", func() {
				val := true
				res := 1
				AssertFunc(val, func() {
					res = 2
				})
				convey.So(res, convey.ShouldEqual, 1)
			})
		})
	})
}
