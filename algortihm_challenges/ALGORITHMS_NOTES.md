# Notes about Algorithms
- Algorithm complexity/Computational complexity - amount of time and resources required
to run a function of the input size.
- It measures how the running time of an function grows with the size of the input
- It's expressed in terms of big O notation

### Big O Notation
- Describes time or space complexity
- Describes worst case scenario (upper limit) of the running time or space used by algorithms as the input size increases
- n represents the input size
- Helps determine the scalability of an algorithm, because helps to understand how the performance of the algorithm
will be affected as the input size grows.
    - It does this by:
      - Analysing the worst-case time or space 
- 
- O(n) - linear growth rate - meaning that as the input size grows,
the running time of the algorithm will increase linearly with the size of the input
  - iterate over the input data a single time. 
  - Examples include linear search and some sorting algorithms like counting sort.


- O(n^2) - quadratic growth rate - meaning that the running time of the algorithm will  increase rapidly.
  - Typically, have nested loops that iterate over the input data multiple times.
increase exponentially with the size of the input.


- O(log n) - logarithmic growth rate - meaning that the running time of the algorithm will
increase logarithmically with the size of the input.
    - Divides the input size into smaller and smaller pieces in each iteration, resulting in a logarithmic growth rate.
    - This is often seen in algorithms that use divide-and-conquer strategies
  
- O(1) - algorithmic complexity is constant and independent of the input size n
  - accessing an element in an array or hash table
  - very efficient and scales extremely well for large inputs.
  - the best complexity for algorithms

### Stack

- a stack is a data structure that follows the Last-In-First-Out (LIFO) principle.
- is a collection of elements, where elements can be added or removed only from one end, known as the top of the stack.
- supports two main operations: 
  - push adds an element to the top of the stack
  - pop removes the top element from the stack.
- Other operations on a stack:
  - peek which returns the top element without removing it
  - isEmpty, which checks if the stack is empty
- Stacks can be implemented using various data structures such as arrays or linked lists.
- Stacks are commonly used in computer algorithms for tasks such as:
  - parsing expressions
  - evaluating arithmetic expressions
  - backtracking in algorithms such as depth-first search.
- Stacks have a limited capacity, which means that it is possible to overflow or underflow the stack.
  - an overflow condition occurs when a push operation is performed on a full stack.
  - an underflow condition occurs when a pop operation is performed on an empty stack.

#### Different types of stacks

- Static stack: 
  - Is implemented using a fixed-size array.
  - The size of the stack remains constant throughout the program execution.
    - Pros:
      - Simple and easy to implement
      - Efficient in terms of time complexity
      - Does not require memory allocation during program execution
    - Cons:
      - Limited capacity due to the fixed size of the array
      - Inefficient in terms of memory usage if the stack is not fully utilized

- Dynamic stack: 
  - Is implemented using a dynamic data structure such as a linked list
  - Its size can be adjusted dynamically during program execution.
    - Pros:
      - Dynamic size, allowing for flexibility in the size of the stack
      - Efficient in terms of memory usage, as only the required memory is used
      - Allows for resizing of the stack during program execution
    - Cons:
      - Less efficient in terms of time complexity compared to the static stack
      - More complex to implement compared to the static stack
      - Memory allocation and deallocation during program execution can lead to memory fragmentation and slower performance.

- Array-based stack: 
  - This type of stack is implemented using an array to store the stack elements.
  - Is operations such as push and pop are performed by manipulating the indices of the array. 
    - Pros:
      - Simple and easy to implement
      - Efficient in terms of time complexity
    - Cons:
      - Limited capacity due to the fixed size of the array
      - Inefficient in terms of memory usage if the stack is not fully utilized

- Linked list-based stack: 
  - This type of stack is implemented using a linked list.
  - Each node in the linked list stores the stack element and a pointer to the next node.
    - Pros:
      - Dynamic size, allowing for flexibility in the size of the stack
      - Efficient in terms of memory usage, as only the required memory is used
    - Cons:
      - Less efficient in terms of time complexity compared to the array-based stack
      - More complex to implement compared to the array-based stack

- Circular stack: 
  - This type of stack is implemented using a circular buffer
  - The last element of the stack is connected to the first element, forming a circle.

- Priority stack: 
  - This type of stack is used to store elements in order of their priority.
  - The element with the highest priority is always placed at the top of the stack.
  - Popping an element removes the element with the highest priority.
    - Pros:
      - Useful when elements need to be stored in order of their priority
      - Enables efficient processing of tasks based on priority levels
    - Cons:
      - Less efficient in terms of time complexity compared to the array-based and linked list-based stacks
      - Priority queues can be complex to implement
