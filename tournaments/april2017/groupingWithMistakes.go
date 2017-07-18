/**
Challenge: groupingWithMistakes
https://codefights.com/tournaments/BLhuiuSY4neuXPXet/B


A teacher told his students to line up and split into n groups by counting in the following way: 1, 2, 3, ..., n - 1, n, n - 1, ..., 3, 2, 1, 2, 3, ..., and so on. But some of the students were being noisy and weren't listening as their neighbors said their number, so it's possible that they may have counted incorrectly.

Given the array students, which represents how the students counted themselves, return the number of students who were placed in the wrong groups (compared to the correct counting order).

Example

For n = 4 and students = [1, 2, 3, 4, 3, 2, 1, 2, 3], the output should be
groupingWithMistakes(n, students) = 0.

None of the students made a mistake, so the order is absolutely correct.

For n = 4 and students = [1, 2, 2, 3, 4, 3, 2, 1, 2], the output should be
groupingWithMistakes(n, students) = 7.

The third student made a mistake, after which the other students started counting incorrectly. So all the students after the second one were put into the wrong groups.

Input/Output

[time limit] 4000ms (go)
[input] integer n

The number of groups into which the students should be split.

Guaranteed constraints:
2 ≤ n ≤ 100.

[input] array.integer students

How the students counted themselves.

Guaranteed constraints:
n ≤ students.length ≤ 800,
1 ≤ students[i] ≤ n.

[output] integer

The number of students who were put into the wrong groups as a result of counting errors.


**/
func groupingWithMistakes(n int, students []int) int {

    arr := generate(len(students),n)
    wrong := 0
    
    for n, v := range students {
        if v != arr[n] {
            wrong++
        }
    }
    return wrong
}

func generate(n, p int) []int {
    arr := make([]int, n)    
    var f func(int) int = add
    val := 1
    for i := 0; i < n; i++ {
        arr[i] = val
        val = f(val)
        if val == p {
            f = sub
        }
        if val == 1 {
            f = add
        }       
    }
    return arr
}

func add(i int) int {
    return i+1
}
func sub(i int) int {
    return i-1
}
