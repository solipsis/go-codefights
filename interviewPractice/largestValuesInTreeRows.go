/*
Challenge: LargestValuesInTreeRows
https://codefights.com/interview-practice/task/m9vC4ALaAeudkwRTF/description

You have a binary tree t. Your task is to find the largest value in each row of this tree. In a tree, a row is a set of nodes that have equal depth. For example, a row with depth 0 is a tree root, a row with depth 1 is composed of the root's children, etc.

Return an array in which the first element is the largest value in the row with depth 0, the second element is the largest value in the row with depth 1, the third element is the largest element in the row with depth 2, etc.

Example

For

t = {
    "value": -1,
    "left": {
        "value": 5,
        "left": null,
        "right": null
    },
    "right": {
        "value": 7,
        "left": null,
        "right": {
            "value": 1,
            "left": null,
            "right": null
        }
    }
}
the output should be largestValuesInTreeRows(t) = [-1, 7, 1].

The tree in the example looks like this:

    -1
   / \
  5   7
       \
        1
In the row with depth 0, there is only one vertex - the root with value -1;
In the row with depth 1, there are two vertices with values 5 and 7, so the largest value here is 7;
In the row with depth 2, there is only one vertex with value 1.
Input/Output

[time limit] 4000ms (go)
[input] tree.integer t

A binary tree of integers.

Guaranteed constraints:
0 ≤ tree size ≤ 5 · 104,
-1000 ≤ node value ≤ 1000.

[output] array.integer

An array of the largest values in each row of t.
*/
//
// Definition for binary tree:
// type Tree struct {
//   Value interface{}
//   Left *Tree
//   Right *Tree
// }
func largestValuesInTreeRows(t *Tree) []int {
    
    r := make([]int, 0)
    q := make([]*Tree, 0) // queue
    
    if t == nil {
        return r
    }
    
    q = append(q, t)
    depth := -1
    depthNode := t // increase depth when we see this node
    
    for len(q) > 0 {
        
        // pop
        cur := q[0]
        q = q[1:] 
        
        // increase depth
        if cur == depthNode {
            depth++
            r = append(r, -2000)
            depthNode = nil
        }
        if cur.Value.(int) > r[depth] {
            r[depth] = cur.Value.(int)
        }
        
        // add children to queue
        if cur.Left != nil {
            q = append(q, cur.Left)
            if depthNode == nil {
                depthNode = cur.Left
            }
        }
        if cur.Right != nil {
            q = append(q, cur.Right)
            if depthNode == nil {
                depthNode = cur.Right
            }
        }
    }  
    return r
}
