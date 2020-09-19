#include <stdio.h>

// codeforces.com/problemset/problem/279/B
int main() {
  int n, t;
  scanf("%d %d", &n, &t);

  int livros[n];
  for (int i = 0; i < n; i++)
    scanf("%d", &livros[i]);

  int soma[n + 1];
  soma[0] = 0;
  for (int i = 1; i < n + 1; i++)
    soma[i] = soma[i - 1] + livros[i - 1];

  int result = 0, result_aux = 0;
  if (t >= soma[n]) {
    result = n;
  } else {
    for (int i = 1; i < n; i++) {
      int left = i, right = n, mid;

      while (left <= right) {
        mid = (left + right) / 2;

        if (soma[mid] - soma[i - 1] <= t) {
          if (soma[mid + 1] - soma[i - 1] > t) {
            result_aux = mid - i + 1;
            break;
          } else
            left = mid + 1;
        } else
          right = mid - 1;
      }
      if (result_aux > result)
        result = result_aux;
    }
  }

  printf("\n%d", result);

  return 0;
}
