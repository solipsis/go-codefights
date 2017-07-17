/*
Challenge: rotateImage
https://codefights.com/interview-practice/task/5A8jwLGcEpTPyyjTB/description

Note: Try to solve this task in-place (with O(1) additional memory), since this is what you'll be asked to do during an interview.

You are given an n x n 2D matrix that represents an image. Rotate the image by 90 degrees (clockwise).

Example

For

a = [[1, 2, 3],
     [4, 5, 6],
     [7, 8, 9]]
the output should be

rotateImage(a) =
    [[7, 4, 1],
     [8, 5, 2],
     [9, 6, 3]]
Input/Output

[time limit] 4000ms (go)
[input] array.array.integer a

Guaranteed constraints:
1 ≤ a.length ≤ 100,
a[i].length = a.length,
1 ≤ a[i][j] ≤ 104.

[output] array.array.integer
*/
func rotateImage(a [][]int) [][]int {

    // to rotate 90 degrees. transpose matrix than reverse rows
    for r := 0; r < len(a)-1; r++ {
        for c := r+1; c < len(a[0]); c++ {
            fmt.Printf("r: %v c: %v\n", r, c)
            a[r][c], a[c][r] = a[c][r], a[r][c]
        }
    }
    for _, row := range a {
        reverse(row)
    }
    
    return a
}
func reverse(r []int) {
    for i,j := 0, len(r)-1; i < j; i,j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
}
