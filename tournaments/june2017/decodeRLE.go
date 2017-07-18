/**
Challenge: decodeRLE
https://codefights.com/tournaments/ZS9S4h3HFKsuWRMnD/B

You like learning about different encoding algorithms. Yesterday you read about run-length encoding (RLE), which works as follows: Run the string data and change all substrings of consecutive equal characters to this character and its number of occurrences. For example, "aabccc" will be changed to "a2b1c3" and "aaaaabcabb" to "a5b1c1a1b2". But just encoding strings is too boring for you! So you decided to decode strings that were encoded using RLE. Given a string code, which is the result of RLE, determine the length of the initial string.

Example

For code = "a2b1c3", the output should be
decodeRLE(code) = 6;
For code = "a5b1c1a1b2", the output should be
decodeRLE(code) = 10.
Input/Output

[time limit] 4000ms (go)
[input] string code

A string, containing lowercase English letters and digits. It is guaranteed that this string was obtained using RLE for string, consisting only from small English letters.

Guaranteed constraints:
2 ≤ code.length ≤ 105.

[output] integer

The length of the string that was initially encoded. It is guaranteed that the answer fits in a 32-bit type.
**/
func decodeRLE(code string) int {

    count := 0
    cur := ""
    for _, c := range code {
        
        if c >= 48 && c <= 57 {
            cur += string(c)
        } else {
            if len(cur) > 0 {
                i, _ := strconv.Atoi(string(cur))              
                count += i
                cur = ""
            }
        }
    }
    if len(cur) > 0 {
                i, _ := strconv.Atoi(string(cur))
                count += i
                cur = ""
            }
    return count
}