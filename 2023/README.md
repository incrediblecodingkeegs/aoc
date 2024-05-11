# 2023

[Reddit Solution thread](https://www.reddit.com/r/adventofcode/search?q=flair_name%3A%22SOLUTION%20MEGATHREAD%22&restrict_sr=1)

## Day 1

- I used regexes even though I knew it would be dirty (spoiler: it was)
- Found issues with regexes that they are mostly linear in nature and created issues with how to parse the string, 
primarily that if a number word trailed off of another one, it had real trouble with this (e.g. oneight)
- Saw a neat solution on reddit where they did a string replace (one -> one1one, two -> two2two etc.)
- Much neater golang solution [here](https://github.com/mnml/aoc/blob/main/2023/01/1.go) that uses string replace

## Day 2

- I used my learnings from Day 1, and made a much neater solution using string replaces
- Was able to very quickly retrofit the program for the Part 2 solution


## Day 3

- Got stuck on this one for aaaages.
- All my unit tests were passing, passed the visual inspection test.  Was the way I was processing results that was the 
problem in the end

## Day 4

- Completed this day using a singly linked list.  My first time using this data structure, but I think it handled the
task well.
- Interestingly, I saw that other solutions used lists to store only the sum of all cards matched, which is probably 
a simpler solution given the part 2 request.  If part 2 had have been more complex though, this might have run into
space complexity issues

## Day 5

- Needed to improve the space complexity of my program to allow it to execute the second part
- Used a state machine to move through the different transitions, although my solution was dependent on the input file
being 100% correct
- Solutions used go routines to improve the execution time, I should look at doing this as well

## Day 6

## Day 7

- Went the wrong way to begin with
- Approach I took in the end could have been simplified, but writing a custom sort function for each of the rank slices
proved valuable when getting to part 2

## Day 8

- My biggest issue with Day 8 was forgetting that lowest common multiple existed
- Started with the most basic solution, a loop that ran through all nodes and their paths and cross checked whether
they had a common step count in their traversal
- Found that this approach was taking too long, and so wrote the program to be concurrent for each different starting node,
although after running for 2 hours, figured even this was too slow
- Then discovered that all paths were looping, and using a lowest common multiple for each of the paths, found the answer
in < 1s of execution time

## Day 9

- Didnt have too much trouble with this one, my logic for the first part was easy to modify to work for the second part