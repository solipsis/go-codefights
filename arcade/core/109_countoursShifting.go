/**
Challenge: Contours Shifting
https://codefights.com/arcade/code-arcade/waterfall-of-integration/LfP67YRDw2rxoLmeP/description

Mark got a rectangular array matrix for his birthday, and now he's thinking about all the fun things he can do with it. He likes shifting a lot, so he decides to shift all of its i-contours in a clockwise direction if i is even, and counterclockwise if i is odd.

Here is how Mark defines i-contours:

the 0-contour of a rectangular array as the union of left and right columns as well as top and bottom rows;
consider the initial matrix without the 0-contour: its 0-contour is the 1-contour of the initial matrix;
define 2-contour, 3-contour, etc. in the same manner by removing 0-contours from the obtained arrays.
Implement a function that does exactly what Mark wants to do to his matrix.

Example

For

matrix = [[ 1,  2,  3,  4],
           [ 5,  6,  7,  8],
           [ 9, 10, 11, 12],
           [13, 14, 15, 16],
           [17, 18, 19, 20]]
the output should be

contoursShifting(matrix) = [[ 5,  1,  2,  3],
                             [ 9,  7, 11,  4],
                             [13,  6, 15,  8],
                             [17, 10, 14, 12],
                             [18, 19, 20, 16]]
For matrix = [[238, 239, 240, 241, 242, 243, 244, 245]],
the output should be
contoursShifting(matrix) = [[245, 238, 239, 240, 241, 242, 243, 244]].

Note, that if a contour is represented by a 1 × n array, its center is considered to be below it.

For

matrix = [[238],
           [239],
           [240],
           [241],
           [242],
           [243],
           [244],
           [245]]
the output should be

contoursShifting(matrix) = [[245],
                             [238],
                             [239],
                             [240],
                             [241],
                             [242],
                             [243],
                             [244]]
If a contour is represented by an n × 1 array, its center is considered to be to the left of it.

Input/Output

[time limit] 4000ms (go)
[input] array.array.integer matrix

Guaranteed constraints:
1 ≤ matrix.length ≤ 100,
1 ≤ matrix[i].length ≤ 100,
1 ≤ matrix[i][j] ≤ 1000.

[output] array.array.integer

The given matrix with its contours shifted.
**/
func contoursShifting(matrix [][]int) [][]int {
    r := make([][]int, len(matrix))
    for n, _ := range r {
        r[n] = make([]int, len(matrix[0]))
    }
    shortestDim := len(matrix)
    if shortestDim > len(matrix[0]) {
        shortestDim = len(matrix[0])
    }
    
    nContours := shortestDim / 2
    if shortestDim % 2 == 1 {
        nContours = shortestDim / 2 +1
    } 
    
    for x := 0; x < nContours; x++ {
        b := x % 2 == 0
        write(r, shift(extract(matrix, x), b), x)
    }
    
    return r
}

func shift(arr []int, right bool) []int {
    i := 1
    if right {
        i = len(arr)-1
    }
    
    if len(arr) == 1 {
        return arr
    }
    reverse(arr[:i])
    reverse(arr[i:])
    reverse(arr)
    return arr
}

func reverse(arr []int) {
    for i,j := 0, len(arr)-1; i < len(arr)/2; i,j = i+1,j-1 {
        arr[i], arr[j] = arr[j], arr[i]
    }
}

func write(res [][]int, shifted []int, layer int) [][]int {
    i := 0
    if (len(res) == 1) {
        res[0] = shifted
        return res
    }
    if len(res[0]) == 1 {
        for _, v := range res {
            v[0] = shifted[i]
            i++
        }
        return res
    }
    for x := layer; x < len(res[0])-layer; x++ {
        res[layer][x] = shifted[i]
        i++
    }
    for y := layer+1; y < len(res)-layer; y++ {
        res[y][len(res[0])-layer-1] = shifted[i]
        i++
    }
    if (layer == len(res)-1-layer || layer == len(res[0])-layer-1) {
        return res
    }
    for x := len(res[0])-2-layer; x >= layer; x-- {
        res[len(res)-layer-1][x] = shifted[i]
        i++
    }
    for y := len(res)-2-layer; y >= layer+1; y-- {
        res[y][layer] = shifted[i]
        i++
    }
    return res
}

func extract(contour [][]int, layer int) []int {
    
    if len(contour) == 1 {
        return contour[0]
    }
    
    r := []int{}
    if len(contour[0]) == 1 {
        for _, v := range contour {
            r = append(r, v[0])
        }
        return r
    }
    
    for x := layer; x < len(contour[0])-layer; x++ {
        r = append(r, contour[layer][x])
    }
    for y := layer+1; y < len(contour)-layer; y++ {
        r = append(r, contour[y][len(contour[0])-layer-1])
    }
    if (layer == len(contour)-1-layer || layer == len(contour[0])-layer-1) {
        return r
    }    
    for x := len(contour[0])-2-layer; x >= layer; x-- {
        r = append(r, contour[len(contour)-layer-1][x])
    }
    for y := len(contour)-2-layer; y >= layer+1; y-- {
        r = append(r, contour[y][layer])
    }
    return r
}
