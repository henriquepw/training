#include <stdio.h>

// https://www.urionlinejudge.com.br/judge/en/problems/view/1029
int fib(int num);

int calls = 0;
int main() {
  int tests;
  scanf("%d", &tests);

  int nums[tests];
  for (int i = 0; i < tests; i++)
    scanf("%d", &nums[i]);

  for (int i = 0; i < tests; i++) {
    calls = -1;
    int f = fib(nums[i]);

    printf("fib(%d) = %d calls = %d\n", nums[i], calls, f);
  }

  return 0;
}

int fib(int num) {
  calls++;
  if (num == 0)
    return 0;
  else if (num == 1)
    return 1;
  else
    return fib(num - 1) + fib(num - 2);
}
