/*
Challenge: addTwoHugeNumbers
https://codefights.com/interview-practice/task/RvDFbsNC3Xn7pnQfH/description

You're given 2 huge integers represented by linked lists. Each linked list element is a number from 0 to 9999 that represents a number with exactly 4 digits. The represented number might have leading zeros. Your task is to add up these huge integers and return the result in the same format.

Example

For a = [9876, 5432, 1999] and b = [1, 8001], the output should be
addTwoHugeNumbers(a, b) = [9876, 5434, 0].

Explanation: 987654321999 + 18001 = 987654340000.

For a = [123, 4, 5] and b = [100, 100, 100], the output should be
addTwoHugeNumbers(a, b) = [223, 104, 105].

Explanation: 12300040005 + 10001000100 = 22301040105.

Input/Output

[time limit] 4000ms (go)
[input] linkedlist.integer a

The first number, without its leading zeros.

Guaranteed constraints:
0 ≤ a size ≤ 104,
0 ≤ element value ≤ 9999.

[input] linkedlist.integer b

The second number, without its leading zeros.

Guaranteed constraints:
0 ≤ b size ≤ 104,
0 ≤ element value ≤ 9999.

[output] linkedlist.integer

The result of adding a and b together, returned without leading zeros in the same format.
*/
// Definition for singly-linked list:
// type ListNode struct {
//   Value interface{}
//   Next *ListNode
// }
//
func addTwoHugeNumbers(a *ListNode, b *ListNode) *ListNode {

    ar := make([]int, 0)
    br := make([]int, 0)
    for a != nil {
        ar = append(ar, a.Value.(int))
        a = a.Next
    }
    for b != nil {
        br = append(br, b.Value.(int))
        b = b.Next
    }
    reverse(ar)
    reverse(br)

    // add until one list empty
    res := make([]int, 0)
    carry := 0
    idx := 0
    for i := 0; i < len(ar) && i < len(br); i++ {
        x := ar[i] + br[i]
        res = append(res, x)
        idx = i+1
    }
    
    // add the larger list
    for idx < len(ar) || idx < len(br) {
        if idx < len(ar) {
            res = append(res, ar[idx])
        }
        
        if idx < len(br) {
            res = append(res, br[idx])
        }
        idx++
    }
    
    // handle carries
    for n, _ := range res {
        res[n] += carry
        carry = 0
        if res[n] > 9999 {
            res[n] -= 10000
            carry = 1
        }
    }
    if carry > 0 {
        res = append(res, carry)
    }
    
    // reverse and convert to list
    reverse(res)  
    ret := new(ListNode)
    n := ret
    for z, v := range res {
        n.Value = v
        n.Next = new(ListNode)
        if z < len(res)-1 {
            n = n.Next
        }
    }
    n.Next = nil
    
    return ret
}

func reverse(arr []int) {
    for i,j := 0, len(arr)-1; i < len(arr)/2; i,j = i+1,j-1 {
        arr[i], arr[j] = arr[j], arr[i]
    }
}

