# Queues and Heaps

## Queues

- its an ADT
- it provides a behavior, not an implementation
- collection of entities kept in a sequence
- can be modified:
  - enqueue-ing - adding at the back of the queue
  - dequeue-ing - removing from the front of the queue
- it's FIFO - first in, first out
- Queues are used everywhere where there is a need for entities
  to be stored to be processed later

## PriorityQueue

- we will implement a Priority Queue
  - it's a Queue where each element has an additional "priority" property:
    - in a standard queue, elements exit the queue in the same order they arrived.
    - in a Priority Queue - the element with higher prio is returned before one with lower prio regardless of order of entry in queue.
    - we will use the Heap Data Structure to implement the Priority Queue

## Heaps

- a tree data structure which must satisfy the Heap property:

  - what is the Heap Property: it is based on the type of heap.
    - there are 2 types of Heaps:
      max-heaps, min-heaps.
    - max-heaps: the root node (the node in the tree without a parent) always has the highest value in the tree, and each node must have a value greater than the value of its children.
    - min-heaps: the root node (the node in the tree without a parent) always has the smallest value in the tree, and each node must have a value smaller than the value of its children.
    - when are heaps useful? When you need to repeatedly access or remove the object with the highest value (for max-heaps) or lowest value (for min-heaps).

- then we will implement our own Priority Queue in a system that simulates an Emergency Room in a hospital, where patients are handled in the order of how severe their injury is, instead of FIFO.
