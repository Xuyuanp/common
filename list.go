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

import "container/list"

type container interface {
	Push(interface{})
	Pop() interface{}
	Len() int
	Reset()
}

type Stack interface {
	container
}

func NewStack() Stack {
	return &stack{list.New()}
}

type stack struct {
	*list.List
}

func (s *stack) Push(v interface{}) {
	s.List.PushBack(v)
}

func (s *stack) Pop() interface{} {
	e := s.Back()
	s.Remove(e)
	return e.Value
}

func (s *stack) Reset() {
	s.List.Init()
}

type Queue interface {
	container
}

func NewQueue() Queue {
	return &queue{list.New()}
}

type queue struct {
	*list.List
}

func (q *queue) Push(v interface{}) {
	q.List.PushBack(v)
}

func (q *queue) Pop() interface{} {
	e := q.List.Front()
	q.List.Remove(e)
	return e.Value
}

func (q *queue) Len() int {
	return q.List.Len()
}

func (q *queue) Reset() {
	q.List.Init()
}
