inpt = [1, 4, -200, 1, 7, 8, 3, 5, 91, 21]

def mergesort(l):
    mid = len(l) // 2
    if mid is 0:
        return l
    else:
        left = mergesort(l[mid:])
        right = mergesort(l[:mid])
        return merge(left, right)

def merge(left, right):
    # Nontypical mergesort because of constant time access to end of list
    l = []
    while left and right:
        if left[-1] < right[-1]:
            l.append(right.pop())
        else:
            l.append(left.pop())

    while left:
        l.append(left.pop())
    while right:
        l.append(right.pop())

    l.reverse()
    return l


print mergesort(inpt)
