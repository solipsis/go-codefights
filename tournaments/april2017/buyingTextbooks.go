/**
Challenge: buyingTextbooks
https://codefights.com/tournaments/BLhuiuSY4neuXPXet/D

The university library needs to buy some new textbooks. There are n textbooks numbered from 1 to n. The library gives you a number of textbooks n and an array textbooks that contains pairs [k, kAmount]. The library needs to purchase at least kAmount copies of each textbook with numbers k, 2 * k, 3 * k, etc.

Help the library calculate how many copies of each book it should buy!

Example

For n = 6 and textbooks = [[3, 1], [2, 2]], the output should be
buyingTextbooks(n, textbooks) = [0, 2, 1, 2, 0, 2].

The library needs the following textbooks:

Each 3rd textbook:
book 3 × 1
book 6 × 1
Each 2nd textbook:
book 2 × 2
book 4 × 2
book 6 × 2
Note that the library has two requirements about book 6 it needs to meet. To meet the first requirement, they need 1 copy of the book, but to meet the second, they need 2 copies. To meet the requirements, they should buy 2 copies of textbook 6.

Altogether, the library should buy two copies of book 2, one copy of book 3, two copies of book 4, and two copies of book 6. The library doesn't need books 1 and 5.

For n = 7 and textbooks = [[1, 2]], the output should be
buyingTextbooks(n, textbooks) = [2, 2, 2, 2, 2, 2, 2].

Every 1st textbook means all the textbooks, so the library needs 2 copies of each textbook.

Input/Output

[time limit] 4000ms (go)
[input] integer n

The number of textbooks.

Guaranteed constraints:
1 ≤ n ≤ 100 .

[input] array.array.integer textbooks

Pairs [k, kAmount], such that the library wants at least kAmount copies of each kth textbook.

Guaranteed constraints:
1 ≤ textbooks.length ≤ n ,
textbooks[i].length = 2,
1 ≤ textbooks[i][0] ≤ n , unique for all pairs,
1 ≤ textbooks[i][1] ≤ 109 .

[output] array.integer

For every book i, how many copies of it the library should purchase.
**/
func buyingTextbooks(n int, textbooks [][]int) []int {

    result := make([]int, n)
    
    for _, v := range textbooks {
        for index, val := range result {
            if (index+1) % v[0] == 0 && v[1] > val {
                result[index] = v[1]
            } 
        } 
    }
    return result
}
