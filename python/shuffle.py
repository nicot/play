import random
l = [1,5,2,5,3,2]
for i, val in enumerate(l):
    rand = random.randint(i, len(l)-1)
    l[i], l[rand] = l[rand], l[i]

print l
