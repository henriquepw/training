#include <bits/stdc++.h>

using namespace std;
 
// https://www.urionlinejudge.com.br/judge/pt/problems/view/1049
int main() {
  string animal[2][2][2] = {
    {{"aguia", "pomba"}, {"homem", "vaca"}},
    {{"pulga", "lagarta"}, {"sanguessuga", "minhoca"}}
  };

  string word1, word2, word3;

  std::cin >> word1 >> word2 >> word3;

  int i = word1 == "vertebrado" ? 0 : 1;
  int j = word2 == "ave" || word2 == "inseto" ? 0 : 1;

  int k = word3 == "hematofago" || word3 == "carnivoro" 
  || (word3 == "onivoro" && word2 == "mamifero")? 0 : 1;

  cout << animal[i][j][k] << endl;

  return 0;
}