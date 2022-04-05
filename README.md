# Wordle Solver

A CLI utility for solving the game
[wordle](https://www.nytimes.com/games/wordle/index.html) by the New York
Times. wordle-solver allows one to find the solution to a wordle puzzle by
entering in their prior guesses to come up with the solution or a set of
possible solutions.

# Example
In this scenario you have already made three guesses. In order to find the
solution, they can be passed to wordle-solver on the command-line as arguments
with colons to delimit the parts. The first part of an argument is the guessed word.
The second part of the argument is the set of matched letters and matched positions
corresponding to the guessed word. In the second part of the argument, a '_' character
means the letter was unmatched, a '-' means the letter was matched but in the wrong
position, and a '+' means the letter was matched and in the correct position.
```
% ./wordle-solver raise:___-_ mount:_---+ donut:_--++
The solution is snout!
%
```
