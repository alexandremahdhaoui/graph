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

// DFS is an implementation of depth-first search on the Node type.
func DFS(node *Node, f func(*Node) bool) bool {
	visitedNodes := make(NodeSet)
	var stack []*Node

	return dfs(node, stack, visitedNodes, concreteVisitor(f))
}

// dfs is the internal recursive func used by DFS.
//
// Returns a boolean, that can be returned as true from concreteVisitor.Visit().
// The return bool let Visitor stop the search.
func dfs(node *Node, stack []*Node, visitedNodes NodeSet, visitor Visitor) bool {
	if !visited(node, visitedNodes) {
		visitedNodes[node] = nil
		stack = append(stack, node)

		if stop := visitor.Visit(node); stop {
			return stop
		}
	}

	for _, adj := range node.AdjacentNodes() {
		if visited(adj, visitedNodes) {
			continue
		}

		if stop := dfs(adj, stack, visitedNodes, visitor); stop {
			return stop
		}
	}

	for len(stack) > 0 {
		stack, node = unstack(stack)
		if stop := dfs(node, stack, visitedNodes, visitor); stop {
			return stop
		}
	}

	return false
}

func BFS(node *Node, f func(*Node) bool) bool {
	visitedNodes := make(NodeSet)
	q := make([]*Node, 0)

	return bfs(node, q, visitedNodes, concreteVisitor(f))
}

func bfs(node *Node, q []*Node, visitedNodes NodeSet, visitor Visitor) bool {
	if !visited(node, visitedNodes) {
		visitedNodes[node] = nil
		if stop := visitor.Visit(node); stop {
			return stop
		}
	}

	for _, adj := range node.AdjacentNodes() {
		if visited(node, visitedNodes) {
			continue
		}

		q = append(q, adj)
	}

	for len(q) > 0 {
		q, node = dequeue(q)
		if stop := bfs(node, q, visitedNodes, visitor); stop {
			return stop
		}
	}

	return false
}

type NodeSet map[*Node]interface{}

func visited(node *Node, visitedNodes NodeSet) bool {
	_, ok := visitedNodes[node]
	return ok
}

type Visitor interface {
	Visit(node *Node) bool
}

type concreteVisitor func(*Node) bool

func (f concreteVisitor) Visit(node *Node) bool {
	return f(node)
}
