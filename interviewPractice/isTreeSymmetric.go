/*
Challenge: 
https://codefights.com/interview-practice/task/tXN6wQsTknDT6bNrf/description

Given a binary tree t, determine whether it is symmetric around its center, i.e. each side mirrors the other.

Example

For

t = {
    "value": 1,
    "left": {
        "value": 2,
        "left": {
            "value": 3,
            "left": null,
            "right": null
        },
        "right": {
            "value": 4,
            "left": null,
            "right": null
        }
    },
    "right": {
        "value": 2,
        "left": {
            "value": 4,
            "left": null,
            "right": null
        },
        "right": {
            "value": 3,
            "left": null,
            "right": null
        }
    }
}
the output should be isTreeSymmetric(t) = true.

Here's what the tree in this example looks like:

    1
   / \
  2   2
 / \ / \
3  4 4  3
As you can see, it is symmetric.

For

t = {
    "value": 1,
    "left": {
        "value": 2,
        "left": null,
        "right": {
            "value": 3,
            "left": null,
            "right": null
        }
    },
    "right": {
        "value": 2,
        "left": null,
        "right": {
            "value": 3,
            "left": null,
            "right": null
        }
    }
}
the output should be isTreeSymmetric(t) = false.

Here's what the tree in this example looks like:

    1
   / \
  2   2
   \   \
   3    3
As you can see, it is not symmetric.

Input/Output

[time limit] 4000ms (go)
[input] tree.integer t

A binary tree of integers.

Guaranteed constraints:
0 ≤ tree size < 5 · 104,
-1000 ≤ node value ≤ 1000.

[output] boolean

Return true if t is symmetric and false otherwise.


*/
//
// Definition for binary tree:
// type Tree struct {
//   Value interface{}
//   left *Tree
//   right *Tree
// }
func isTreeSymmetric(t *Tree) bool {
    
    if t == nil {
        return true
    }

    v := []int{}
    inOrder(t, &v)
    if len(v) % 2 != 1 {
        return false
    }
    
    for i := 0; i < len(v)/2; i++ {
        if v[i] != v[len(v)-i-1] {
            return false
        }
    }

    return true
}

func inOrder(t *Tree, vals *[]int) {
    
    // leaf
    if t == nil {
        return
    }
      
    inOrder(t.Left, vals)
    *vals = append(*vals, t.Value.(int))
    inOrder(t.Right, vals)
}
