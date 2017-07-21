/**
Challenge: Star Rotation
https://codefights.com/arcade/code-arcade/list-backwoods/A5meC5REAcDfHeuFf/description

Consider a (2k+1) × (2k+1) square subarray of an integer integers matrix. Let's call the union of the square's two longest diagonals, middle column and middle row a star. Given the coordinates of the star's center in the matrix and its width, rotate it 45 · t degrees clockwise preserving position of all matrix elements that do not belong to the star.

Example

For

matrix = [[1, 0, 0, 2, 0, 0, 3],
          [0, 1, 0, 2, 0, 3, 0],
          [0, 0, 1, 2, 3, 0, 0],
          [8, 8, 8, 9, 4, 4, 4],
          [0, 0, 7, 6, 5, 0, 0],
          [0, 7, 0, 6, 0, 5, 0],
          [7, 0, 0, 6, 0, 0, 5]]
width = 7, center = [3, 3] and t = 1, the output should be

starRotation(matrix, width, center, t) = [[8, 0, 0, 1, 0, 0, 2],
                                          [0, 8, 0, 1, 0, 2, 0],
                                          [0, 0, 8, 1, 2, 0, 0],
                                          [7, 7, 7, 9, 3, 3, 3],
                                          [0, 0, 6, 5, 4, 0, 0],
                                          [0, 6, 0, 5, 0, 4, 0],
                                          [6, 0, 0, 5, 0, 0, 4]]
For

matrix = [[1, 0, 0, 2, 0, 0, 3],
          [0, 1, 0, 2, 0, 3, 0],
          [0, 0, 1, 2, 3, 0, 0],
          [8, 8, 8, 9, 4, 4, 4],
          [0, 0, 7, 6, 5, 0, 0]]
width = 3, center = [1, 5] and t = 81, the output should be

starRotation(matrix, width, center, t) = [[1, 0, 0, 2, 0, 0, 0],
                                          [0, 1, 0, 2, 3, 3, 3],
                                          [0, 0, 1, 2, 0, 0, 0],
                                          [8, 8, 8, 9, 4, 4, 4],
                                          [0, 0, 7, 6, 5, 0, 0]]
Input/Output

[time limit] 4000ms (go)
[input] array.array.integer matrix

A two-dimensional array of integers.

Guaranteed constraints:
3 ≤ matrix.length ≤ 15,
3 ≤ matrix[i].length ≤ 15,
matrix[i].length == matrix[j].length,
0 ≤ matrix[i][j] ≤ 99.

[input] integer width

An odd integer representing the star's width. It equals the length of the sides of the bounding square for the star.

Guaranteed constraints:
3 ≤ width ≤ min(matrix.length, matrix[i].length).

[input] array.integer center

An array of two integers.

Guaranteed constraints:
center.length = 2,
(width - 1) / 2 ≤ center[0] < matrix.length - (width - 1) / 2,
(width - 1) / 2 ≤ center[1] < matrix[i].length - (width - 1) / 2.

[input] integer t

A non-negative integer denoting how many times the star should be rotated by 45 degrees.

Guaranteed constraints:
0 ≤ t ≤ 109.

[output] array.array.integer

An array with specified star rotated by 45 · t degrees.
**/
func starRotation(star [][]int, width int, center []int, t int) [][]int {

    zero := func(x int) int { return x}
    add := func(x int) int { return x + 1}
    sub := func(x int) int { return x - 1}
    
    // parse out each line of star
    arr := [][]int {
        parse(sub, zero, width, center, star),
        parse(sub, add, width, center, star),
        parse(zero, add, width, center, star),
        parse(add, add, width, center, star),
        parse(add, zero, width, center, star),
        parse(add, sub, width, center, star),
        parse(zero, sub, width, center, star),
        parse(sub, sub, width, center, star),
    }
    
    // rotate the star
    times := 8 - (t % 8)
    reverse(arr[:times])
    reverse(arr[times:])
    reverse(arr)
    
    // overrite the star
    write(sub, zero, width, center, arr[0], star)
    write(sub, add, width, center, arr[1], star)
    write(zero, add, width, center, arr[2], star)
    write(add, add, width, center, arr[3], star)
    write(add, zero, width, center, arr[4], star)
    write(add, sub, width, center, arr[5], star)
    write(zero, sub, width, center, arr[6], star)
    write(sub, sub, width, center, arr[7], star)
    
    return star
}

func write(r, c func(int) int, width int, center, arr []int, star [][]int) {
    i, j := center[0], center[1]
    for z := 0; z < width / 2; z++ {
        i = r(i)
        j = c(j)
        star[i][j] = arr[z]
    }
}

func parse(r, c func(int) int, width int, center []int, star [][]int) []int {
    arr := make([]int, 0)
    i, j := center[0], center[1]
    for z := 0; z < width / 2; z++ {
        i = r(i)
        j = c(j)
        arr = append(arr, star[i][j])
    }
    return arr
}

func reverse(arr [][]int) {
    for i,j := 0, len(arr)-1; i < j; i,j=i+1,j-1 {
        arr[i], arr[j] = arr[j], arr[i]
    }
}
