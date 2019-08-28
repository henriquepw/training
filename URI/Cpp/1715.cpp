#include <bits/stdc++.h>

using namespace std;

int main() {
  int n = 0, m = 0, result = 0;

  cin >> n >> m;

  while(n--) {
    int gol;
    bool is_gol = true;

    for(int i = 0; i < m; i++){
      cin >> gol;
      if (gol == 0) is_gol = false;
    }

    if (is_gol) result++;   
  }

  cout << result << endl;

  return 0;
}