/*
Challenge: groupingDishes
https://codefights.com/interview-practice/task/xrFgR63cw7Nch4vXo/description

You have a list of dishes. Each dish is associated with a list of ingredients used to prepare it. You want to group the dishes by ingredients, so that for each ingredient you'll be able to find all the dishes that contain it (if there are at least 2 such dishes).

Return an array where each element is a list with the first element equal to the name of the ingredient and all of the other elements equal to the names of dishes that contain this ingredient. The dishes inside each list should be sorted lexicographically. The result array should be sorted lexicographically by the names of the ingredients in its elements.

Example

For
  dishes = [["Salad", "Tomato", "Cucumber", "Salad", "Sauce"],
            ["Pizza", "Tomato", "Sausage", "Sauce", "Dough"],
            ["Quesadilla", "Chicken", "Cheese", "Sauce"],
            ["Sandwich", "Salad", "Bread", "Tomato", "Cheese"]]
the output should be
  groupingDishes(dishes) = [["Cheese", "Quesadilla", "Sandwich"],
                            ["Salad", "Salad", "Sandwich"],
                            ["Sauce", "Pizza", "Quesadilla", "Salad"],
                            ["Tomato", "Pizza", "Salad", "Sandwich"]]
For
  dishes = [["Pasta", "Tomato Sauce", "Onions", "Garlic"],
            ["Chicken Curry", "Chicken", "Curry Sauce"],
            ["Fried Rice", "Rice", "Onions", "Nuts"],
            ["Salad", "Spinach", "Nuts"],
            ["Sandwich", "Cheese", "Bread"],
            ["Quesadilla", "Chicken", "Cheese"]]
the output should be
  groupingDishes(dishes) = [["Cheese", "Quesadilla", "Sandwich"],
                            ["Chicken", "Chicken Curry", "Quesadilla"],
                            ["Nuts", "Fried Rice", "Salad"],
                            ["Onions", "Fried Rice", "Pasta"]]
Input/Output

[time limit] 4000ms (go)
[input] array.array.string dishes

An array of dishes. dishes[i] for each valid i contains information about the ith dish: the first element of dishes[i] is the name of the dish and the following elements are the ingredients of that dish. Both the dish name and the ingredient names consist of English letters and spaces. It is guaranteed that all dish names are different. It is also guaranteed that ingredient names for one dish are also pairwise different.

Guaranteed constraints:
1 ≤ dishes.length ≤ 500,
2 ≤ dishes[i].length ≤ 10,
1 ≤ dishes[i][j].length ≤ 50.

[output] array.array.string

The array containing the grouped dishes.
*/
import "sort"
func groupingDishes(dishes [][]string) [][]string {

    ingredientCount := make(map[string]map[string]int) // map each ingredient to the dishes it belongs to
    
    for _, dish := range dishes {
        for _, ingredient := range dish[1:] {
            if ingredientCount[ingredient] == nil {
                ingredientCount[ingredient] = make(map[string]int)
            }
            ingredientCount[ingredient][dish[0]]++
        }
    }
    sortedIngredients := make([]string, 0)
    for k := range ingredientCount {
        sortedIngredients = append(sortedIngredients, k)
    }
    sort.Strings(sortedIngredients)    
    
    res := make([][]string, 0)
    n := 0
    for _, v := range sortedIngredients {
        if len(ingredientCount[v]) > 1 {
            res = append(res, []string{v})
            sortDishes := make([]string, 0)
            for k := range ingredientCount[v] {
                sortDishes = append(sortDishes, k)
            }
            sort.Strings(sortDishes)
            for _, x := range sortDishes {           
                res[n] = append(res[n], x)
            }
            n++
        }
    }   
    return res 
}
