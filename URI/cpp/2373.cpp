#include <bits/stdc++.h>

using namespace std;

// https://www.urionlinejudge.com.br/judge/pt/problems/view/2373
int main() {
  int n = 0, l = 0, c = 0, total = 0;

  cin >> n;
  while (n--) {
    cin >> l >> c;

    if (l > c) total += c;
  }

  cout << total << endl;
   
  return 0;
}