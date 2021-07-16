# -*- coding: UTF-8 -*-

from typing import Optional


class Node:

    def __init__(self, data: str, next=None) -> None:
        self.data = data
        self._next = next


class LinkedQueue:

    def __init__(self):
        self._head: Optional[Node] = None
        self._tail: Optional[Node] = None

    def enqueue(self, value: str):
        node = Node(value)
        if self._tail:
            self._tail._next = node
        else:
            self._head = node
        self._tail = node

    def dequeue(self) -> Optional[str]:
        if self._head:
            node = self._head
            self._head = node._next
            return node

    def __repr__(self) -> str:
        values = []
        current: Node = self._head._next
        while current:
            values.append(current.data)
            current = current._next
        return "->".join(value for value in values)


if __name__ == "__main__":
    q = LinkedQueue()
    for i in range(10):
        q.enqueue(str(i))
    print(q)

    for _ in range(3):
        q.dequeue()
    print(q)

    q.enqueue("7")
    q.enqueue("8")
    print(q)
