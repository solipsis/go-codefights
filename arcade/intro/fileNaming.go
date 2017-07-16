/*
Challenge: fileNaming
https://codefights.com/arcade/intro/level-12/sqZ9qDTFHXBNrQeLC/description

You are given an array of desired filenames in the order of their creation. Since two files cannot have equal names, the one which comes later will have an addition to its name in a form of (k), where k is the smallest positive integer such that the obtained name is not used yet.

Return an array of names that will be given to the files.

Example

For names = ["doc", "doc", "image", "doc(1)", "doc"], the output should be
fileNaming(names) = ["doc", "doc(1)", "image", "doc(1)(1)", "doc(2)"].

Input/Output

[time limit] 4000ms (go)
[input] array.string names

Guaranteed constraints:
5 ≤ names.length ≤ 15,
1 ≤ names[i].length ≤ 15.

[output] array.string
*/
func fileNaming(names []string) []string {
    
    m := make(map[string]int)
    r := []string{}
    for _, s := range names {
        m[s] += 1
        if (m[s] > 1) {
            for z := 1; ; z++ {
                dup := s + "(" + strconv.Itoa(z) + ")"
                if m[dup] == 0 {
                    m[dup] += 1
                    r = append(r,dup)
                    break
                }
            }  
        } else {
            r = append(r,s)
        }
    }
    
    return r
}
