# codeforces.com/problemset/problem/430/B

import copy


def list_sum(matrix):
    result = []
    aux = matrix[0]
    for x in range(1, len(matrix)):
        if(aux[0] == matrix[x][0]):
            aux[1] += matrix[x][1]
        else:
            result.append(aux)
            aux = matrix[x]
    else:
        result.append(aux)

    return result


def verif(posib, matrix):
    del(matrix[posib])
    matrix = list_sum(matrix)
    print(matrix)

    while True:
        pos = ispos(matrix)
        if len(pos) == 0:
            break
        else:
            for x in pos:
                del(matrix[x])
            matrix = list_sum(matrix)

        print(matrix)

    return matrix


def ispos(matrix):
    posibilidades = []
    for x in range(len(matrix)):
        if matrix[x][1] >= 3:
            posibilidades.append(x)
    return posibilidades


n, colors, boll = [int(x) for x in input().split(' ')]
lista = [int(x) for x in input().split(' ')]

soma = []
aux = [lista[0], 1]
for x in range(1, n):
    if lista[x] == aux[0]:
        aux[1] += 1
    else:
        soma.append(aux)
        aux = [lista[x], 1]
else:
    soma.append(aux)

posibilidades = []
for x in range(len(soma)):
    if soma[x][0] == boll and soma[x][1] >= 2:
        posibilidades.append(x)

if len(posibilidades) == 0:
    print('0')
else:
    result = 0
    for x in posibilidades:
        matrix = copy.deepcopy(soma)
        del(matrix[x])
        matrix = list_sum(matrix)

        while True:
            pos = ispos(matrix)
            if len(pos) == 0:
                break
            else:
                for i in pos:
                    del(matrix[i])
                if len(matrix) > 0:
                    matrix = list_sum(matrix)

        aux = 0
        for j in matrix:
            aux += j[1]
        aux = n - aux
        if aux > result:
            result = aux

    print(result)
