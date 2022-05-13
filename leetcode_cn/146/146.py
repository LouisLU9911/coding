from typing import Dict, Optional


class Node:

    def __init__(self, _key=None, _value=None, _pre=None, _next=None):
        self.key = _key
        self.value = _value
        self.pre = _pre
        self.next = _next

    def __repr__(self):
        return f"<{self.key},{self.value}>{self.next}"


class LRUCache:

    def __init__(self, capacity: int):
        self._cache: Dict[int, Node] = dict()
        self.head: Node = None
        self.tail: Node = None
        self.capacity = capacity

    def __repr__(self):
        return f"<{self.head}>"

    def _update(self, key: int, value: Optional[int] = None):
        if value is not None:
            self._cache[key].value = value
        node = self._cache[key]

        # remove this node
        if node.pre is None:
            return
        else:
            pre = node.pre
        if node.next is not None:
            node.next.pre = pre
        else:
            self.tail = pre
        pre.next = node.next
        # print(f"after rm: {self}")

        # move this node to head
        node.pre = None
        node.next = self.head
        self.head.pre = node
        self.head = node
        # print(f"move to head: {self}")

    def _append(self, key: int, value: int):
        if self.tail is None:
            self.tail = Node(key, value)
            self.head = self.tail
        else:
            self.tail.next = Node(key, value, self.tail)
            self.tail = self.tail.next
        self._cache[key] = self.tail

    def _pop(self) -> Node:
        if self.tail is None:
            return
        key = self.tail.key
        node = self._cache[key]
        self._cache.pop(key)
        if node.pre is None:
            self.head = None
            self.tail = None
        else:
            self.tail = node.pre
            self.tail.next = None
        node.next = None
        node.pre = None
        return node

    def get(self, key: int) -> int:
        if key in self._cache:
            self._update(key)
            res = self._cache[key].value
        else:
            res = -1
        print(self)
        return res

    def put(self, key: int, value: int) -> None:
        if key in self._cache:
            self._update(key, value)
        else:
            if len(self._cache) == self.capacity:
                self._pop()
            self._append(key, value)
            self._update(key)
        print(self)


# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)
