/**
Challenge: surviveIt
https://codefights.com/challenge/Bw63XoZg24SpfJtcR

The system will keep destroying alien ships until there's only one left. What number will this spacecraft have?

Example

For spacecrafts = 10, the output should be
surviveIt(spacecrafts) = 8.

Initially, there are 10 spacecrafts. After your first attack, only half of them will survive: 2, 4, 6, 8 and 10.
After your second attack, only 2 spacecrafts will remain: 4 and 8.
Finally, spacecraft 4 will be vaporized, so spacecraft 8 will be the only one left.

Output/Input

[time limit] 4000ms (go)
[input] integer spacecrafts

Guaranteed constraints:
1 ≤ spacecrafts ≤ 109.

[output] integer

The number of the last spacecraft to survive your attacks.
**/
import "math"

func surviveIt(s int) int {
    return 1 << uint(math.Log2(float64(s)))
}
