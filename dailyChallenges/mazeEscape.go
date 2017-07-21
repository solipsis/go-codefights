/**
Challenge: mazeEscape
https://codefights.com/challenge/dF3Qvny3Yn23xvLRb/solutions


The Royal Priest of Ur has been kidnapped by the Mad Dragon of Akkad. The gods have chosen you to rescue him! The Dragon has hidden the Priest at the end of a maze, and you must successfully traverse the maze in order to find him. Your task is to determine the minimum number of moves needed to escape the maze, or -1 if it isn't possible to escape.

The maze is a grid of width w and height h.

A kindly demigod has given you a list impassable, provided as an array of integers that represent the coordinates of impassable cells. Each integer pair at indices 2 · i and 2 · i + 1 (where  0 ≤ i < impassable.length/2 - 1) represents an impassable cell at coordinate (x, y). For example, the list [1, 2, 3, 4] represents two impassable cells at (1, 2) and (3, 4) (but NOT at (2, 3)).

You start at cell (0, 0) and escape by reaching the opposite corner at (w - 1, h - 1). It is guaranteed that the start and end cells are not impassable.

You can only move up, down, left, and right. You can't move diagonally.

Example
For w = 4, h = 4 and impassable = [1, 0, 1, 1, 2, 2, 3, 1], the output should be
mazeEscape(w, h, impassable) = 6.


[time limit] 4000ms (go)
[input] integer w

The width of the maze.

Guaranteed constraints:
2 ≤ w ≤ 150.

[input] integer h

The height of the maze.

Guaranteed constraints:
2 ≤ h ≤ 150.

[input] array.integer impassable

Each integer pair at indices 2 · i and 2 · i + 1 represents an impassable cell at coordinate (x, y).

Guaranteed constraints:
0 ≤ impassable.length ≤ w · h - 2 .

[output] integer

Return the minimum number of moves needed to escape from the maze and find the Priest, or -1 if it is not possible.
**/
func mazeEscape(w int, h int, imp []int) int {
    seen := make(map[cell]bool)
    for i := 0; i < len(imp); i+=2 {
        seen[cell{imp[i], imp[i+1]}] = true
    }
    
    queue := make([]cell, 0)
    queue = append(queue, cell{0,0})
    depthTrigger := queue[0]
    depth := 0
    
    for len(queue) > 0 {
        c := queue[0]
        seen[c] = true
        queue = queue[1:] // pop
        
        x := c.x
        y := c.y
        
        if y > 0 {
            enqueue(x,y-1, &queue, seen)     
        }
        if c.y < h-1 {
            enqueue(x, y+1, &queue, seen)
        }
        if c.x > 0 {
            enqueue(x-1,y, &queue, seen)
        }
        if c.x < w-1 {
            enqueue(x+1,y, &queue, seen)
        }

        if c.x == w-1 && c.y == h-1 {
            return depth
        }
        if c == depthTrigger && len(queue) > 0 {
            depth++
            depthTrigger = queue[len(queue)-1]
        }
    }
    return -1
}

func enqueue(x,y int, q *[]cell, s map[cell]bool) {
    c := cell{x,y}
    if !s[c] {
        s[c] = true
        *q = append(*q, c) 
    }
}

type cell struct {
    x, y int
}