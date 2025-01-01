# Linked List: Key Concepts, Time Complexity, and Space Complexity

## 1. **Singly Linked List**
- **Description**: Each node contains data and a reference to the next node.
- **Time Complexity**:
  - Access: O(n)
  - Search: O(n)
  - Insert/Remove (at head): O(1)
  - Insert/Remove (at tail): O(n) for singly linked list (O(1) with tail pointer)
- **Space Complexity**: O(n) (where n = number of nodes)

---

## 2. **Doubly Linked List**
- **Description**: Each node contains data, a reference to the next node, and a reference to the previous node.
- **Time Complexity**:
  - Access: O(n)
  - Search: O(n)
  - Insert/Remove (at head): O(1)
  - Insert/Remove (at tail): O(1)
- **Space Complexity**: O(n) (due to extra space for previous pointer)

---

## 3. **Circular Linked List**
- **Description**: The last node points back to the first node, forming a circle.
- **Time Complexity**:
  - Access: O(n)
  - Search: O(n)
  - Insert/Remove (at head): O(1)
  - Insert/Remove (at tail): O(1) (if tail pointer is maintained)
- **Space Complexity**: O(n)

---

## 4. **Circular Doubly Linked List**
- **Description**: A doubly linked list where the last node points to the first, and the first node points to the last.
- **Time Complexity**:
  - Access: O(n)
  - Search: O(n)
  - Insert/Remove (at head/tail): O(1)
- **Space Complexity**: O(n)

---

## 5. **Operations on Linked List**
### a) **Insert**
- At head: O(1)
- At tail: O(n) (O(1) with tail pointer)

### b) **Delete**
- At head: O(1)
- At tail: O(n) (O(1) with tail pointer)

### c) **Search**
- O(n) (linear search)

### d) **Traversal**
- O(n)

### e) **Reverse**
- O(n)

---

## 6. **Memory Management**
- Linked lists use dynamic memory allocation, unlike arrays, which use contiguous memory.
- Each node requires extra space for storing pointers.

--- 

## 7. **Applications**
- Dynamic memory allocation
- Efficient insertions/deletions at the ends
- Implementing queues, stacks, graphs, and other data structures.
- basically you can say linked list dont have indexing, but you can access the data at any node by traversing the list.

---

## 8. **Disadvantages**
- Linked lists are not efficient for insertions/deletions at the beginning or middle of the list.
- Linked lists are not efficient for random access.
- Linked lists are not efficient for sorting.


