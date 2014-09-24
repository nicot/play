in1 = "1234"
in2 = "-4321"

def atoi(s):
    s = list(s)
    num = 0
    if s[0] is '-':
        sign = -1
        s.pop(0)
    else:
        sign = 1

    for digit in s:
        digit = ord(digit) - ord('0')
        num *= 10
        num += digit

    return num*sign

print atoi(in1)
print atoi(in2)
