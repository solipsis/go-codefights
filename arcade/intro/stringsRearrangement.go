/*
Challenge: strings rearrangement 
https://codefights.com/arcade/intro/level-7/PTWhv2oWqd6p4AHB9/description

Given an array of equal-length strings, check if it is possible to rearrange the strings in such a way that after the rearrangement the strings at consecutive positions would differ by exactly one character.

Example

For inputArray = ["aba", "bbb", "bab"], the output should be
stringsRearrangement(inputArray) = false;

All rearrangements don't satisfy the description condition.

For inputArray = ["ab", "bb", "aa"], the output should be
stringsRearrangement(inputArray) = true.

Strings can be rearranged in the following way: "aa", "ab", "bb".
Input/Output

[time limit] 4000ms (go)
[input] array.string inputArray

A non-empty array of strings of lowercase letters.

Guaranteed constraints:
2 ≤ inputArray.length ≤ 10,
1 ≤ inputArray[i].length ≤ 15.

[output] boolean
*/
var perms [][]string
func stringsRearrangement(inputArray []string) bool {
  
    perms = make([][]string, 0)
    n := len(inputArray)
    
    permute(inputArray, 0, n-1)
    for _, perm := range perms {
        fail := false
        prev := perm[0]
        for _, str := range perm[1:] {
            if !oneDiff(prev, str) {
                fail = true
            }
            prev = str
        }
        if !fail { 
            return true
        }
    }
    return false
}


func permute(arr []string, l, r int) {
    if l == r {
        cp := make([]string, len(arr))
        copy(cp, arr)
        perms = append(perms, cp)
        return 
    }
    for i := l;i <= r; i++ {
        swap(arr, l, i)
        permute(arr, l+1, r)
        swap(arr, l, i)
    } 
}

func swap(arr []string, i, j int) {
    arr[i], arr[j] = arr[j], arr[i]
}

func oneDiff(a, b string) bool {
    dif := 0
    for n, _ := range a {
        if a[n] != b[n] {
            dif++
        }
    }
    return dif == 1
}

