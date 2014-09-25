class BiNode:
    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None

def insert(node, value):
    if node.value > value:
        if node.right:
            insert(node.right, value)
        else:
            node.right = BiNode(value)
    else:
        if node.left:
            insert(node.left, value)
        else:
            node.left = BiNode(value)

def inOrder(node):
    if not node:
        return ""
    else:
        return inOrder(node.right) + str(node.value) + inOrder(node.left)

def bfs(node):
    if not node:
        return ""
    else:
        return str(node.value) + bfs(node.left) + bfs(node.right)

def check(node, ma, mi):
    if not node:
        return True
    if ma > node.value or mi < node.value:
        return False
    return check(node.left, node.value, mi) and\
           check(node.right, ma, node.value)

t = BiNode(0)
l = [1,2,3,4,1,4,3,2,6]

for i in l:
    insert(t, i)

print inOrder(t)
print bfs(t)
print check(t, -200, 200)
