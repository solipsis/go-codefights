/*
Challenge: digitsProduct
https://codefights.com/arcade/intro/level-12/NJJhENpgheFRQbPRA/description

Given an integer product, find the smallest positive (i.e. greater than 0) integer the product of whose digits is equal to product. If there is no such integer, return -1 instead.

Example

For product = 12, the output should be
digitsProduct(product) = 26;
For product = 19, the output should be
digitsProduct(product) = -1.
Input/Output

[time limit] 4000ms (go)
[input] integer product

Guaranteed constraints:
0 ≤ product ≤ 600.

[output] integer
*/
func digitsProduct(product int) int {

    if (product == 1) {
       return 1
    }
    
    for i := 0; i <= 10000; i++ {
        
        t := i
        prod := 1
        for t > 0 {
            rem := t % 10
            prod *= rem
            t /= 10
        }
        if prod == product {
            return i
        }
    }
    return -1
}
