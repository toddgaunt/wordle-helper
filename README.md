# Wordle Helper

[![Go](https://github.com/toddgaunt/wordle-helper/actions/workflows/go.yml/badge.svg)](https://github.com/toddgaunt/wordle-helper/actions/workflows/go.yml)

A command-line utility for solving the [wordle](https://www.nytimes.com/games/wordle/index.html) game by the New York Times.
This program can help a person find the solution to a wordle puzzle by providing the prior guesses entered into the puzzle.
Using these prior guesses, the world-helper will output a list of possible solutions to pick from.

# Example
In this scenario, three guesses have already been made.
In order to find the solution,
these guesses can be passed to wordle-helper on the command-line as arguments.
Each argument consists of two parts delimited with a colon.
The first part of the argument is the guessed word itself, such as the word "slate".
The second part of the argument, following the delimiting colon, is a series of underscores, hyphens, and plus-signs that correlate to the same index as the letter of the word in the first part of the argument.
Underscores imply the letter at the same index in the guessed word aren't in the solution, unmatched letters.
Hyphens imply the letter at the same index in the guessed word are in the solution, matched letters.
Plus-signs imply the letter at the same index of the guessed word are in the solution at that exact position, matched positions.
```
% ./wordle-helper raise:___-_ mount:_---+ donut:_--++
The solution is snout!
%
```
