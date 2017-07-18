/**
Challenge: possibleCapitals
https://codefights.com/tournaments/vbZdzDTcajNBS7xN6/D

The inhabitants of the country Murmuria want to choose a new capital city. To calculate which cities can become the capital, they draw a map of the country. The map is an n × m matrix in which some cells are marked as cities. A city can become the capital if, in this map, it has at least one city in the same row and at least one city in the same column except itself.

Given a (0, 1) matrix country where country[i][j] = 1 indicates a city, return the number of all the possible capitals in Murmuria.

Example

For

country = [[0, 0, 1, 0, 1],
           [1, 0, 0, 0, 0],
           [1, 1, 1, 1, 1],
           [0, 0, 0, 1, 0],
           [1, 0, 1, 1, 0]]
the output should be
possibleCapitals(country) = 9.

There are 12 cities on this map, and only 3 of them don't meet the requirements needed to be the capital: the city in the second row, the city in the fourth row, and city in the second column.

Input/Output

[time limit] 4000ms (go)
[input] array.array.integer country

The map of the country. country[i][j] = 1 if there is a city at that location and 0 otherwise.

Guaranteed constraints:
1 ≤ country.length ≤ 100,
1 ≤ country[i].length ≤ 100,
country[i].length = country[j].length,
0 ≤ country[i][j] ≤ 1.

[output] integer

The number of possible capital cities in this country.
**/
func possibleCapitals(country [][]int) int {

    count := 0
    for i := 0; i < len(country); i++ {
        for j := 0; j < len(country[0]); j++ {
            if country[i][j] == 1 {
                colcount := 0
                for col := 0; col < len(country[0]); col++ {
                    if country[i][col] == 1 && col != j {
                        colcount++
                    }
                }
                rowcount := 0
                for row := 0; row < len(country); row++ {
                    if country[row][j] == 1 && row != i {
                        rowcount++
                    }
                }
                if rowcount > 0 && colcount > 0 {
                    count++
                }
            }
        }
    }
    return count
}
