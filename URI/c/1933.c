#include <stdio.h>

// https://www.urionlinejudge.com.br/judge/en/problems/view/1933
int main() {
  int card1, card2;
  scanf("%d %d", &card1, &card2);

  if (card1 == card2)
    printf("%d\n", card1);
  else
    printf("%d\n", (card1 > card2) ? card1 : card2);

  return 0;
}
