/*
Challenge: almontIncreasingSequence
https://codefights.com/arcade/intro/level-2/2mxbGwLzvkTCKAJMG

Given a sequence of integers as an array, determine whether it is possible to obtain a strictly increasing sequence by removing no more than one element from the array.

Example

For sequence = [1, 3, 2, 1], the output should be
almostIncreasingSequence(sequence) = false;

There is no one element in this array that can be removed in order to get a strictly increasing sequence.

For sequence = [1, 3, 2], the output should be
almostIncreasingSequence(sequence) = true.

You can remove 3 from the array to get the strictly increasing sequence [1, 2]. Alternately, you can remove 2 to get the strictly increasing sequence [1, 3].

Input/Output

[time limit] 4000ms (go)
[input] array.integer sequence

Guaranteed constraints:
2 ≤ sequence.length ≤ 105,
-105 ≤ sequence[i] ≤ 105.

[output] boolean

Return true if it is possible to remove one element from the array in order to get a strictly increasing sequence, otherwise return false.
*/

func almostIncreasingSequence(sequence []int) bool {

    for i := 0; i < len(sequence)-1; i++ {
        if sequence[i+1] <= sequence[i] {
            a := sequence[i+1:]
            b := []int{sequence[i]}
            if i > 0 {
                a = append([]int{sequence[i-1]}, a...)
                b = append([]int{sequence[i-1]}, b...)
            }     
            b = append(b, sequence[i+2:]...)        
            return isIncreasing(a) || isIncreasing(b)
        }
    }
    return true
}



 func isIncreasing(a[]int) bool {
     for i := 0; i < len(a)-1; i++ {
         if a[i+1] <= a[i] {
             return false
         }
     }
     return true
 }
