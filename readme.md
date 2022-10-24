# Master Wizard Quest
### (Test task for Pixelmatic)

Seba the Master Wizard is stuck in a maze filled with miscellaneous demons and hungry animals. He
has to find the way out as quickly as possible! Luckily he has a map of the entire maze. He wishes
there was a magic trick available to tell him the right way out of the maze.

Each room in the maze is represented either as a string describing what is in the room or as a
dictionary of the possible ways to get out of the room. The string describing what is in the room will
either be some demon or hungry animal (“dragon”, “tiger”, “ogre”, or similar) or the word “exit”,
meaning that you have found the exit. The dictionary of ways to get out of the room will have keys
like “forward”, “left”, “right” and similar, and the value for each key will be the representation of a
new room.

## Examples
1. Small Maze
   Input
   `{“forward”: “tiger”, “left”: {“forward”: {“upstairs”: “exit”}, “left”: “dragon”}, “right”: {“forward”:
   “dead end”}}`
   Output
   `[“left”, “forward”, “upstairs”]`
2. Exit
   Input
   `{“forward”: “exit”}`
   Output
   `[“forward”]`
3. No Exit
   Input
   `{“forward”: “tiger”, “left”: “ogre”, “right”: “demon”}`
   Output
   `"Sorry"`

**Note**: The response is a list of strings showing how to get out of the maze or the word “sorry” if that
is impossible.

## Expected work

● Maze Solver: Propose a function to find the **shortest way** out of any maze given in arguments
using Golang

● Maze Generator (Bonus): Explain in one or two paragraphs how to generate an infinity of
maze. Think of exceptions and weird cases

# Solution for Maze Solver
> Used BFS algorithm to find the shortest path

## How to test
`go test -v`

# Maze Generator
Suppose, we have a set of directions `{forward, left, right, upstairs, downstairs, behind the tree, ...}` and a set of endings `{exit, dragon, pool with sharks, ...}`

To start, I will generate a first chunk of infinity of maze. As current maze is basically a tree, I would need an algorithm for that, for example [this one](https://www.geeksforgeeks.org/random-tree-generator-using-prufer-sequence-with-examples/). Each node can be a struct like this, for example:
```go
type Path struct {
	Direction string
	Ending string
	NextPath *Path // there cannot be ending and nexpath at one time
}
```

Also it makes sence to add a rules:
1. There must be at least one Path with ending "exit" (or "borderline"). When user hits this ending - generate next chunk of a maze.
2. The paths from one 'parent' must be distinct.