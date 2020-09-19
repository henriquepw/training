package main

import "fmt"

func mergeStrings(str1, str2 string, len int, result *string) {
  for j := 0; j < len; j++ {
    *result += string(str1[j]) + string(str2[j])
  }
}

// Q1238 https://www.urionlinejudge.com.br/judge/pt/problems/view/1238
func Q1238() {
  var N int
  fmt.Scanf("%d", &N)

    for i := 0; i < N; i++ {
      var str1, str2, result string

      fmt.Scanf("%s %s", &str1, &str2)

      str1Len := len(str1)
      str2Len := len(str2)
  
      if(str1Len < str2Len) {
        mergeStrings(str1, str2, str1Len, &result);
        result += str2[str1Len:str2Len]
      } else {
        mergeStrings(str1, str2, str2Len, &result);
        result += str1[str2Len:str1Len]
      }

    fmt.Printf("%s\n", result)
  }
}