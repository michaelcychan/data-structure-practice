# data-structure-practice
This repo is for me to learn the basic data strcutures

I attempt to learn different data structures by implementing my own versions using TypeScript.

## Basic Linked List
I created a Linked List Class `MyNumLinkedList` to store numbers. Each `MyNumLinkedList` consists of instances of the `MyNumberNode` class.  

Each `MyNumberNode` stores a number and a pointer to the next `MyNumberNode`.  

The `MyNumLinkedList` class has the following functions:  
```
.length()
// to return the length of the list

.addNewHead(`value`)
// to add a new node at head

.removeNode(`value`)
// to remove the value

.reverseIndex(`target`)
// to find the reverse index of the target

.swapping(a, b)
// to swap the position of the targets in the List

.findMid()
// to find the middle node (tail-centered) of the List

```