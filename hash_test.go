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
	"bytes"
	"crypto/md5"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

func TestHash(t *testing.T) {
	convey.Convey("Test hash", t, func() {
		data := []byte("foobar")
		buf := bytes.NewBuffer(data)
		sum, n, err := Hash(md5.New(), buf)
		msum := md5.Sum(data)
		convey.So(err, convey.ShouldBeNil)
		convey.So(n, convey.ShouldEqual, len(data))
		convey.So(string(sum), convey.ShouldEqual, string(msum[:]))
	})

	convey.Convey("Test HashFile", t, func() {
		tmpfile := fmt.Sprintf("tmp-%d", time.Now().Unix())
		file, err := os.OpenFile(tmpfile, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0600)
		convey.So(err, convey.ShouldBeNil)
		data := []byte("foobar")
		_, err = file.Write(data)
		convey.So(err, convey.ShouldBeNil)
		file.Close()
		defer func() {
			os.Remove(tmpfile)
		}()

		sum, n, err := HashFile(md5.New(), tmpfile)
		convey.So(err, convey.ShouldBeNil)
		convey.So(n, convey.ShouldEqual, len(data))
		msum := md5.Sum(data)
		convey.So(string(sum), convey.ShouldEqual, string(msum[:]))
	})

	convey.Convey("Test MD5", t, func() {
		data := []byte("foobar")
		buf := bytes.NewBuffer(data)
		sum, n, err := MD5(buf)
		msum := md5.Sum(data)
		convey.So(err, convey.ShouldBeNil)
		convey.So(n, convey.ShouldEqual, len(data))
		convey.So(string(sum), convey.ShouldEqual, string(msum[:]))
	})

	convey.Convey("Test MD5File", t, func() {
		tmpfile := fmt.Sprintf("tmp-%d", time.Now().Unix())
		file, err := os.OpenFile(tmpfile, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0600)
		convey.So(err, convey.ShouldBeNil)
		data := []byte("foobar")
		_, err = file.Write(data)
		convey.So(err, convey.ShouldBeNil)
		file.Close()
		defer func() {
			os.Remove(tmpfile)
		}()

		sum, n, err := MD5File(tmpfile)
		convey.So(err, convey.ShouldBeNil)
		convey.So(n, convey.ShouldEqual, len(data))
		msum := md5.Sum(data)
		convey.So(string(sum), convey.ShouldEqual, string(msum[:]))
	})
}
