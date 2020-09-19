# https://www.urionlinejudge.com.br/judge/en/problems/view/1120

while True:
    d, m = input().split(' ')
    if int(d) == 0 and int(m) == 0:
        break

    result = '0'
    for i in m:
        if i != d:
            result += i

    print(int(result))
