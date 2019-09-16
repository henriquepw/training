#include <bits/stdc++.h>

using namespace std;
 
// https://www.urionlinejudge.com.br/judge/pt/runs/add/2454
int main() {
 
  bool p = 0, r = 0;

  cin >> p >> r;

  if (!p) cout << 'C' << endl;
  else {
    if (r) cout << 'A' << endl;
    else cout << 'B' << endl;
  }

  return 0;
}