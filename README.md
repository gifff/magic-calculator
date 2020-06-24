# Magic Calculator

## Problem Statement

Please write a program that handles the following use-cases:

- [ ] Sum X & Y, and print the result
  - Input : 1, 2
  - Output : 3

- [ ] Multiply X & Y, and print the result
  - Input : 1, 2
  - Output: 2

- [ ] Find first N prime number, and print the result
  - Input: 4
  - Output : 2, 3, 5, 7

- [ ] Find the first N Fibonacci sequence, and print the result
  - Input: 4
  - Output : 0, 1, 1, 2

### Assumption

1. For case of Sum and Multiply, input values are integer and may be negative
   number ranging from -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
   (inclusive) (int64 range).
2. For Find first N prime number and Find the first N Fibonacci sequence cases,
   N must be integer that is larger than or equal to 0. It will output empty
   when N is 0, and panic if negative.
3. The result is not guaranteed to always be correct due to int64 limitation to
   hold the result. It might be overflow or underflow.
   For example, the limitation for N in Fibonacci case is 93, because on the
   94th case, the result is overflowed.

## System Requirements

1. Go 1.14.x

## How to run test

```shell
$ go test ./...
```

## How to run the program

First, you have to build the program

```shell
$ go build
```

Then, execute the compiled binary. For Linux/Unix, the command is:

```shell
$ ./magic-calculator
```
