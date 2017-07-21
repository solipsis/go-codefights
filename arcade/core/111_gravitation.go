/**
Challenge: Gravitation
https://codefights.com/arcade/code-arcade/waterfall-of-integration/hqrYesGKEaKQnv7Sv/description

You are given a vertical box divided into equal columns. Someone dropped several stones from its top through the columns. Stones are falling straight down at a constant speed (equal for all stones) while possible (i.e. while they haven't reached the ground or they are not blocked by another motionless stone). Given the state of the box at some moment in time, find out which columns become motionless first.

Example

For

rows = ["#..##",
        ".##.#",
        ".#.##",
        "....."]
the output should be gravitation(rows) = [1, 4].

Check out the image below for better understanding:



Input/Output

[time limit] 4000ms (go)
[input] array.string rows

A non-empty array of strings of equal length consisting only of #-s and .-s describing the box at a specific moment in time. Sharps represent stones, and dots represent empty cells. row[0] corresponds to the upper row. Last element of rows corresponds to the ground level.

Guaranteed constraints:
2 ≤ rows.length ≤ 10,
2 ≤ rows[i].length ≤ 10.

[output] array.integer

A sorted array containing numbers of all columns (leftmost column's number is 0) in which movements will stop at the same second and earlier than in any other column. Assume that if there are no stones in a column then movement stops immediately, i.e. after 0 seconds.
**/
func gravitation(rows []string) []int {
    
    arr := make([][]rune,len(rows))  
    for n, row := range rows {
        for _, r := range row {
            arr[n] = append(arr[n], r)
        }
    }

    time := make([]int, len(rows[0]))
    res := []int{}
    minTime := 99
    
    // start + largest gap == time till done
    for col := 0; col < len(arr[0]); col++ {
        
        start := 0
        gap := 0
        maxGap := 0
        first := false
        found := false
        
        for row := len(arr)-1; row >= 0; row-- {          
            if !first && arr[row][col] == '.' {
                start++
            }
            if found && arr[row][col] == '.' {
                gap++
            }
            if arr[row][col] == '#' {
                first = true
                found = true
                if gap > maxGap {
                    maxGap = gap
                }
                gap = 0
            }
         
        }
        if !first {
            time[col] = 0
        } else {
            time[col] = start + maxGap
        }
        if time[col] < minTime {
            minTime = time[col]
        }
    }
    
    for n, v := range time {
        if v == minTime {
            res = append(res, n)
        }
    }
    return res
}
