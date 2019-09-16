#include <bits/stdc++.h>

using namespace std;
 
// https://www.urionlinejudge.com.br/judge/pt/runs/add/1021
int main() {
  int coins[] = {10000, 5000, 2000, 1000, 500, 200, 100, 50, 25, 10, 5, 1};
 
  double n = 0;
  cin >> n;
  n *= 100;

  cout << "NOTAS:" << endl;
  for (int x : coins) {
    if (x == 100) cout << "MOEDAS:" << endl;

    int total = n / x;
    n = n - (total * x);

    if (x > 100) 
      printf("%d nota(s) de R$ %d.00\n", total,  x / 100);
    else 
      printf("%d moedas(s) de R$ %.2f\n", total,  (double) x / 100);
  }

  return 0;
}