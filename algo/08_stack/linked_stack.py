"""
    Stack based upon linked list
    基于链表实现的栈
"""

from typing import Optional

class Node:

    def __init__(self,data:int,next=None) -> None:
        self._data = data
        self._next = next

class LinkedStack:

    def __init__(self) -> None:
        self._top:Node = None

    def push(self, value: int):
        top = self._top
        node = Node(value)
        self._top = node
        node._next = top

    def pop(self) -> Optional[int]:
        top = self._top
        self._top = self._top._next
        return top

    def __repr__(self) -> str:
        current = self._top
        nums = []
        while current:
            nums.append(current._data)
            current = current._next
        return " ".join(f"{num}]" for num in nums)


if __name__ == "__main__":
    stack = LinkedStack()
    for i in range(9):
        stack.push(i)
    print(stack)
    for _ in range(3):
        stack.pop()
    print(stack)