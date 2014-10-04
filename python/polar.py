import math
import sys

def angle(a, b):
    angle = math.atan2(a, b)
    if angle < 0:
        angle += math.pi
    return angle

def mag(a, b):
    return math.sqrt(a**2 + b**2)

def ins(a, b, l):
    ang = angle(a, b)
    i = 0
    while ang < l[i]:
        i += 1
    angl, al, bl = l[i]
    if ang == angl:
        if mag(a, b) > mag(al, bl):
            l.insert(i, (ang, a, b))
        else:
            l.insert(i-1, (ang, a, b))
        return
    l.insert(i+1, (ang, a, b))
    l.append((ang, a, b))
    return

l = []
lines = sys.stdin.readlines()
lines.pop(0)

for line in lines:
    a, b = map(int, line.split(' '))
    ins(a, b, l)

for line in l:
    ang, a, b = line
    print str(a) + ' ' + str(b)
