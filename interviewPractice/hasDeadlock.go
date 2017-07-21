/**
Challenge: hasDeadlock
https://codefights.com/interview-practice/task/ZTqKqNwK6ZL6GDpJ5/description

Note: Try to solve this task in O(m + n) or O(n) time, where n is a number of vertices and m is a number of edges, since this is what you'll be asked to do during an interview.

In a multithreaded environment, it's possible that different processes will need to use the same resource. A wait-for graph represents the different processes as nodes in a directed graph, where an edge from node i to node j means that the node j is using a resource that node i needs to use (and cannot use until node j releases it).

We are interested in whether or not this digraph has any cycles in it. If it does, it is possible for the system to get into a state where no process can complete.

We will represent the processes by integers 0, ...., n - 1. We represent the edges using a two-dimensional list connections. If j is in the list connections[i], then there is a directed edge from process i to process j.

Write a function that returns True if connections describes a graph with a directed cycle, or False otherwise.

Example

For connections = [[1], [2], [3, 4], [4], [0]], the output should be
hasDeadlock(connections) = true.


This graph contains a cycle.

For connections = [[1, 2, 3], [2, 3], [3], []], the output should be
hasDeadlock(connections) = false.


This graph doesn't contain a directed cycle (there are two paths from 0 to 3, but no paths from 3 back to 0).

Input/Output

[time limit] 4000ms (go)
[input] array.array.integer connections

A representation of the graphs edges. connection.length is the number of vertices. If j is in the list connections[i], then there is a directed edge from process i to process j.

Guaranteed constraints:
1 ≤ connections.length ≤ 500,
0 ≤ connections[i][j] < connections.length,
connections[i][j] ≠ connections[i][k] for j ≠ k,
i not in connections[i].

[output] boolean

Return True if connections describes a graph with a directed cycle, or False otherwise.
**/
func hasDeadlock(connections [][]int) bool {

    stack := make([]int, 0)
    visited := make(map[int]bool)
    
    ret := false
    for i := 0; i < len(connections); i++ {
        if !visited[i] {
            ret = ret || dfs(i, stack, connections, visited)
        }
    }
    
    return ret
}

func dfs(node int, stack []int, c [][]int, visited map[int]bool) bool {
    for _, v := range stack {
        if v == node {
            return true
        }
    }
    visited[node] = true
    stack = append(stack, node)
    ret := false
    
    for _, v := range c[node] {
        ret = ret || dfs(v, stack, c, visited)
    }
    return ret
}
