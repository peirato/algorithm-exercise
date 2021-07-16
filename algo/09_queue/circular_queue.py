from typing import Optional
from itertools import chain


class CircularQueue:

    def __init__(self, capacity):
        self._items :list = []
        self._capacity = capacity
        self._head = 0
        self._tail = 0

    def enqueue(self, item: str) -> bool:
        if (self._tail + 1) % self._capacity == self._head:
            return False
        self._items.append(item)
        self._tail = (self._tail + 1) % self._capacity

    def dequeue(self) -> Optional[str]:
        pass

    def __repr__(self) -> str:
        # return '-'.join(self._items[self._head:self._tail])
        return '-'.join(self._items)


if __name__ == "__main__":
    queue = CircularQueue(10)
    for i in range(1,20):
        queue.enqueue(str(i))
        print(str(i) + ':' + str(queue))
