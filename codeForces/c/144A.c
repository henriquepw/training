#include <stdio.h>

// codeforces.com/problemset/problem/144/A
int main() {
  int n;
  scanf("%d", &n);

  int soldier[n];
  for (int i = 0; i < n; i++)
    scanf("%d", &soldier[i]);

  int pos_max = 0, pos_min = n - 1;
  for (int i = 0; i < n; i++)
    if (soldier[i] > soldier[pos_max])
      pos_max = i;
    else if (soldier[i] <= soldier[pos_min])
      pos_min = i;

  int k = (pos_min > pos_max) ? 0 : 1;
  printf("%d", (pos_max + n - 1 - pos_min - k));

  return 0;
}
