/**
Challenge: splitPresentCost
https://codefights.com/tournaments/ZS9S4h3HFKsuWRMnD/E

You're going to buy a present for your friend's birthday party. The present costs s dollars. Your friend John is going to the party too and has no idea what to buy, so he wants to go in on your present and split the cost with you. Both you and John are going to spend no less than l dollars for the gift, but also not more than r dollars. Now you want to know the number of ways it's possible to split s dollars between the both of you into two amounts of a and b dollars such that l ≤ a ≤ r and l ≤ b ≤ r.

Note: For example, a = 2, b = 3 and a = 3, b = 2 are two different ways to split 5 dollars.

Example

For l = 2, r = 4 and s = 5, the output should be
splitPresentCost(l, r, s) = 2.

There are only two ways to split 5 dollars between you and John:

a = 2 and b = 3;
a = 3 and b = 2.
For l = 1, r = 3 and s = 7, the output should be
splitPresentCost(l, r, s) = 0.

Input/Output

[time limit] 4000ms (go)
[input] integer64 l

Guaranteed constraints:
1 ≤ l ≤ 1015.

[input] integer64 r

Guaranteed constraints:
l ≤ r ≤ 1015.

[input] integer64 s

Guaranteed constraints:
1 ≤ s ≤ 1015.

[output] integer64

The number of ways to split the cost of the present s into two values a and b such that l ≤ a ≤ r and l ≤ b ≤ r.
**/
func splitPresentCost(l int64, r int64, s int64) int64 {

    count := int64(0)
    if r > s {
        r = s
    }
    if l > s {
        return 0
    }
    if l + r < s {
        l += s - (l+r)
        if l > r {
            return 0
        }
    }
    if l + r > s {
        r -= (l+r) - s
        if l > r {
            return 0
        }
    }
    
    if (r-l > 1) {
        count = (((r - l) / 2) * 2) + 1
    } else {
        if l == r {
            count = 1
            
        } else if r - l == 1 {
            return 2

        } else {
            count = 2
        }
    }

    if (r % 2 != l%2) {count++}
    return count
}

