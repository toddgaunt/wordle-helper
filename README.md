# Wordle Solver

A CLI utility for solving the game
[wordle](https://www.nytimes.com/games/wordle/index.html) by the New York
Times. wordle-solver allows one to find the solution to a wordle puzzle by
entering in their prior guesses to come up with the solution or a set of
possible solutions.

# Example
In the scenario of a game where you have already made three guesses. They can be
passed as arguments delimited by colons. The first part of an argument is the guess word.
The second part of the argument is the set of matched letters. The third part of an
argument is the matched letter positions.
```
% ./wordle-solver raise:s mount:ount:____t donut:onut:____t
The solution is snout!
%
```
