# https://www.urionlinejudge.com.br/judge/pt/problems/view/2095

n = int(input())

enemies = [int(i) for i in input().split(' ')]
allies = [int(i) for i in input().split(' ')]

enemies.sort()
allies.sort()

wins = 0
for index in range(n):
	if allies[index] > enemies[wins]:
		wins += 1

print(wins)
