class Node:
    def __init__(self, value):
        self.value = value
        self.child = None

class Stack:
    def __init__(self):
        self.head = None

    def pop(self):
        if self.head is None:
            return None
        ret = self.head.value
        self.head = self.head.child
        return ret

    def push(self, value):
        new = Node(value)
        new.child = self.head
        self.head = new

class Queue:
    def __init__(self):
        self.head = None
        self.tail = None

    def dequeue(self):
        if self.head is None:
            return None
        ret = self.head.value
        self.head = self.head.child
        return ret

    def enqueue(self, value):
        new = Node(value)
        # I'd like this to be in init, but we could have empty queues.
        if self.head is None or self.tail is None:
            self.head, self.tail = new, new
        self.tail.child = new
        self.tail = self.tail.child

root = Queue()
l = [1,2,3,4,1,4,3,2,6]

for i in l:
    root.enqueue(i)

for i in range(9):
    print root.dequeue()
