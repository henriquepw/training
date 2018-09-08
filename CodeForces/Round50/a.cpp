#include <bits/stdc++.h>

using namespace std;

int main() {
  unsigned long long n, k;
  cin >> n >> k;

  unsigned long long result;
  
  if (n > k)
    result = 1;
  else if (k % n == 0)
    result = k / n;
  else
    result = k / n + 1;

  cout << result << '\n';

  return 0;
}
