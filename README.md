# Go Generic

## Introduction

This project showcases a simple implementation of generic types in Go. The goal is to demonstrate how to create reusable code that can work with different data types without sacrificing type safety or performance.

## Features

1. **Utilities Package**

The `utils` package contains various utility functions that demonstrate the use of generics in Go. These functions include:

- `PrintWithEmoji`: Prints a message with an emoji prefix.
- `PrintWithEmojiOnlyStringNum`: Prints a message only if it's a string or an integer.
- `PrintWithEmojiOnlyNumFloat`: Prints a message only if it's an integer or a float.

2. **Stack Implementation**

The `utils` package also includes a generic stack implementation, which supports storing elements of any type. The stack provides the following functionalities:

- `Push`: Pushes an item onto the stack.
- `Pop`: Pops an item from the top of the stack.
- `Peek`: Retrieves the top item from the stack without removing it.
- `IsEmpty`: Checks if the stack is empty.
- `Size`: Returns the size of the stack.

3. **Student Models**

The `studentmodels` package demonstrates how to create generic structs in Go. Specifically, it includes a `Student` struct that can store a student's name, score, and age. The score can be of any type (string, integer, or float).

## Usage

The `main` package showcases the usage of the generic utilities and student models. It includes examples such as:

- Printing messages with emojis using various data types.
- Finding the minimum and maximum values of integers using generic functions.
- Using the generic stack implementation to store strings and integers.
- Creating student instances with different types of scores and manipulating their scores.

## How to Run

To run the examples, simply execute the main package. Ensure that all the necessary packages are imported correctly and that your Go environment is properly set up.

```bash
go run main.go
```

## Conclusion

This project serves as a basic demonstration of how to leverage generics in Go to create reusable and type-safe code. While generics are still relatively new in Go, they provide a powerful way to write more flexible and efficient code. As the language continues to evolve, the use of generics is expected to become more prevalent in the Go ecosystem.
