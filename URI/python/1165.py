
# https://www.urionlinejudge.com.br/judge/pt/problems/view/1165

import math

n = int(input())

for _ in range(n):
  x = int(input())
  raiz = int(math.sqrt(x))

  primo = True
  if (x % 2 != 0):
    for i in range(3, raiz + 1, 2):
      if(x % i == 0):
        primo = False
        break
  elif(x != 2):
    primo = False
  
  print(x, 'eh primo' if primo else 'nao eh primo')

