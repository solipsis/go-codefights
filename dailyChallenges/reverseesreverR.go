/**
Challenge: ReversesreveR
https://codefights.com/challenge/WeEiBDEkJtJhsRfZu

This is a reverse challenge, enjoy!

Input/Output

[time limit] 4000ms (go)
[input] string inputString

A valid (in terms of this challenge) mathematical expression in the format <a><op><b>, where <a> and <b> are integers not greater than 1000, and <op> is one of the following operations: +, -, /, *, %, ^.

[output] integer
**/
import "regexp"
import "math"

func ReversesreveR(i string) int {
    g, _ := regexp.Compile("[+-/\\*%^]")
    z := g.Split(i, -1)
   
    a, _ := strconv.Atoi(r(z[0]))
    b, _ := strconv.Atoi(r(z[1]))
    
    switch p := string(i[len(z[0])]); p {
        case "+":
        return a + b
        case "-":
        return a - b
        case "/":
        return a / b
        case "*":
        return a * b
        case "%":
        return a % b
        case "^":
        return int(math.Pow(float64(a), float64(b))) 
    }
    
    return 0
}

func r(a string) string {
    r := []rune(a)
    for i,j := 0, len(a)-1; i < len(a)/2; i, j = i+1,j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}