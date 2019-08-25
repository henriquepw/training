#include <bits/stdc++.h>

using namespace std;

int main() {
  int n = 0, x = 0;

  cin >> n;
  while (n--) {
    cin >> x;

    bool primo = true;
    for (int i = 2; i < x; i++) {
      if(x % i == 0) {
        primo = false;
        break;
      }
    }

    cout << x << (primo ? " eh primo" : " nao eh primo") << endl;
  }
   
  return 0;
}