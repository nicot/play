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
    # Not actually linear time because popping from the left side
    # takes linear time
    l = []
    while left:
        if right:
            if left[0] > right[0]:
                l.append(right.pop(0))
            else:
                l.append(left.pop(0))
        else:
            l.append(left.pop(0))

    while right:
        l.append(right.pop(0))

    return l

print mergesort(inpt)
