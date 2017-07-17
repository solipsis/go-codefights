/**
Challenge: reverseInverse
https://codefights.com/challenge/bzkdx7vP3QJ8DgQTL

Given a string s, your mission is to apply the following easy-to-comprehend algorithm to it:

Find all the words in s, where a word is a sequence of consecutive alphanumeric characters with no other letters around it;
Reverse the characters in each word;
For each word, swap the cases of its characters so that the case of a character at each position differs from the case at the corresponding position of the original (unreversed) word.
Return the obtained string as the answer.

Example

For s = "So, what is CodeFights?", the answer should be
reverseInverse(s) = "oS, TAHW SI sTHGiFEDOC?".

There are 4 words in s: "So", "what", "is", and "CodeFights". Let's take the word "CodeFights" as an example:

The letters 'C' at index 0 and letter 'F' at index 4 are uppercase, while all the other letters are lowercase;
"codefights" reversed becomes "sthgifedoc";
With the cases swapped, the letters at indices 0 and 4 should be lowercase and all the other letters should be uppercase;
Thus, the final word is "sTHGiFEDOC".
Input/Output

[time limit] 4000ms (go)
[input] string s

A string containing only alphanumeric characters and punctuation marks.

Guaranteed constraints:
0 ≤ s.length < 500.

[output] string

The result of applying the algorithm described above to s.
**/
func reverseInverse(s string) string {
    
    w := ""
    res := ""
    caps := make([]int, len(s))
    for n, r := range s {
        if r >= 65 && r <= 90 {
            w += string(r)
            caps[n] = 1
        } else if r >= 97 && r <= 122 {
            w += string(r)
            caps[n] = 2
        } else if r >= 48 && r <= 57 {
            w += string(r)
        } else {
            res += reverse(w) + string(r)
            w = ""
        }
    }
    res += reverse(w)
    z := ""
    for n, r := range res {
        if caps[n] == 1 && r < 97 {
            z += string(r+32)
        } else if caps[n] == 2 && r > 96{
            z += string(r-32)
        } else {
            z += string(r)
        }
    }
    return z
}

func reverse(s string) string {
    r := []rune(s)
    for i,j:=0,len(s)-1; i < len(s)/2; i,j=i+1,j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}

