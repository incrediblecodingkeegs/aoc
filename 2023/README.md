# 2023

## Day 1

[Reddit thread](https://www.reddit.com/r/adventofcode/comments/1883ibu/2023_day_1_solutions/?sort=confidence)

After Thoughts:
- I used regexes even though I knew it would be dirty (spoiler: it was)
- Found issues with regexes that they are mostly linear in nature and created issues with how to parse the string, 
primarily that if a number word trailed off of another one, it had real trouble with this (e.g. oneight)
- Saw a neat solution on reddit where they did a string replace (one -> one1one, two -> two2two etc.)
- Much neater golang solution [here](https://github.com/mnml/aoc/blob/main/2023/01/1.go) that uses string replace

## Day 2

After Thoughts:
- I used my learnings from Day 1, and made a much neater solution using string replaces
- Was able to very quickly retrofit the program for the Part 2 solution


## Day 3

After thoughts:
- Got stuck on this one for aaaages.
- All my unit tests were passing, passed the visual inspection test.  Was the way I was processing results that was the 
problem in the end