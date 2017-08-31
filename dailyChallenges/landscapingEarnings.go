/**
Challenge: landscapingEarnings
https://codefights.com/challenge/ckYJtErc8vh8bA5Yu/solutions/JxLRaAQZZo5qLTehT

You are making a system that will help your client, a landscaping company, assign their employees to projects in such a way that their earnings are maximized. The company offers a few different services, each of which take an employee a certain amount of time to complete.

:tree_removal => {time: 120, money: 300},
:mowing => {time: 30, money: 14},
:raking => {time: 10, money: 7},
:trimming => {time: 5, money: 4},
:digging => {time: 4, money: 3}
The company policy is that an employee must be scheduled for EXACTLY the amount of time they say they can work that month. The company wants to maximize the profits that they can make from a given employee in a month.

Employees can work anywhere from 1 to 10 jobs a month, as long as the total amount of time is equal to the number of hours they put in for that month. An employee putting in more than 10 jobs a month is not allowed. An employee putting in fewer than 4 hours of work a month is also not allowed.

Example
For time = 20, the output should be
landscapingEarnings(time) = 16.
Here are the possible ways to assign jobs to this person:

raking {x2} => time = 20, money = 14;
raking, trimming, trimming => time = 20, money = 15;
trimming {x4} => time = 20, money = 16;
digging {x5} => time = 20, money = 15.
So the most profitable assignment of jobs would be to have the employee do 4 trimmings, making 16 money for the company.

Input/Output

[time limit] 4000ms (go)
[input] integer time

The amount of time an employee can work in a month.

Guaranteed constraints:
4 ≤ time ≤ 1200.

[output] integer

The max amount of profit the company can make from a given employee in a month.
**/
func landscapingEarnings(time int) int {   
    memo := make([][]int, 10)
    for n := range memo {
        memo[n] = make([]int, 1201)
 
   }
    // first row of dp table  
    memo[0][4] = 3
    memo[0][5] = 4
    memo[0][10] = 7
    memo[0][30] = 14
    memo[0][120] = 300
    
    // bottom up dp approach
    tm := map[int]int {120:300, 30:14, 10:7, 5:4, 4:3}
    for d := 1; d < 10; d++ {
        for t := 0; t < time+1; t++ {
            max := memo[d-1][t]
            for tk, vm := range tm {
                if t - tk < 1 {
                    continue
                }
                v := memo[d-1][t-tk]
                if v > 0 && v + vm > max {
                    max = v + vm
                }
            }
            memo[d][t] = max  
        }
    }
    return memo[9][time]
}
