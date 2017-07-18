/**
Challenge: addEven
https://codefights.com/tournaments/vbZdzDTcajNBS7xN6/E

Given an undirected graph graph that is represented by its adjacency matrix, return whether or not is it possible to add no more than two edges to this graph in order to make all the degrees of nodes even. Keep in mind that in the resulting graph there should be at most one edge between any pair of nodes.

Example

For

graph = [[false, true, false, false],
         [true, false, true, false],
         [false, true, false, true],
         [false, false, true, false]]
the output should be
addEven(graph) = true.



Nodes 1 and 4 have odd degrees. By connecting them with an edge, we get a graph where all the nodes have even degrees.

For

graph = [[false, true, true, true],
         [true, false, true, false],
         [true, true, false, true],
         [true, false, true, false]]
the output should be
addEven(graph) = false.



Nodes 1 and 3 have odd degrees. But they are already connected with an edge, and there is no way to add one or two edges to get only even degrees in this graph.

For

graph = [[false, true, false, false],
         [true, false, false, false],
         [false, false, false, true],
         [false, false, true, false]]
the output should be
addEven(graph) = true.



All the nodes have odd degrees. By adding edges (1, 4) and (2, 3), we get a graph where all the nodes have even degrees.

Input/Output

[time limit] 4000ms (go)
[input] array.array.boolean graph

The adjacency matrix of the given graph.

Guaranteed constraints:
1 ≤ graph.length ≤ 50,
graph[i].length = graph.length,
graph[i][j] = graph[j][i].

[output] boolean

Return true if it's possible to add no more than two edges to make a graph where all the nodes have only even degrees, otherwise return false.
**/
func addEven(graph [][]bool) bool {

    // true if num odd nodes not already connected == 0, 2, 4  
    blah := make([]int, len(graph))
    for i := 0; i < len(graph); i++ {
        for j := 0; j < len(graph[0]); j++ {
            if graph[i][j] {
                blah[i]++
            }
        }
    }
    odd := 0
    z := make([]int, 0)
    fmt.Println(blah)
    for n, v := range blah {
        fmt.Println(n)
        if v % 2 == 1  {
            odd++
            z = append(z, n)
        }
    }
    
    if (odd > 4) {
        return false
    }
    
    if (odd == 0) {
        return true
    }
    if (odd % 2 == 1) {
        return false
    }
    
    fmt.Println("odd: ", z)
    
    if (odd == 4) {
        // can we pair up
        if (!graph[z[0]][z[1]] && !graph[z[1]][z[2]]) || 
        (!graph[z[1]][z[2]] && !graph[z[0]][z[3]]) || 
        (!graph[z[1]][z[3]] && !graph[z[0]][z[2]]) {
            return true
        }
        return false
    }
    
    // is there an even node that both odd nodes don't connect to
    for n, v := range blah {
        if !graph[z[0]][z[1]] {
            return true
        }
        if v % 2 == 0 {
            
            fmt.Printf("EvenNode n: %v v: %v\n", n, v)
            if (!graph[n][z[0]]) && (!graph[n][z[1]]) {
                return true
            }
        } 
        
    }
    
    return false
}
