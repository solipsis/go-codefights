/**
Challenge: CrosswordFormation
https://codefights.com/arcade/code-arcade/labyrinth-of-nested-loops/W5Sq7taLSzNHh8GiF/description

You're a crossword fanatic, and have finally decided to try and create your own. However, you also love symmetry and good design, so you come up with a set of rules they should follow:

the crossword must contain exactly four words;
these four words should form four pairwise intersections;
all words must be written either left-to-right or top-to-bottom;
the area of the rectangle formed by empty cells inside the intersections isn't equal to zero.
Given 4 words, find the number of ways to make a crossword following the above-described rules. Note that two crosswords which differ by rotation are considered different.

Example

For words = ["crossword", "square", "formation", "something"], the output should be
crosswordFormation(words) = 6.

The six crosswords can be formed as shown below:

  f                         f                             f
  o                     c r o s s w o r d     c r o s s w o r d
c r o s s w o r d           r   o                   q     r
  m   q                     m   m                   u     m
  a   u            ;  s q u a r e          ;        a     a
  t   a                     t   t                   r     t
  i   r                     i   h             s o m e t h i n g
s o m e t h i n g           o   i                         o
  n                         n   n                         n
                                g                               
                                                              
    c         s               s                                      
f o r m a t i o n       c     q               c         s          
    o         m         r     u               r         o      
    s q u a r e       f o r m a t i o n       o         m            
    s         t    ;    s     r            ;  s q u a r e                  
    w         h         s o m e t i n g       s         t         
    o         i         w                     w         h     
    r         n         o                   f o r m a t i o n            
    d         g         r                     r         n         
                        d                     d         g             
Input/Output

[time limit] 4000ms (go)
[input] array.string words

An array of distinct strings, the words you need to use in your crossword.

Guaranteed constraints:
words.length = 4,
3 ≤ words[i].length < 15.

[output] integer

The number of ways to make a correct crossword of the desired formation.
**/
func crosswordFormation(words []string) int {

    // Should probably have built a graph and done a DFS :D
    // but slight pruning gets it under the time limit 
    c := 0
    // for each of 4 words
    for _, w1 := range words {
        // for each of 3 remaining y starting at top
        for _, w2 := range words {
            if w2 == w1 {continue}
            // typewriter left // possibly put characters in word in map
            for x2 := 0; x2 < len(w2)-2; x2++ {
                for y2 := 0; y2 < len(w1)-2; y2++ {
                    if w1[y2] != w2[x2] {continue}
                    // for each of 2 remaining
                    for _, w3 := range words {
                        if w3 == w1 || w3 == w2 {continue}
                        for x3 := 0; x3 < len(w3)-2; x3++ {
                            for y3 := y2+2; y3 < len(w1); y3++ {                            
                                if w1[y3] != w3[x3] {continue}                                
                                // for last word
                                for _, w4 := range words {
                                    if w4 == w1 || w4 == w2 || w4 == w3 {continue}
                                    for x4 := 2; x4+x2 < len(w2) && x4+x3 < len(w3); x4++ {
                                        for y4 := 0; y4 + (y3-y2) < len(w4); y4++ { // fix
                                            // top left joint
                                            if w1[y2] != w2[x2] {continue}
                                            // bottom left
                                            if w1[y3] != w3[x3] {continue}
                                            // top right
                                            if w2[x4+x2] != w4[y4] {continue}
                                            // bottom right
                                            if w3[x4+x3] != w4[y4+(y3-y2)] {continue}             
                                            c++
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }    
    return c        
}

