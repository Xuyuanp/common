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
	"crypto/md5"
	"hash"
	"io"
	"os"
)

// Hash computes the content from reader by the provided hash function
// and return the checksum, content length, and error.
func Hash(h hash.Hash, r io.Reader) ([]byte, int64, error) {
	n, err := io.Copy(h, r)
	if err != nil {
		return nil, 0, err
	}
	return h.Sum(nil), n, nil
}

// Hash computes the content from file by the provided hash function
// and return the checksum, content length, and error.
func HashFile(h hash.Hash, name string) ([]byte, int64, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()
	return Hash(h, file)
}

// MD5 computes the md5 checksum of the content of r
func MD5(r io.Reader) ([]byte, int64, error) {
	return Hash(md5.New(), r)
}

// MD5 computes the md5 checksum of file
func MD5File(name string) ([]byte, int64, error) {
	return HashFile(md5.New(), name)
}
