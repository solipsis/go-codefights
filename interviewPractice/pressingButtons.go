/**
Challenge: pressingButtons
https://codefights.com/interview-practice/task/SyBbdPYC86RdkKvjH/description


Given a string of digits, return all of the possible non-empty letter combinations that the number could represent. The mapping of digits to letters is the same as you would find on a telephone's buttons, as in the example below:

The resulting array should be sorted lexicographically.

Example

For buttons = "42", the output should be
pressingButtons(buttons) = ["ga", "gb", "gc", "ha", "hb", "hc", "ia", "ib", "ic"].

Input/Output

[time limit] 4000ms (go)
[input] string buttons

A string composed of digits ranging from '2' to '9'.

Guaranteed constraints:

0 ≤ buttons.length ≤ 7.

[output] array.string
**/
func pressingButtons(buttons string) []string {
    
    arr := []string{2:"abc",3:"def",4:"ghi",5:"jkl",6:"mno",7:"pqrs",8:"tuv",9:"wxyz"}
    res := []string{}
    perms(buttons, arr, "", &res)
  
    fmt.Println(len(res))
    if len(res) == 0 {
        return []string{}
    }
    return res
    
}

func perms(buttons string, arr []string, cur string, result *[]string) {
    
    if len(cur) == len(buttons) {
        if len(buttons) > 0 {
            *result = append(*result, cur)
        }
        return
    }
    
    num, _ := strconv.Atoi(string(buttons[len(cur)]))
   
    letters := arr[num]
    for _, c := range letters {
        perms(buttons, arr, cur + string(c), result)
    }
}
