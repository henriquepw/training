#include <bits/stdc++.h>

using namespace std;

// https://www.urionlinejudge.com.br/judge/pt/problems/view/1566
int main() {
  int N, NC, value;

  scanf("%d", &N);

  for (int i = 0; i < N; i++) {
    scanf("%d", &NC);

    int heights[231] = {};
    for (int j = 0; j < NC; j++) {
      scanf("%d", &value);
      heights[value]++;
    }

    bool notFirst = false;
    for (int j = 20; j < 231; j++) {
      if (heights[j] > 0) {
        for (int h = 0; h < heights[j]; h++) {
          if (notFirst) printf(" ");
          else notFirst = true;
          
          printf("%d", j);
        }
      }
    }
    
    printf("\n");
  }

  return 0;
}