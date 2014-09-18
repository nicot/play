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

root = Stack()
l = [1,2,3,4,1,4,3,2,6]

for i in l:
    root.push(i)

for i in range(9):
    print root.pop()
