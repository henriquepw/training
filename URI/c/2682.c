#include <stdio.h>

// https://www.urionlinejudge.com.br/judge/pt/problems/view/2682
int main() {
  int n = 0, stay = 1, anterior = 0, prox = 1;

  while (scanf("%d", &n) != EOF) {
    if (stay)
      if (n < anterior)
        prox = n + 1;
      else
        stay = 0;

    anterior = n;
  }

  printf("%d\n", prox);
  return 0;
}
