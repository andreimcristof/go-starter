# Stacks

## What we learn

1. What is a stack
2. How to implement one
3. A concrete usage of stacks in a simulated undo/redo mechanism

## Theory - What is a stack

- an abstract data type ( an ADT, abstract data type = a definition of an interface, AKA a minimal expected behavior.
- It is called abstract because it does not say how the implementation must be done (like data structures do), but it just expresses some expected behavior).
- the expected behavior of a Stack is a LIFO mechanism (last in, first out). Like a stack of plates put on top of each other, you can only operate on the top level - you can insert and remove only at the top of the stack.

## How to implement

- Possible implementations of Stack: by using arrays (slices in Golang), or linked lists.

  - Today we demonstrate by implementing a stack with both methods
  - we use a real life example of an UNDO/REDO mechanism controlled by 2 stacks - we have a program such as well-known Photoshop which offers the UNDO/REDO mechanism.

- We will implement the methods on the stack:
  - push - to add an item on top
  - pop - to remove an item from top
  - peek - to inspect the top most item
