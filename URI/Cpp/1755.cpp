#include <bits/stdc++.h>

using namespace std;

int main() {
  int t;
  cin >> t;

  while (t--) {
    int d, n;
    cin >> d >> n;

    float troco, marca, max = 0;
    for (int i = 0; i < n; i++) {
      cin >> marca;

      int x = 1;
      do {
        x++;
        troco = marca * x;
      } while (troco <= d);

      troco = d - (marca * (x - 1));
      if (troco > max)
        max = troco;
    }

    cout << fixed << setprecision(2) << max << endl;
  }

  return 0;
}
