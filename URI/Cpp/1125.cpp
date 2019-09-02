#include <bits/stdc++.h>

using namespace std;

// incomplete
// https://www.urionlinejudge.com.br/judge/pt/problems/view/1125
int main() {
  int G = 0, P = 0; //G = Corridas / P = Pilotos
  cin >> G, P;

  while (G--) {
    int rank[P];

    for(int j = 0; j < P; j++) cin >> rank[j];

    int S = 0;
    cin >> S;

    int points[S][P];
    for(int j = 0; j < S; j++)
      for(int k = 0; k < S; k++)
        cin >> points[j][k];


  }

  return 0;
}