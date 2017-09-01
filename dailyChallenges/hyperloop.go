/**
Challenge: hyperloop
https://codefights.com/challenge/2RbXjnH68WacjvTrk


It's vacation time! Luckily, a nationwide network of Hyperloop tunnels has just been built, so all you have to do is find the lowest cost route to your destination.

Example
For city1 = "Los Angeles", city2 = "San Francisco" and

tunnels = [
  ["Los Angeles","San Francisco","1000"],
  ["Los Angeles","Fresno","150"],
  ["Fresno","San Francisco","300"]
]
the output should be
hyperloop(city1, city2, tunnels) = 450.

For this case, there are 2 options for a trip from LA to San Francisco. The first option is a direct trip costing 1000. The second option is a trip with one stop at Fresno costing 150 + 300 = 450. So, the lowest cost trip is 450.

Input/Output

[time limit] 4000ms (go)
[input] string city1

The city where the trip begins.

Guaranteed constraints:
2 ≤ city1.length < 15.

[input] string city2

The destination city.

Guaranteed constraints:
2 ≤ city2.length < 15.

[input] array.array.string tunnels

The list of one-way tunnels between 2 cities. Each tunnel is an array of 3 strings: tunnels[i][0] is the city at the beginning of the tunnel, tunnels[i][1] is the city at the end of the tunnel, and tunnels[i][2] is the cost of the trip for the tunnel.

Guaranteed constraints:
2 < tunnels.length < 50,
tunnels[i].length = 3,
2 ≤ tunnels[i][0].length < 15,
2 ≤ tunnels[i][1].length < 15,
0 ≤ int(tunnels[i][2]) < 1500.

[output] integer

The lowest possible cost of a trip from city1 to city2
**/
func hyperloop(src string, dest string, tunnels [][]string) int {

    // small bounds so simple DFS and pick shortest
    edges := make(map[string][]edge)    
    visited := make(map[string]bool)
    for _, v := range tunnels {
        edges[v[0]] = make([]edge, 0)
    }
    for _, v := range tunnels {
        edge := edge{ v[1], v[2]}
        edges[v[0]] = append(edges[v[0]], edge)
    }
   
    visited[src] = true
    return dfs(src, dest, edges, visited)
}

func dfs(node, dest string, edges map[string][]edge, visited map[string]bool) int {

    // we found the end
    if node == dest {
        return 0
    }
    
    // find the shortest path to dest node
    min := 999999 
    for _, edge := range edges[node] {
        if visited[edge.dest] {
            continue
        }      
        cost, _ := strconv.Atoi(edge.cost)
        visited[edge.dest] = true
        r := dfs(edge.dest, dest, edges, visited) + cost
        visited[edge.dest] = false
 
        if r < min {
            min = r
        }
    }
    return min
}

type edge struct {
    dest, cost string
}
