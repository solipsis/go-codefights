/**
Challenge: steganographer
https://codefights.com/challenge/PyzYWH2rf7eJDh9QH


A steganographer is a special piece of software that lets Alice send messages to Bob through a simple bitmap picture, without Eve knowing that the picture actually contains a message.

The process is really quite simple: Since the human eye can barely distinguish the difference between a 255 and 254 color code in the RGBA construct, one can use the least significant bit of each color to hide a message. To make it more complex, it's possible to use only one of the colors, and it is even possible to encrypt the message prior to hiding it (which is overkill, so we won't do it here).

Your task is to generate a program that will take an image encoded as an rgba_arr array of pixels, and extract a message from it using a 4-bit code representing the colors used in encryption. Each 32-bit integer in rgba_arr, represents 4 groups of 8 bits which are respectively Red, Green, Blue, and Alpha colors.

Example

For

rgb_arr = [
 [1852001592, 642517586,  1375810157, 978611461,  209602908,  1815751467, 1918763865],
 [2099972387, 625962066,  209155390,  662511729,  1074152308, 1464016421, 240729689],
 [1544184132, 1882073136, 371947289,  839922279,  459299100,  1416848664, 1818624513],
 [187696431,  50614298,   1260153149, 1948346715, 1028787982, 1426982168, 1228945414],
 [795629403,  1450318166, 2063688524, 510462577,  1214057803, 1848402996, 139541769]
]
and code = 3, the output should be
steganographer(rgb_arr, code) = "CodeFights".

Here's the breakdown:
If we take the first line of rgba_arr and convert it to binary, we obtain the following:

01101110011000110100110100111000
00100110010011000000101001010010
01010010000000010011001001101101
00111010010101000110110100000101
00001100011111100100100101011100
01101100001110100010101100101011
01110010010111100000001101011001
You can further define the codes by splitting each 32-bit string into four 8-bit chunks. Since the code is 310 = 00112, it means the third and the fourth chunks (which stand for the blue and alpha segments) should be used. In other words, you should extract the least significant bits from them.


01101110 01100011 01001101 00111000
00100110 01001100 00001010 01010010
01010010 00000001 00110010 01101101
00111010 01010100 01101101 00000101
00001100 01111110 01001001 01011100
01101100 00111010 00101011 00101011
01110010 01011110 00000011 01011001
The ASCII representation of 1000011 1101111 is "Co". Continue the same process throughout the matrix to extract all the characters.

Input/Output

[time limit] 4000ms (go)
[input] array.array.integer64 rgba_arr

A matrix of integers, each of which represents a 4 × 8 bits sequence of the RGBA encoding.

Guaranteed constraints:
1 ≤ rgba_arr.length ≤ 1000,
rgba_arr[i].length = 7,
1 ≤ rgba_arr[i][j] ≤ 231.

[input] integer code

A 4-bit code indicating which colors to use in the steganographer.

Guaranteed constraints:
1 ≤ code ≤ 15.

[output] string

The message that's been retrieved from the picture.
**/
func steganographer(rgba_arr [][]int64, code int) string {
  
    sv := ""
    r := ""
    for _, arr := range rgba_arr {
        for _, v := range arr {
            s := fmt.Sprintf("%032b", v)
          
            for i := 0; i < 4; i++ {
                if (1 << uint(3-i)) & code > 0 {
                    t := s[i*8:i*8+8]
                    sv += string(t[7])               
                }
                if len(sv) == 7 {
                    conv, _ := strconv.ParseInt(sv, 2, 64)
                    r += string(conv)
                    sv = ""
                }                            
            }
        }
        
    } 
    return r
}
