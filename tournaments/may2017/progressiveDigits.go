/**
Challenge: progressiveDigits
https://codefights.com/tournaments/vbZdzDTcajNBS7xN6/G

Finite arithmetic progressions by modulo 10 are almost like ordinary arithmetic progressions, but with each element taken by modulo 10. For example, the arithmetic progression 5, 7, 9, 11, 13, 15 turns into 5, 7, 8, 1, 3, 5 when each element is taken by modulo 10. These progressions can be converted to numbers, using the individual elements as digits. Doing so, we get the number 579135 from the progression 5, 7, 9, 1, 3, 5.

Given two integers, l and r, return all the numbers in the segment [l, r] in ascending order, such that each of them has been converted to a number from the finite arithmetic progression by modulo 10.

Example

For l = 100 and r = 200, the output should be
progressiveDigits(l, r) = [109, 111, 123, 135, 147, 159, 161, 173, 185, 197].

All the progressions here have an initial number 1. Let's look more closely at the first three numbers:

109 was taken from the finite arithmetic progression by modulo 10 1, 0, 9, which corresponds to the arithmetic progression 1, 10, 19 (a common difference of 9).

111 was taken from the arithmetic progression 1, 1, 1 (a common difference of 0).

123 was taken from the arithmetic progression 1, 2, 3 (a common difference of 1).

Input/Output

[time limit] 4000ms (go)
[input] integer64 l

The left point of a segment.

Guaranteed constraints:
1 ≤ l ≤ r ≤ 1015.

[input] integer64 r

The right point of a segment.

Guaranteed constraints:
1 ≤ l ≤ r ≤ 1015.

[output] array.integer64

All the numbers in the segment [l, r] that are taken from some finite arithmetic progression by modulo 10, returned in ascending order.
**/
import "math"
import "sort"

func progressiveDigits(l int64, r int64) []int64 {

    minDigits := math.Log10(float64(l)) + 1
    maxDigits := math.Log10(float64(r)) + 1
    
    m := make(map[int]bool)
    
    fmt.Printf("min: %v max: %v\n", minDigits, maxDigits)
    
    for i := 1; i <= 9; i++ {
        s := strconv.Itoa(i)
        for j := 0; j <= 9; j++ {
            temp := s
            val := i
            if int64(val) >= l && int64(val) <= r {
                m[val] = true
            }
            
            cur, _ := strconv.Atoi(temp)
    
            for int64(cur) < r {
                val = (val + j) % 10
                temp = temp + strconv.Itoa(val)
                cur, _ = strconv.Atoi(temp)
                if int64(cur) >= l && int64(cur) <= r {
                    m[cur] = true
                }
            }
        }
    }
    
    sorted := make([]int, len(m))
    ret := make([]int64, len(m))
    x := 0
    for k := range m {
        sorted[x] = k
        x++
    }
    fmt.Println(m)
    sort.Ints(sorted)
    for n, v := range sorted {
        ret[n] = int64(v)
    }
    
    return ret
}
