/*
Copyright 2023 Alexandre Mahdhaoui

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package graph

func unstack[T any](stack []T) ([]T, T) {
	switch i := len(stack); {
	case i > 1:
		return stack[:i-1], stack[i-1]
	case i == 1:
		return make([]T, 0), stack[0]
	default:
		panic("length of stack should be > 0")
	}
}

func dequeue[T any](q []T) ([]T, T) {
	switch i := len(q); {
	case i > 1:
		return q[1:], q[0]
	case i == 1:
		return make([]T, 0), q[0]
	default:
		panic("length of queue should be > 0")
	}
}
