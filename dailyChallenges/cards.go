/**
Challenge: cards
https://codefights.com/challenge/KPYtAdvnN7Pv27TXr

Your task is to write a function that prints out the nth card of the standard card deck. Each of the cards has a rank which can be, in order of ascending value, a number in range from 2 to 10 (depicted as '0'), a Jack ('J'), a Queen ('Q'), a King ('K'), or an Ace ('A'). Each of the cards also has a suit: clubs ('C'), diamonds ('D'), hearts ('H'), or spades ('S').

The cards in the deck are sorted first by their suits in lexicographical order, and then by the value of their ranks. Return the nth (0-based) card in the deck as a string with a length of two, with the first character representing the rank and the second character representing the suit.

Example

For n = 0, the output should be
cards(n) = "2C".

The very first card in the deck is 2 of clubs, so the answer is "2C".

For n = 34, the output should be
cards(n) = "OH".

The 34th card in the deck is 10 of hearts, making the answer "0H".

Input/Output

[time limit] 4000ms (go)
[input] integer n

Guaranteed constraints:
0 ≤ n ≤ 51.

[output] string

A string with a length of 2, representing the nth card in the deck in the format described above.
**/
func cards(n int) string {

    z := "234567890JQKACDHS"
    x := n%13
    y := n/13+13
    return z[x:x+1] + z[y:y+1]
}

