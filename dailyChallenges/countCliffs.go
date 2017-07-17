/**
Challenge: countCliffs


You have a seaMap of part of the Caribbean Sea, represented as a 2D array. The map shows the positions of dangerous cliffs that a wise sea dog should steer clear of.

As the captain of the most notorious pirate ship on the high seas, you're not afraid of these little rocks and are willing to navigate your ship through them. However, you're mildly curious about the number of cliff lines in the area. You assume that cliffs form a line if they are connected to each other, i.e. if each cliff has another cliff directly to the left, to the right, above, or below it.

Looking at the seaMap, calculate the number of cliff lines on it.

Example

For

seaMap = [[0,1,1], 
          [0,1,1], 
          [0,0,0]]
the output should be
countCliffs(seaMap) = 1.

There is only one cliff line on the seaMap, which consists of 4 cliffs and forms a square.

Input/Output

[time limit] 4000ms (go)
[input] array.array.integer seaMap

A rectangular array representing your map of the Caribbean. seaMap[i][j] = 1 if there's a cliff at (i, j), and seaMap[i][j] = 0 otherwise.

Guaranteed constraints:
3 ≤ seaMap.length ≤ 100,
3 ≤ seaMap[0].length ≤ 100,
0 ≤ seaMap[i][j] ≤ 1.

[output] integer

The number of cliff lines.
**/
func countCliffs(s [][]int) int {
   
    c := 0
    for n, z := range s {
        for m, _ := range z {
            if s[n][m] == 1 {
                c++
                v(n, m, &s)
            }
        }
    } 
    return c  
}

func v(i, j int, s *[][]int) {
    (*s)[i][j] = 0
    // above
    if i > 0 && (*s)[i-1][j] > 0 {
        v(i-1, j, s)
    }
    // down
    if i < len(*s)-1 && (*s)[i+1][j] > 0 {
        v(i+1, j, s)
    }
    // left
    if j > 0 && (*s)[i][j-1] > 0 {
        v(i, j-1, s)
    }
    // right 
    if j < len((*s)[0])-1 && (*s)[i][j+1] > 0 {
        v(i, j+1, s)
    }
}