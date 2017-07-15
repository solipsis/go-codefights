/*
Challenge: Is Lucky
https://codefights.com/arcade/intro/level-3/3AdBC97QNuhF6RwsQ/description

Ticket numbers usually consist of an even number of digits. A ticket number is considered lucky if the sum of the first half of the digits is equal to the sum of the second half.

Given a ticket number n, determine if it's lucky or not.

Example

For n = 1230, the output should be
isLucky(n) = true;
For n = 239017, the output should be
isLucky(n) = false.
Input/Output

[time limit] 4000ms (go)
[input] integer n

A ticket number represented as a positive integer with an even number of digits.

Guaranteed constraints:
10 ≤ n < 106.

[output] boolean

true if n is a lucky ticket number, false otherwise.
*/

func isLucky(n int) bool {

    s := strconv.Itoa(n)
    var a, b uint8
    for _, r := range s[0:len(s)/2] {
        a += uint8(r) - 48
    }
    for _, r := range s[len(s)/2:] {
        b += uint8(r) - 48
    }
  
    return a == b
}
