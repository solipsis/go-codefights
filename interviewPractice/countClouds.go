/*
Challenge: countClouds
https://codefights.com/interview-practice/task/HdgqPhHqs3NciAHqH/description

Given a 2D grid skyMap composed of '1's (clouds) and '0's (clear sky), count the number of clouds. A cloud is surrounded by clear sky, and is formed by connecting adjacent clouds horizontally or vertically. You can assume that all four edges of the skyMap are surrounded by clear sky.

Example

For

skyMap = [['0', '1', '1', '0', '1'],
          ['0', '1', '1', '1', '1'],
          ['0', '0', '0', '0', '1'],
          ['1', '0', '0', '1', '1']]
the output should be
countClouds(skyMap) = 2;

For

skyMap = [['0', '1', '0', '0', '1'],
          ['1', '1', '0', '0', '0'],
          ['0', '0', '1', '0', '1'],
          ['0', '0', '1', '1', '0'],
          ['1', '0', '1', '1', '0']]
the output should be
countClouds(skyMap) = 5.

Input/Output

[time limit] 4000ms (go)
[input] array.array.char skyMap

A 2D grid that represents a map of the sky, as described above.

Guaranteed constraints:

0 ≤ skyMap.length ≤ 300,
0 ≤ skyMap[i].length ≤ 300.

[output] integer

The number of clouds in the given skyMap, as described above.
*/
func countClouds(skyMap [][]string) int {
    if len(skyMap) == 0 || len(skyMap[0]) == 0 {
        return 0
    }

    c := 0
    for i := 0; i < len(skyMap); i++ {
        for j := 0; j < len(skyMap[0]); j++ {
            if skyMap[i][j] == "1" {
                c++
                follow(i, j, skyMap)
            }
        }
    }
    return c
}

func follow(i, j int, skyMap [][]string) {
    if skyMap[i][j] != "1" {
        return
    }
    skyMap[i][j] = "*"
    if i > 0 {
        follow(i-1, j, skyMap)
    }
    if i < len(skyMap)-1 {
        follow(i+1, j, skyMap)
    }
    if j > 0 {
        follow(i, j-1, skyMap)
    }
    if j < len(skyMap[0])-1 {
        follow(i, j+1, skyMap)
    }
}
  