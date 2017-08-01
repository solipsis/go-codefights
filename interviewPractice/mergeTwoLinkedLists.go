/**
Challenge: mergeTwoLinkedLists
https://codefights.com/interview-practice/task/6rE3maCQwrZS3Mm2H/description


Note: Your solution should have O(l1.length + l2.length) time complexity, since this is what you will be asked to accomplish in an interview.

Given two singly linked lists sorted in non-decreasing order, your task is to merge them. In other words, return a singly linked list, also sorted in non-decreasing order, that contains the elements from both original lists.

Example

For l1 = [1, 2, 3] and l2 = [4, 5, 6], the output should be
mergeTwoLinkedLists(l1, l2) = [1, 2, 3, 4, 5, 6];
For l1 = [1, 1, 2, 4] and l2 = [0, 3, 5], the output should be
mergeTwoLinkedLists(l1, l2) = [0, 1, 1, 2, 3, 4, 5].
Input/Output

[time limit] 4000ms (go)
[input] linkedlist.integer l1

A singly linked list of integers.

Guaranteed constraints:
0 ≤ list size ≤ 104,
-109 ≤ element value ≤ 109.

[input] linkedlist.integer l2

A singly linked list of integers.

Guaranteed constraints:
0 ≤ list size ≤ 104,
-109 ≤ element value ≤ 109.

[output] linkedlist.integer

A list that contains elements from both l1 and l2, sorted in non-decreasing order.
**/
// Definition for singly-linked list:
// type ListNode struct {
//   Value interface{}
//   Next *ListNode
// }
//
func mergeTwoLinkedLists(l1 *ListNode, l2 *ListNode) *ListNode {
    
    // edge cases
    if l1 == nil && l2 == nil {
        return nil
    } 
    if l1 == nil {
        return l2
    } else if l2 == nil {
        return l1
    }
    
    // initialize
    var start, a, b *ListNode
    if l1.Value.(int) < l2.Value.(int) {
        start = l1
        a = start.Next
        b = l2
    } else {
        start = l2
        b = start.Next
        a = l1
    }
    

    // node that follows path
    cur := start

    // keep finding min until one list empty
    var tmp *ListNode
    for a != nil && b != nil {
        if a.Value.(int) < b.Value.(int) {
            tmp = a
            a = a.Next
            cur.Next = tmp
            cur = cur.Next
        } else {
            tmp = b
            b = b.Next
            cur.Next = tmp
            cur = cur.Next
        }
    }
    // append remaining elements
    if a != nil {
        cur.Next = a
    } else {
        cur.Next = b
    }    
    return start
}

