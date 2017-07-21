/**
Challenge: bitsToFloat
https://codefights.com/interview-practice/task/S6JDgmp425yD3NZcQ/description

You have a float number stored in a integer in a following format:

The fractional part is stored in bits 0-15 (counting from the least significant bit)
The integer part is stored in bits 16-30
Return the input converted to a regular float (integer part.fractional part).
Example

For bits = 65539, the output should be
bitsToFloat(bits) = 1.3.

6553910 = 00000001 00000000 00000000 000000112, so the integer part is 1 and the fractional part is 112 = 310.

Input/Output

[time limit] 4000ms (go)
[input] integer bits

Guaranteed constraints:
0 ≤ bits ≤ 231 - 1.

[output] float

Return bits converted to a float.
**/
func bitsToFloat(bits int) float64 {
  intp := float64(bits & 0xFFFF)
  frap := float64(bits & 0xFFFF)
  for frap >= 1. {
    frap /= 10.
  }
  return intp + frap
}
