lst = [3,4,2,1,6,2,1,4,5,7,8,9,0]

def quicksort(l, r):
    if r-l > 1:
        pivot = partition(l, r)
        quicksort(l, pivot)
        quicksort(pivot, r)

def partition(l, r):
    idx = l+1
    for i in range(l+1, r):
        if lst[i] < lst[l]:
            lst[i], lst[idx] = lst[idx], lst[i]
            idx += 1
    lst[idx-1], lst[l] = lst[l], lst[idx-1]
    return idx

quicksort(0, len(lst))
print lst
