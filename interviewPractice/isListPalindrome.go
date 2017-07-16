/*
Challenge: isListPalindrome
https://codefights.com/interview-practice/task/HmNvEkfFShPhREMn4/description

Note: Try to solve this task in O(n) time using O(1) additional space, where n is the number of elements in l, since this is what you'll be asked to do during an interview.

Given a singly linked list of integers, determine whether or not it's a palindrome.

Example

For l = [0, 1, 0], the output should be
isListPalindrome(l) = true;
For l = [1, 2, 2, 3], the output should be
isListPalindrome(l) = false.
Input/Output

[time limit] 4000ms (go)
[input] linkedlist.integer l

A singly linked list of integers.

Guaranteed constraints:
0 ≤ list size ≤ 5 · 105,
-109 ≤ element value ≤ 109.

[output] boolean

Return true if l is a palindrome, otherwise return false.
*/
// Definition for singly-linked list:
// type ListNode struct {
//   Value interface{}
//   Next *ListNode
// }
//
func isListPalindrome(l *ListNode) bool {
    
    le := length(l)
    if le < 2 {
        return true
    }
    if le == 2 {
        return l.Value == l.Next.Value
    }
    
    mid := (le-1) / 2
    m := l
    for x := 0; x < mid; x++ {
        m = m.Next
    }
  
    // reverse list after midpoint
    reverse(m)
    // check if both halves the same
    m = m.Next
    for m != nil && l != nil {
        if m.Value != l.Value {
            return false
        }
        m = m.Next
        l = l.Next
    }
    return true  
}

func length(l *ListNode) int {
    
    x := 0
    for l != nil {
        x++
        l = l.Next
    }
    return x
}

// reverse everything after node n
func reverse(n *ListNode) {
    
    l := n
    m := l.Next
    r := m.Next  
    for r != nil {
        m.Next = l
        l = m
        m = r 
        r = r.Next
    }
    m.Next = l
    n.Next.Next = nil
    n.Next = m 
}

