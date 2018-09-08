#include <bits/stdc++.h>

using namespace std;

int main() {
  long long l, r, x, y, k;
  cin >> l >> r >> x >> y >> k;

  short int find = 0;
  for (float a = k; a <= r; a += k) {
    if (a < l)
      continue;
    else {
      if (a / k >= x and a / k <= y) {
        find = 1;
        break;
      }
    }
  }

  string result = find ? "YES" : "NO";
  cout << result << "\n";

  return 0;
}
