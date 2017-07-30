/**
Challenge: Digit Difference Sort
https://codefights.com/arcade/code-arcade/sorting-outpost/2SFFWqkhkqC7mMBse/description


Given an array of integers, sort its elements by the difference of their largest and smallest digits. In the case of a tie, that with the larger index in the array should come first.

Example

For a = [152, 23, 7, 887, 243], the output should be
digitDifferenceSort(a) = [7, 887, 23, 243, 152].

Here are the differences of all the numbers:

152: difference = 5 - 1 = 4;
23: difference = 3 - 2 = 1;
7: difference = 7 - 7 = 0;
887: difference = 8 - 7 = 1;
243: difference = 4 - 2 = 2.
23 and 887 have the same difference, but 887 goes after 23 in a, so in the sorted array it comes first.

Input/Output

[time limit] 4000ms (go)
[input] array.integer a

An array of integers.

Guaranteed constraints:
0 ≤ sequence.length ≤ 104,
1 ≤ sequence[i] ≤ 105.

[output] array.integer

Array a sorted as described above.
**/
import "strings"
import "sort"
func digitDifferenceSort(a []int) []int {

    items := make(Items, 0)
    for n, v := range a {
        str := strconv.Itoa(v)
        arr := strings.Split(str, "")
        min, max := 999999, 0
        for _, s := range arr {
            i, _ := strconv.Atoi(s)
            if i < min {
                min = i
            }
            if i > max {
                max = i
            }
        }
        items = append(items, item{v, max - min, n})
    }
    sort.Sort(items)
    for n := range a {
        a[n] = items[n].val
    }
    
    return a
}

type item struct {
    val, diff, idx int
}

type Items []item
func (it Items) Len() int { return len(it) }
func (it Items) Swap(i, j int) { it[i], it[j] = it[j], it[i] }
func (it Items) Less(i, j int) bool { 
    if it[i].diff < it[j].diff {
        return true
    } else if it[i].diff > it[j].diff {
        return false
    }
    return it[i].idx > it[j].idx
}