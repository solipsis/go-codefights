/**
Challenge: alienBit
https://codefights.com/challenge/aEPJbWzkSjEuuH3cH

Earth-Alien Transmissions Services (EATS) have intercepted some transmissions from a nearby star. It appears that aliens are trying to communicate with us, but the transmissions are just long strings of numbers with a decimal place at the beginning.

You, a genius cryptographer and linguist, managed to figure out what these numbers mean. It looks like the aliens just have a different way of computing! Rather than having just two states like we do (1 and 0), the alien computers measure the intensity of a single particle with an accuracy that allows reading nearly unlimited states between 0 and 1. This means that a long string's worth of data can be stored in one alien "bit'.

Conveniently, the aliens also use ASCII - what luck! Now all you need to do is write a program that turns an alien bit abit into a translated string.

Example

For abit = "0.116101115116035049", the output should be
alienBit(abit) = "test#1".

The alien bit can be split into 6 ASCII symbols:

116 -> 't'
101 -> 'e'
115 -> 's'
116 -> 't'
35 -> '#'
49 -> '1'
Thus, this alien transmission contains the message "test#1".

Input/Output

[time limit] 4000ms (go)
[input] string abit

A string that represents a number between 0 and 1, with the number of decimal places divisible by 3.

Guaranteed constraints:
5 ≤ abit.length ≤ 1001.

[output] string

A translated string of ASCII characters.
**/
func alienBit(a string) (r string) {
    
    for i := 2; i < len(a); i+=3 {
        i, _ := strconv.Atoi(a[i:i+3])
        r += string(i)
    }
    return r  
}
