def fib(a, b, n):
    count = 2
    x = list()
    while count < n:
        x.append(b**2+a)
        a = b
        b = x[count-2]
        count += 1
    return x.pop()

x = raw_input().split(' ')
a, b, c = [int(n) for n in x]
print fib(a, b, c)
