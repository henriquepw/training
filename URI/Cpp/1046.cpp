#include <bits/stdc++.h>

using namespace std;
 
// https://www.urionlinejudge.com.br/judge/pt/runs/add/1046
int main() {
 
  int init = 0, final = 0, result = 0;

  if (init == final) result = 24;
  else result = (final > init) ? final - init : (24 - init) + final;
  
  cout << "O JOGO DUROU " << result << " HORA(S)" << endl;

  return 0;
}