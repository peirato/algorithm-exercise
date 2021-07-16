# -*- coding: utf-8 -*-

from typing import Optional


class ArrayQueue:

    def __init__(self, capacity: int):
        self._capacity = capacity
        self._items: list = []
        self._head: int = 0
        self._tail: int = 0

    def enqueue(self, item: str) -> bool:

        if self._tail == self._capacity:
            if self._head == 0:
                return False
            else :
                for i in range(0,self._tail-self._head):
                    self._items[i] = self._items[i+self._head]
                self._tail = self._tail - self._head
                self._head = 0

        self._items.insert(self._tail,item)
        self._tail += 1

    def dequeue(self) -> Optional[str]:
        if self._tail == self._head:
            return None
        else :
            item :str = self._items[self._head]
            self._head += 1
            return item
    def __repr__(self) -> str:
        return " ".join(item for item in self._items[self._head : self._tail])

if __name__ == "__main__":
    queue: ArrayQueue = ArrayQueue(10)
    
    for ii in range(1,3):
        for i in range(1,10):
            queue.enqueue(str(i))
            print(queue)

        for i in range(1,11):
            queue.dequeue()
            print(queue)
