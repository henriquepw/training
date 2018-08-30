#include<stdio.h>

// codeforces.com/problemset/problem/18/C

int main(){
	int n;
	scanf("%d", &n);
	
	int tape[n];
	for(int i = 0; i < n; i++)
		scanf("%d", &tape[i]);
	
	int sum[n];
	sum[0] = tape[0];
	for(int i = 1; i < n; i++)
		sum[i] = sum[i - 1] + tape[i]; 
	
	int cut = 0;
	for(int i = 0; i < n-1; i++)
		if(sum[i] == sum[n - 1] - sum[i]) cut++;

	
	printf("%d", cut);
		
	return 0;
}
