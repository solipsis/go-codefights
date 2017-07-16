/*
Challenge: validTime
https://codefights.com/arcade/intro/level-12/ywMyCTspqGXPWRZx5/description

Check if the given string is a correct time representation of the 24-hour clock.

Example

For time = "13:58", the output should be
validTime(time) = true;
For time = "25:51", the output should be
validTime(time) = false;
For time = "02:76", the output should be
validTime(time) = false.
Input/Output

[time limit] 4000ms (go)
[input] string time

A string representing time in HH:MM format. It is guaranteed that the first two characters, as well as the last two characters, are digits.

[output] boolean

true if the given representation is correct, false otherwise.
*/
func validTime(time string) bool {

    if len(time) != 5 {
        return false
    } 
    h, _ := strconv.Atoi(string(time[:2]))
    m, _ := strconv.Atoi(string(time[3:])) 
    if (h > 23 || m > 59) {
        return false
    }  
    return true
}
