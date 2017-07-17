/**
Challenge: incrementCipher
https://codefights.com/challenge/ea9kjzNzE6dQGw8MY

You are given a string s, which consists only of lowercase English letters, and an array of numbers, nums. For each ni = nums[i], increment every nith letter in the given string - i.e. the letters with indices s[ni], s[2 · ni], s[3 · ni] and so on (1-based). For letters, incrementing means changing the letter to the next one in the alphabet. For example, a becomes b, b becomes c, and z cycles to a.

Example

For s = "abc" and nums = [1, 2, 3], the output should be
incrementCipher(s, nums) = "bde".

First, we should increment every 1st letter, so we just increment every letter and get s = "bcd". Then we increment every 2nd letter and get s = "bdd" (since the length of s is 3, there is only one such letter). And then, after incrementing every 3rd letter, we get the answer: s = "bde".

Input/Output

[time limit] 4000ms (go)
[input] string s

A string composed of lowercase English letters.

Guaranteed constraints:
1 ≤ s.length ≤ 100 .

[input] array.integer nums

Guaranteed constraints:
1 ≤ nums.length ≤ 50 ,
1 ≤ nums[i] ≤ 50 .

[output] string

The newly ciphered string.
**/
func incrementCipher(s string, z []int) string {

    r := []byte(s)
    for _, v := range z {
        for n := range r {
            if (n+1) % v < 1 {
                r[n]++
                if r[n] > 122 {
                    r[n] = 97
                } 
            }
        }
        
    }
    
    return string(r)
}
