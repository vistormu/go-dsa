# queues

a queue is a data structure that stores a collection of elements in a linear order. it is First In First Out (FIFO) by nature, meaning that the first element added to the queue will be the first one to be removed. this is similar to a line of people waiting for service, where the first person in line is the first one to be served.

these are the basic operations that can be performed on a queue:
- `enqueue`: add an element to the end of the queue
- `dequeue`: remove and return the element at the front of the queue
- `peek`: return the element at the front of the queue without removing it
- `isEmpty`: check if the queue is empty
- `length`: return the number of elements in the queue
- `clear`: remove all elements from the queue

## `queuelinkedlist.go`

it uses a linked list to store the elements of the queue.

pros:
- dynamic size: the queue can grow and shrink as needed, without the need to resize an array
- constant time complexity for enqueue and dequeue operations: both operations take O(1) time, regardless of the size of the queue
- no wasted space: the queue only uses as much memory as it needs to store the elements

cons:



## `queuearray.go`

it uses an array to store the elements of the queue.


## `queuearrayring.go`

it uses a circular array to store the elements of the queue. a circular array is an array that wraps around when it reaches the end, allowing for efficient use of space and time. it needs a fixed size.


## `queuelinkedlistring.go`

it uses a circular linked list to store the elements of the queue. a circular linked list is a linked list where the last node points back to the first node, allowing for efficient use of space and time. it is unbounded and dynamic in size.
