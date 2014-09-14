# Take an array an array of heights and find where water would collect.

tray = [1, 2, 5, 1, 7, 4, 1, 2, 4, 1, 6]
right = []
left = []
level = []
water = 0

#def sweep(l):

for idx, val in enumerate(reversed(tray)):
    if idx == 0 or val > right[idx-1]:
        right.append(val)
    else:
        right.append(right[idx-1])
right.reverse()

for idx, val in enumerate(tray):
    if idx == 0 or val > left[idx-1]:
        left.append(val)
    else:
        left.append(left[idx-1])

for i, j in zip(left, right):
    if i > j:
        level.append(j)
    else:
        level.append(i)

for l, t in zip(level, tray):
    water += l - t

print water
