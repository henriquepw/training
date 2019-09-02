#include <bits/stdc++.h>

using namespace std;

// https://www.urionlinejudge.com.br/judge/pt/problems/view/2423
int main() {
  int a, b, c;

  cin >> a >> b >> c;
  a /= 2, b /= 3, c /= 5;

  int min = a < b ? a : b;
  min = min < c ? min : c;

  cout << min << endl;

  return 0;
}
