/**
Challenge: queueForShower
https://codefights.com/tournaments/vbZdzDTcajNBS7xN6/H

You're in your dormitory, standing in the queue for the shower. Each student needs exactly ki minutes in the shower. Unfortunately, sometimes the dormitory's water just turns off and no one can clean up during those times. And of course nobody wants the water turn off while they're in the shower! So if a student realizes that he doesn't have enough time to clean up before the water turns off, he lets the student nearest to him in the queue who can clean up before the water turns off go first.

You are given an array students that represents the queue and the amount of time each student in the queue needs to spend in the shower, and a sorted array timetable that contains the time segments (in minutes) [a, b) during which the water will be turned off. Return the time (in minutes) at which all the students will be clean. The whole process starts at time 0.

Example

For students = [4, 10, 7, 3] and timetable = [[10, 15], [22, 25]], the output should be
queueForShower(students, timetable) = 35.

0:00. Student 1 gets into the shower.
4:00. Student 1 finishes, but student 2 doesn't have enough time because he needs 10 minutes but the water will be turned off in 6 minutes. Student 3 doesn't have enough time to complete his shower either, so they let student 4 go first.
7:00. Student 4 finishes and there are 3 minutes left before the water turns off. Students 2 and 3 still need to wait.
10:00-15:00. The water is turned off.
15:00. Student 2 still doesn't have enough time to shower, because the water will be turned off in 7 minutes. So student 3 takes his shower.
22:00. Student 3 finishes.
22:00-25:00. The water is turned off.
25:00. Finally, student 2 has enough time to shower.
35:00. Student 2 finishes.
Input/Output

[time limit] 4000ms (go)
[input] array.integer students

The queue for the shower. Student i in the queue needs to spend students[i] minutes in the shower.

Guaranteed constraints:
1 ≤ students.length ≤ 1000,
1 ≤ students[i] ≤ 1000.

[input] array.array.integer timetable

The timetable of when the water will be turned off. The array contains time segments [a, b), during which the water is turned off and nobody can shower.

Guaranteed constraints:
0 ≤ timetable.length ≤ 1000,
timetable[i].length = 2,
1 ≤ timetable[i][0] < timetable[i][1] < timetable[i+1][0] < timetable[i+1][1] ≤ 109.

[output] integer

The time in minutes at which all the students will be clean.
**/
func queueForShower(students []int, timetable [][]int) int {

    time := 0
    
    for len(students) > 0 {
        next := nextOut(&timetable, time)
        found := false
        for i := 0; i < len(students); i++ {
            if time + students[i] <= next[0]  {
                
                time += students[i]     
                students = append(students[0:i], students[i+1:]...)
                found = true
                break
            }
        }
        if !found {
            time = next[1]
        }

    }
    
    return time
}

func nextOut(timetable *[][]int, time int) []int {
    
    if len(*timetable) > 0 {
        if time < (*timetable)[0][0] {
            return (*timetable)[0]
        }
        
        v := (*timetable)[0]
        *timetable = (*timetable)[1:]
        return v
    }
    
    return []int{99999999999, 999999999999}
}

