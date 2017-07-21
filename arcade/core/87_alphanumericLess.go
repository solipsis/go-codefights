/**
Challenge: Alphanumeric Less
https://codefights.com/arcade/code-arcade/lab-of-transformations/bBgu7NEfk2GoQuNzf/description

An alphanumeric ordering of strings is defined as follows: each string is considered as a sequence of tokens, where each token is a letter or a number (as opposed to an isolated digit, as is the case of lexicographic ordering). For example, the tokens of a string "ab01c004" are [a, b, 01, c, 004]. In order to compare two strings, you break them down into tokens and compare corresponding pairs of tokens with each other (i.e. first token of the first string, with the first token of the second string etc).

Here is how tokens are compared:

If a letter is compared with another letter, the usual order applies.
A number is always less than a letter.
When two numbers are compared, their values are compared. Leading zeros, if any, are ignored.
If at some point one string has no more tokens left while the other one still does, the one with fewer tokens is considered smaller.

If the two strings s1 and s2 appear to be equal, consider the smallest index i such that tokens(s1)[i] and tokens(s2)[i] (where tokens(s)[i] is the ith token of string s) differ only by the number of leading zeros. If no such i exists, the strings are indeed equal. Otherwise, the string whose ith token has more leading zeros is considered less.

Here are some examples of comparing strings using alphanumeric ordering.

"a" < "a1" < "ab"
"ab42" < "ab000144" < "ab00144" < "ab144" < "ab000144x"
"x11y012" < "x011y13"
Example

For s1 = "a" and s2 = "a1", the output should be
alphanumericLess(s1, s2) = true;
For s1 = "ab" and s2 = "a1", the output should be
alphanumericLess(s1, s2) = false;
For s1 = "b" and s2 = "a1", the output should be
alphanumericLess(s1, s2) = false.
Input/Output

[time limit] 4000ms (go)
[input] string s1

String, consisting of Latin letters and digits.

Guaranteed constraints:
1 ≤ s1.length ≤ 20.

[input] string s2

String, consisting of Latin letters and digits.

Guaranteed constraints:
1 ≤ s2.length ≤ 20.

[output] boolean

true if s1 is alphanumerically strictly less than s2, false otherwise.
**/
import "math/big"
func alphanumericLess(s1 string, s2 string) bool {
    
    // very ugly should revisit this one
    a := tokenize(s1)
    b := tokenize(s2)
    firstNumTie := 0
    for n, _ := range a {
        if n >= len(b) {
            return false
        }
        
        n1 := false
        n2 := false
        if a[n][0] < 65 {
            n1 = true
        }
        if b[n][0] < 65 {
            n2 = true
        }
        if n1 && !n2 {
            return true
        }
        if !n1 && n2 {
            return false
        }
        if n1 && n2 {
            // one edge case overflows int64 :(
            i1 := big.NewInt(0)
            i2 := big.NewInt(0)
            i1, _ = i1.SetString(a[n], 10)
            i2, _ = i2.SetString(b[n], 10)
            if i1.Cmp(i2) < 0 {
                return true
            }
            if i1.Cmp(i2) > 0 {
                return false
            }
            if firstNumTie == 0 {
                if len(a[n]) > len(b[n]) {
                    firstNumTie = -1
                }
                if len(a[n]) < len(b[n]) {
                    firstNumTie = 1
                }
            }
        }
        if !n1 && !n2 {
            if a[n] < b[n] {
                return true
            }
            if a[n] > b[n] {
                return false
            }
        }
        
    }
    if len(a) < len(b) {
        return true
    }
    if len(a) > len(b) {
        return false
    }
    if firstNumTie < 0 {
        return true
    }   
    return false
}

func tokenize(s string) []string {
    tokens := make([]string, 0)
    t := s[0]
    s = s[1:]
    tok := "" + string(t)
    num := t < 65
    if !num {
        tokens = append(tokens, tok)
        tok = ""
    }
    for n := range s {
        if s[n] < 65 {   
            tok += string(s[n])
            num = true
        } else {
            if num {
                tokens = append(tokens, tok)       
            }
            tokens = append(tokens, string(s[n]))
            tok = ""
            num = false
        }
    }
    if len(tok) > 0 {
        tokens = append(tokens, tok)
    }
    
    return tokens
}
