class Node:
    def __init__(self, value):
        self.value = value
        self.child = None

class Stack:
    def __init(self):
        self.head = Node():

    def pop(self):
        ret = self.head.value
        self.head = self.head.child
        return ret

    def push(self, value):
        insert(self.head, value)

def insert(n, v):
    if n.child is None:
        n.child = Node(v)
    else:
        insert(n.child, v)

def walk(n):
    if n is not None:
        print n.value
        walk(n.child)

root = Node(0)
l = [1,2,3,4,1,4,3,2,6]

for i in l:
    insert(root, i)

walk(root)
