#include <stdio.h>

// https://www.urionlinejudge.com.br/judge/en/problems/view/2189
int main() {
  int participantes, teste = 1;

  while (1) {
    scanf("%d", &participantes);

    if (participantes == 0)
      break;

    int ingressos[participantes];
    for (int i = 0; i < participantes; i++)
      scanf("%d", &ingressos[i]);

    printf("Teste %d\n", teste);
    for (int i = 0; i < participantes; i++) {
      if (ingressos[i] == i + 1) {
        printf("%d\n\n", ingressos[i]);
        break;
      }
    }

    teste++;
  }

  return 0;
}
