/*
Challenge: removeKFromList
https://codefights.com/interview-practice/task/gX7NXPBrYThXZuanm/description

Note: Try to solve this task in O(n) time using O(1) additional space, where n is the number of elements in the list, since this is what you'll be asked to do during an interview.

Given a singly linked list of integers l and an integer k, remove all elements from list l that have a value equal to k.

Example

For l = [3, 1, 2, 3, 4, 5] and k = 3, the output should be
removeKFromList(l, k) = [1, 2, 4, 5];
For l = [1, 2, 3, 4, 5, 6, 7] and k = 10, the output should be
removeKFromList(l, k) = [1, 2, 3, 4, 5, 6, 7].
Input/Output

[time limit] 4000ms (go)
[input] linkedlist.integer l

A singly linked list of integers.

Guaranteed constraints:
0 ≤ list size ≤ 105,
-1000 ≤ element value ≤ 1000.

[input] integer k

An integer.

Guaranteed constraints:
-1000 ≤ k ≤ 1000.

[output] linkedlist.integer

Return l with all the values equal to k removed.
*/
// Definition for singly-linked list:
// type ListNode struct {
//   Value interface{}
//   Next *ListNode
// }
//
func removeKFromList(l *ListNode, k int) *ListNode {

    if (l == nil) {
        return l
    }
    
    // find first node
    i := l.Value.(int)
    for i == k {
        l = l.Next
        if l == nil {
            return l
        }
        i = l.Value.(int)
    }
    
    cur := l.Next
    ptr := l
    // remove items
    for cur != nil {
        i := cur.Value.(int)
        if i != k {
            ptr.Next = cur
            ptr = cur
        }
        cur = cur.Next
    }
    ptr.Next = nil
    
    
    return l;
}
