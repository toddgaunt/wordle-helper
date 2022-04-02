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
The second part of the argument is the set of matched letters. The third part of an
argument is the matched letter positions.
```
% ./wordle-solver raise:s mount:ount:____t donut:onut:____t
The solution is snout!
%
```
