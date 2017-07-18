/**
Challenge: wordsRotationCc
https://codefights.com/challenge/c6TgF2asHrFEN8X3Z

Given a matrix of words of the same length, rotate it 90° in the counterclockwise direction.

Input/Output

For

words = ["apple", 
         "anger", 
         "monks", 
         "stink"]
the output should be

wordsRotationCc(words) = ["ersk", 
                          "lekn", 
                          "pgni", 
                          "pnot", 
                          "aams"]
Input/Output

[time limit] 4000ms (go)
[input] array.string words

An array of strings of equal length.

Guaranteed constraints:
1 ≤ words.length ≤ 20,
1 ≤ words[i].length ≤ 20.

[output] array.string

The words matrix, rotated counterclockwise 90°.
**/
func wordsRotationCc(w []string) (a []string) {

    for c := len(w[0])-1; c >= 0; c-- {
        s := ""
        for r := 0; r < len(w); r++ {
            s += string(w[r][c])
        }
        a = append(a, s)
    }
    return
}
