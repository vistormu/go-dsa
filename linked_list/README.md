# linked list

a linked list is a data structure that consists of a sequence of elements, each of which contains a reference (or link) to the next element in the sequence. 

linked lists include the following operations:
- `append`   O(1): add an element to the end of the list
- `prepend`  O(1): add an element to the beginning of the list
- `pop`      O(1): remove the first element from the list
- `pop_last` O(n): remove the last element from the list
- `insert`   O(n): add an element at a specific index
- `remove`   O(n): remove an element from the list
- `get`      O(n): get an element at a specific index
- `length`   O(1): get the length of the list

## `linked_list.go`

## `doubly_linked_list.go`

a double linked lists maintains a reference to the previous element in the sequence as well as the next element. this allows for more efficient insertion and deletion of elements, as well as traversal in both directions.

now, `pop_last` operation is O(1) because we can access the last element directly.

the main drawback of doubly linked lists is that they require more memory to store the additional reference to the previous element. this can be a significant overhead if the list is very large.

## `circular_linked_list.go`
