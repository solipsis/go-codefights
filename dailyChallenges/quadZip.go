/**
Challenge: quadZip
https://codefights.com/challenge/z4rytKFWoT7xrJACN

Your task is to implement the quadZip function, which zips four strings. Here's how the function should work: It should take the first character from each string, then the second character from each string, and so on, until no characters are left in the longest string.

Given four strings, aptly named one, two, three, and four, apply the quadZip function to them and return the answer.

Example

For one = "one", two = "two", three = "three" and four = "four",
the output should be
quadZip(one, two, three, four) = "ottfnwhoeoruere".

Here's how the answer can be obtained:

Take the first character of each string: "ottf";
Take the second character of each string: "nwho";
Take the third character of each string: "eoru";
There are no characters left in the first two words, so the fourth letter can be taken from the third and the fourth words: "er";
Now the only characters left are in the third word, so the last character is "e".
Thus, the final answer is "ottfnwhoeoruere".

Input/Output

[time limit] 4000ms (go)
[input] string one

Guaranteed constraints:
1 ≤ one.length ≤ 75.

[input] string two

Guaranteed constraints:
1 ≤ two.length ≤ 75.

[input] string three

Guaranteed constraints:
1 ≤ three.length ≤ 75.

[input] string four

Guaranteed constraints:
1 ≤ four.length ≤ 100.

[output] string

The four words, zipped together.
**/
import "strings"
func quadZip(one string, two string, three string, four string) string {

    one = strings.Replace(one, "\\", "",-1)
    two = strings.Replace(two, "\\", "",-1)
    three = strings.Replace(three, "\\", "",-1)
    four = strings.Replace(four, "\\", "",-1)
    strs := []string{one, two, three, four}
    res := ""
    for i := 0; i <= 100; i++ {
        for _, v := range strs {
            if i < len(v) {
                res += string(v[i])
            }
        }
    }
    
    return res
}
