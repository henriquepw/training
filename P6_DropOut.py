class Jogador:
    def __init__(self, name):
        self.nome = name
        self.carta = 0
        self.ativo = True


def cartas_diferentes():
    for x in range(0, num_jog - 1):
        if jogadores[x].ativo:
            if jogadores[x].carta != jogadores[x + 1].carta:
                return True
    return False


def print_vencedores():
    for x in range(num_jog):
        if jogadores[x].ativo:
            print(jogadores[x].nome, end=' ')
    print()


def qtd_ativos():
    ativos = 0
    for x in range(num_jog):
        if jogadores[x].ativo:
            ativos += 1
    return ativos


if __name__ == '__main__':
    num_jog = 2

    while True:
        num_jog = int(input())

        if num_jog == 0:
            break

        baralho = []
        jogadores = []
        for x in input().split(' '):
            jogadores.append(Jogador(x))

        for x in range(4):
            baralho += [int(x) for x in input().split(' ')]

        cont, j = 0, 0
        ativos = qtd_ativos()
        while cont < 52:  # (ativos * (52 // ativos)):
            menor = 14
            for i in range(num_jog):
                if jogadores[i].ativo:
                    jogadores[i].carta = baralho[cont]

                    if menor > jogadores[i].carta:
                        menor = jogadores[i].carta

                    cont += 1

                    if cont == 52:
                        break

            if cont == 52:
                break

            ativos = qtd_ativos()
            inativos_roda = 0
            if cartas_diferentes():
                for i in range(num_jog):
                    if jogadores[i].carta == menor:
                        jogadores[i].ativo = False
                        inativos_roda += 1

                    if inativos_roda == 4:
                        break
            else:
                break

            if qtd_ativos() == 1:
                break

        if num_jog != 0:
            print_vencedores()
