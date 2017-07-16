/*
Challenge: longestWord
https://codefights.com/arcade/intro/level-12/s9qvXv4yTaWg8g4ma/description

Define a word as a sequence of consecutive English letters. Find the longest word from the given string.

Example

For text = "Ready, steady, go!", the output should be
longestWord(text) = "steady".

Input/Output

[time limit] 4000ms (go)
[input] string text

Guaranteed constraints:
4 ≤ text.length ≤ 50.

[output] string

The longest word from text. It's guaranteed that there is a unique output.
*/
import "regexp"

func longestWord(text string) string {

    //var words [][]byte
    max := 0
    maxW := ""
    re := regexp.MustCompile("[A-Za-z]+")
    words := re.FindAllString(text, -1)
    for _, word := range words {
        if len(word) > max {
            max = len(word)
            maxW = word
        }
    }
    
    return maxW
}
