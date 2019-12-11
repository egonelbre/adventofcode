# Advent of Code 2019

https://adventofcode.com/2019

In the style of "Cut the Red Wire" (see http://www.carlopescio.com/2011/06/cut-red-wire.html).

The goal is to not do any testing of the code. It should work correctly the first time you run it.

| Day | Part | Result | Comments |
|-----|------|--------|----------|
|   1 |   1  | Success |  |
|     |   2  | Success |  |
|   2 |   1  | Failure | Computer worked correctly, but missed instructions about changing code before executing. Two lines to fix the bug. |
|     |   2  | Success |  |
|   3 |   1  | Success |  |
|     |   2  | Failure | Code was correct, but forgot to call it from main. |
|   4 |   1  | Success |  |
|     |   2  | Failure | Counted half-clusters. |
|   5 |   1  | Success |  |
|     |   2  | Failure | Copy-paste mistake, instead of `==` had `<` in equals implementation. |
|   6 |   1  | Success |  |
|     |   2  | Success |  |
|   7 |   1  | Success |  |
|     |   2  | Failure | Misread the assignment and thought the output should be once the amplifiers finish amplifying instead of running max output from all phases. Also missed that part 2 used different phase values. Also it seems, it was across all permutations of phase inputs. |
|   8 |   1  | Success |  |
|     |   2  | Failure | Forgot to set Width, Height in constructor and in `Index` used `y * image.Height`. |
|   9 |   1  | Failure | Off by one error in setting code length. |
|     |   2  | Success |  |
|  10 |   1  | Failure | Tired. Counted from non-asteroids. Miscalculated prime offsets. |
|     |   2  | Failure | Out-of-bounds. Wrong Angle implementation for the orientation. |
