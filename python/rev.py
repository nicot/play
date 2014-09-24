inp = "Hello, World!"

def rev(s):
    s = list(s)
    l = len(s)
    for i in range(l//2):
        end = l - i - 1
        s[i], s[end] = s[end], s[i]

    return ''.join(s)

print rev(inp)
