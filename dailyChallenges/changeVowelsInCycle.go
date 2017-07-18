/**
Challenge: changeOfVowelsInCycle
https://codefights.com/challenge/CA2ShAfLkb5D5JHCG

The vowels are 'a', 'e', 'i', 'o', and 'u' and can be upper- and lowercase.

Example

For cycle = 1 and text = "potato", the output should be
changeOfVowelsInCycle(cycle, text) = "ototap".

Reversed, the text becomes "otatop". The vowels in the text are 'o', 'a' and 'o' (in this exact order). When shifted by one, they become 'o', 'o' and 'a', so the final answer is "ototap".

Input/Output

[time limit] 4000ms (go)
[input] integer cycle

A number that defines by how many characters the vowels should be shifted.

Guaranteed constraints:
1 ≤ cycle ≤ 350.

[input] string text

The text, consisting of upper- and lowercase English letters and whitespace characters.

Guaranteed constraints:
1 ≤ word.length ≤ 100.

[output] string

The text, modified as described above.
**/
func changeOfVowelsInCycle(c int, t string) string {

    //z([]byte(text))
    k := []rune(t)
    z(k)
    v := []rune{}
    for _, r := range k {
        switch r {
        case 'a','e','i','o','u','A','E','I','O','U': 
            v = append(v, r)
        }
    }
    
    if (len(v) > 0) {
        c %= len(v)
        z(v[:len(v)-c])
        z(v[len(v)-c:])
        z(v)
    }
    x := 0
    for n, r := range k {
        switch r {   
        case 'a','e','i','o','u','A','E','I','O','U': 
            k[n] = v[x]
            x++
        }
    }  
    return string(k)
}


func z(s []rune) {
    for i,j := 0,len(s)-1; i < len(s)/2; i,j = i+1,j-1 {
        s[i], s[j] = s[j], s[i]
    }
}
