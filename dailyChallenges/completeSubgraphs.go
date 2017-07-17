/*
challenge: completeSubgraphs

https://codefights.com/challenge/tK2Zzw5WjkSpbFuRv/solutions

You are given a simple graph and the number of vertices in it, n. Return the number of complete subgraphs of the given graph that have more than two vertices. A subgraph is said to be complete if all the vertices are connected by an edge.

Example
For n = 6 and
graph = [['A', 'B'], ['B', 'C'], ['B', 'D'], ['C', 'D'], ['E', 'F'], ['F', 'B'], ['F', 'C'], ['F', 'D']], the output should be
completeSubgraphs(n, graph) = 5.
There are five complete subgraphs: BDC, BDF, BCF, DCF, and BDCF.

Input/Output

[time limit] 4000ms (go)
[input] integer n

The number of vertices in the given graph.

Guaranteed constraints:
0 ≤ n ≤ 15.

[input] array.array.char graph

The matrix representing the given graph. graph[i] denotes an edge between the vertices graph[i][0] and graph[i][1]. Vertices are represented using uppercase English letters.

Guaranteed constraints:
0 ≤ graph.length < 100,
graph[i].length = 2.

[output] integer

The number of complete subgraphs of the given graph that have more than two vertices. A subgraph is complete if all the vertices are connected with an edge.
*/

import "strings"
import "sort"

var perms = make(map[string]bool)
func completeSubgraphs(n int, graph [][]string) int {
    if len(graph) < 1 || len(graph[0]) < 1 {
        return 0
    }

    m := make(map[string]map[string]bool)
    for _, pair := range graph {
        if m[pair[0]] == nil {
            m[pair[0]] = make(map[string]bool)
        }
        if m[pair[1]] == nil {
            m[pair[1]] = make(map[string]bool)
        }
        m[pair[0]][pair[1]] = true
        m[pair[1]][pair[0]] = true
    }
    
    for k, v := range m {
        s := []string{k}
        for sk := range v {
            s = append(s, sk)
        }
        sort.Strings(s)
        combos("", strings.Join(s, ""))
    }

    res := 0
    // see if everything in this permutation is connected
    for perm := range perms {
        arr := strings.Split(perm, "")
        con := true
        for _, str := range arr {
            for _, val := range arr {
                if str != val && !m[str][val] {
                    con = false
                }
            }
        }
        if con {
            res++
        }
        
    }
  
    return res
}

// generate all combinations
func combos(cur, rem string) {
    if len(cur) > 2 {
        perms[cur] = true
    }
    if len(rem) < 1 {
        return
    }
    combos(cur + string(rem[0]), rem[1:])
    combos(cur, rem[1:])
}

