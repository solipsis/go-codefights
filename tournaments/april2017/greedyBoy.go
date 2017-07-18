/**
Challenge: greedyBoy
https://codefights.com/tournaments/BLhuiuSY4neuXPXet/A

John wants to invite some friends to his birthday party. He knows that each friend needs to eat exactly k portions of food in order to be happy. They will still be hungry if they eat fewer than k portions, and they will get too full if they eat more than k. John is going to buy some snacks (as many as necessary) and possibly one of the cakes from the local supermarket. Each snack has 10 portions, and the ith cake from the store has cakes[i] portions. It's guaranteed that all of the cakes have fewer than 10 portions.

What is the smallest amount of friends that John can invite to his party, such that everyone is happy and no food is left over? Remember that John needs k portions for himself as well, and he wants at least one friend at his party.

Given the value k and an array cakes that contains the amount of portions for each cake option, return the smallest number of friends John can invite.

Example

For k = 7 and cakes = [6, 8, 9], the output should be
greedyBoy(k, cakes) = 3.

The smallest number of people John can invite is 3. Including John, there are 4 people at the party who eat 28 portions in total. The 28 portions are divided among two snacks (10 portions in each) and one cake (8 portions).

For k = 13 and cakes = [3], the output should be
greedyBoy(k, cakes) = 9.

The smallest number of people John can invite is 9. Including John, there are 10 people at the party who eat 130 portions (13 snacks and 0 cake). Note that since John wants to invite at least one friend, the output can't be 0.

Input/Output

[time limit] 4000ms (go)
[input] integer k

The number of food portions each guest needs in order to be happy.

Guaranteed constraints:
1 ≤ k ≤ 100.

[input] array.integer cakes

An array that contains the number of portions for each cake option.

Guaranteed constraints:
1 ≤ cakes.length ≤ 9,
1 ≤ cakes[i] ≤ 9 .

[output] integer

The smallest number of people that John can invite to his party.
**/
func greedyBoy(k int, cakes []int) int {

    cakes = append(cakes, 0)
    tot := 2
    for  {
        for _, c := range cakes {
            if (tot * k - c) % 10 == 0  {
                return tot-1
            }
        }
        tot++
    }
    return -1
    
}
