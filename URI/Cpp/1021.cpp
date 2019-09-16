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

    double total = (x == 1) ? n : (int) (n / x);
    n -= (total * x);
      
    if (x > 100)
      printf("%.0f nota(s) de R$ %d.00\n", total,  x / 100);
    else
      printf("%.0f moeda(s) de R$ %.2f\n", total,  (float) x / 100);
  }

  return 0;
}