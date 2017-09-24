/**
Challenge: stackIt
https://codefights.com/challenge/qTHXJTeYezBvkHSoR

The local supermarket is planning a HUGE sale of XYZ soda and the manager, Jim, is thinking they should build a perfect triangular structure using all their XYZ stock. He does not want any soda packs being left over so every single one must be used to construct the structure. However, he is not sure if this is possible given their current stock so he explains his plan to the newly hired stock boy, Tommy, and asks him to have an answer by tomorrow.
That night Tommy goes home and pulls out his building blocks to simulate the structure:

The next morning Jim asks Tommy:
Can you do it?

Example
For packs = 6, the output should be
stackIt(packs) = true.

Input/Output

[time limit] 4000ms (go)
[input] integer64 packs

Guaranteed constraints:
0 ≤ packs < 1012.

[output] boolean
**/
func stackIt(p int64) bool {
    var x, z int64
    
    // triangular numbers xn = n(n+1)/2
    for z = 1; x < p; x,z = z*(z+1)/2, z+1 {
    }
    return x == p
}
