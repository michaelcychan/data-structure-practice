# Sorting comparison

## Simple comparison
| Method | Best Case | Worst Case | Average Case | Space Complexity |
| :---: | :---: | :---: | :---: | :---: |
| Bubble Sort | Ω(n) | O(n^2) | O(n^2) | O(1) |
| Merge Sort | Ω(n log n) | O(n log n) | O(n log n) | O(n) |
| Quick Sort | Ω(n log n) | O(n^2) | O(n log n) | O(log n) |

## Quicksort
Quicksort is an efficient algorithm. Ideally, this process of dividing the array will produce sub-lists of nearly equal length, otherwise, the runtime of the algorithm suffers.

There is no fixed view on whetheer quicksort is in-place sort. Most believes it isn't because it take the space complexity of O(log n) instead of O(1).

## Merge Sort
Merge sort is a divide-and-conquer sorting algorithm that breaks the list-to-be-sorted into smaller parts. In a divide-and-conquer algorithm, the data is continually broken down into smaller elements until sorting them becomes incredibly simple.  

It’s important to note that **merge sort makes a copy of the entire list** during its process, meaning it does not sort in place, which adds to the space complexity.

## Why this space complexity?
**Quicksort**: For quicksort, recursion requiring O(log(n)) space is correct. Since quicksort calls itself on the order of log(n) times (in the average case, worst case number of calls is O(n)), at each recursive call a new stack frame of constant size must be allocated. Hence the O(log(n)) space complexity.

**Mergesort**: Since mergesort also calls itself on the order of log(n) times, why the O(n) space requirement? The extra space comes from the merge operation. Most implementations of merge use an auxiliary array with length equal to the length of the merged result, since in-place merges are very complicated. In other words, to merge two sorted arrays of length n/2, most merges will use an auxiliary array of length n. The final step of mergesort does exactly this merge, hence the O(n) space requirement.