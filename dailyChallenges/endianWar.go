/**
Challenge: endianWar
https://codefights.com/challenge/ChcFLSa3rfJsKNgkC

The fight on whether to store numbers starting with their most significant digit or their least significant digit has been going on for years. It even got a name and is called the Endian War by some specialists.

Joe Stoy in his (excellent, by the way) book "Denotational Semantics", tells the following story about Alan Turing: "...One early British computer had numbers running from right to left (because the spot on an oscilloscope tube runs from left to right, but in serial logic the least significant digits are dealt with first). Turing used to mystify audiences at public lectures when, quite by accident, he would slip into this mode even for decimal arithmetic, and write things like 73+42=16...".

You are given an expression that was presumably written by Alan Turing. Return true if it is a correct expression written in the little-endian decimal format, or return false otherwise.

Example

For expression = 73+42=16, the output should be
endianWar(expression) = true.

In the little-endian decimal format, the expression becomes 37 + 24 = 61, which is correct.

For expression = "5+8=13", the output should be
endianWar(expression) = false.

In the little-endian decimal format, the result of the expression should be 31.

Input/Output

[time limit] 4000ms (go)
[input] string expression

An expression in the format a+b=c, where a, b and c are decimal numbers (possibly with leading zeros) that are guaranteed to be smaller than 106.

The expression is guaranteed to be valid.

[output] boolean

Return true if expression is valid in the little-endian decimal notation, otherwise return false.
**/
import "regexp"

func endianWar(e string) bool {
    
    r := []rune(e)
    for i,j:=0,len(e)-1;i<len(e)/2;i,j=i+1,j-1 {
        r[i],r[j]=r[j],r[i]
    }
    
    k,_ := regexp.Compile("[+=]")
    
    z := k.Split(string(r), 3)
    c,_ := strconv.Atoi(z[0])
    b,_ := strconv.Atoi(z[1])
    a,_ := strconv.Atoi(z[2])
    return c == b+a
}
