/**
Challenge: convertToOneCase
https://codefights.com/tournaments/vbZdzDTcajNBS7xN6/A

Alex often finds words that have letters both in upper and lower case, and they make him feel terrible. He's so tired of this that he decides to convert all these words to one letter case.

To choose the case, he calculates the number of lower case letters and the number of upper case letters. If these numbers have the same parity, Alex converts the word to lower case. Otherwise, he converts the word to upper case. Note that Alex doesn't convert words that only have letters that are one case.

Given the word that Alex wants to convert, return the modified word.

Example

For word = "KeY", the output should be
convertToOneCase(word) = "KEY".

This word contains 2 upper case letters and 1 lower case letter. 2 and 1 have the opposite parity, so this word should be converted to upper case.

For word = "FOObar", the output should be
convertToOneCase(word) = "foobar".

This word contains 3 letters in both upper and lower case. 3 and 3 have the same parity, so this word should be converted to lower case.

For word = "chamomile", the output should be
convertToOneCase(word) = "chamomile".

The letters in this word are only in one case, so Alex doesn't convert it.

Input/Output

[time limit] 4000ms (go)
[input] string word

The word that Alex wants to convert. It contains only English letters.

Guaranteed constraints:
1 ≤ word.length ≤ 1000.

[output] string

The modified word.
**/
import "strings"

func convertToOneCase(word string) string {

    l := 0
    u := 0
    for _, v := range word {
        if v >= 97 {
            l++
        } else {
            u++
        }
    }
    if (u == len(word) || l == len(word)) {
        return word
    }
    
    fmt.Println(l)
    fmt.Println(u)
    if l % 2 == u % 2 {
        return strings.ToLower(word)
    }
    return strings.ToUpper(word)
}
