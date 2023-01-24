# Trees
This is an exercise on Tree data structure implemented with Go.  

## Tree Node
Each Tree node has `value` to store an integer and a `children`, a slice storing the pointers of its children. Because it is a slice, a tree node can have multiple number of children.  

`GetValue()` and `GetChildren()` are used to retrieve the current value and children of each node; `SetValue()` is used to set or amend the value; and `AddChild()` is used to add a new child node to a tree node.  

`PrintTree()` can be used to print the whole tree.  

`BreadthFirstSearch()` and `DepthFirstSearch()` are implemented to search if a target is stored in the tree.

## BreadthFirstSearch
It is the level-order search. `BreadthFirstSearch()` is implemented iteratively, that means the children of each node is stored, by using the FIFO pricinple, each node is visited level by level.  

The function returns the path to target, or returns an error if target not found.  

## DepthFirstSEarch
This searches as far as possible along each branch before backtracking. `DepthFirstSearch()` is implemented recursively. An additional path parameter is used to store the current path. To follow the Interface Segregation Principle, `DepthFirstSearch()` actually wraps around a `Dfs()` so that users do not need to place an empty path parameter in the outer function. The path is taken care of by design.

It returns the path to the target, or returns an error if target not found.  