#include <stdio.h>

// https://www.urionlinejudge.com.br/judge/en/problems/view/1300
int main() {
  int angle;

  while (scanf("%d", &angle) != EOF)
    printf((angle % 6 == 0) ? "Y\n" : "N\n");

  return 0;
}
