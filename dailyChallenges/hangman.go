/**
Challenge: hangman
https://codefights.com/challenge/TuFtqDrA3bcpeW9Ed/solutions/2a8DBX2kHbLZtezdC


It is a common tradition to play Hangman game during classes. Too bad it's difficult to do so if you and your friend sit far from each other. You, however, came up with a great idea.

First, you write down a word. After that your friend writes down distinct letters and passes the list with them to you.
Now you look at the letters in his list one by one. If the current letter is present in your word, you erase all occurrences of this letter from it, otherwise you call it a miss. If you erase the entire word before 6 misses, then your friend wins. If you count 6 misses or run out of letters before the word is erased, you win.

Given the word you made and your friend's letters, return true if he wins or false otherwise.

Example

For word = "hello" and letters = "aenmrolhtg", the output should be
hangman(word, letters) = true.

The stages of the game are:

'a' - 1st miss;
'e' - correct letter (_ e _ _ _);
'n' - 2nd miss;
'm' - 3rd miss;
'r' - 4th miss;
'o' - correct letter (_ e _ _ o);
'l' - correct letter (_ e l l o);
'h' - correct letter (h e l l o);
Other letters don't matter since the word is guessed already.

Input/Output

[time limit] 4000ms (go)
[input] string word

A non-empty word composed of lowercase Latin letters.

Guaranteed constraints:
1 ≤ word.length ≤ 15.

[input] string letters

A string of distinct lowercase Latin letters.

Guaranteed constraints:
5 ≤ letters.length ≤ 15.

[output] boolean

true if the word is guessed, false otherwise.
**/
import "strings"
func hangman(w, l string) bool {
    // see how many matches we have until we have len(w) matches or 6 misses
    m := 6
    g := len(w)
    for _, r := range l {   
        n := g
        g -= len(strings.Split(w, string(r))) - 1
        if n == g {
            m--
        }
        
        if m < 1 {
            break
        }
    }
    return g < 1
}
