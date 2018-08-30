#include <stdio.h>

// codeforces.com/problemset/problem/4/A

int main(){
 	int w;
	
	scanf("%d", &w);
	
	printf((w % 2 == 0) && (w >= 4)? "Yes" : "No");
	
	return 0;
}
