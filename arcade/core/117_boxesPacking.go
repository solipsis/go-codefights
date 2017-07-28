/**
Challenge: Boxes Packing
https://codefights.com/arcade/code-arcade/sorting-outpost/9y4wLpcqnNozn92tG/description


You are given n rectangular boxes, the ith box has the length lengthi, the width widthi and the height heighti. Your task is to check if it is possible to pack all boxes into one so that inside each box there is no more than one another box (which, in turn, can contain at most one another box, and so on). More formally, your task is to check whether there is such sequence of n different numbers pi (1 ≤ pi ≤ n) that for each 1 ≤ i < n the box number pi can be put into the box number pi+1.
A box can be put into another box if all sides of the first one are less than the respective ones of the second one. You can rotate each box as you wish, i.e. you can swap its length, width and height if necessary.

Example

For length = [1, 3, 2], width = [1, 3, 2] and height = [1, 3, 2], the output should be
boxesPacking(length, width, height) = true;
For length = [1, 1], width = [1, 1] and height = [1, 1], the output should be
boxesPacking(length, width, height) = false;
For length = [3, 1, 2], width = [3, 1, 2] and height = [3, 2, 1], the output should be
boxesPacking(length, width, height) = false.
Input/Output

[time limit] 4000ms (go)
[input] array.integer length

Array of positive integers.

Guaranteed constraints:
1 ≤ length.length ≤ 104,
1 ≤ length[i] ≤ 2 · 104.

[input] array.integer width

Array of positive integers.

Guaranteed constraints:
width.length = length.length,
1 ≤ width[i] ≤ 2 · 104.

[input] array.integer height

Array of positive integers.

Guaranteed constraints:
height.length = length.length,
1 ≤ height[i] ≤ 2 · 104.

[output] boolean

true if it is possible to put all boxes into one, false otherwise.
**/
import "sort"
func boxesPacking(length []int, width []int, height []int) bool {

    boxes := make(boxes, 0)
    for i := 0; i < len(length); i++ {
        arr := []int{length[i], width[i], height[i]}
        sort.Ints(arr)
        boxes = append(boxes, arr)
    }
    sort.Sort(sort.Reverse(boxes))
    for i := 0; i < len(boxes)-1; i++ {
        b1, b2 := boxes[i], boxes[i+1]
        for j := 0; j < len(b1); j++ {
            if b1[j] <= b2[j] {
                return false
            }
        }
    }
    return true
}

type boxes [][]int
func (b boxes) Len() int { return len(b) }
func (b boxes) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b boxes) Less(i, j int) bool {
    x, y := b[i], b[j]
    for z := 0; z < len(x); z++ {
        if x[z] < y[z] {
            return true
        } else if x[z] > y[z] {
            return false
        }
    }
    return false
}
