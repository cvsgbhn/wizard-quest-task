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