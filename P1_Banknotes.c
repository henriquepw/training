#include <stdio.h>
 
#define mod(x) ((x >= 0)? x : -1 * x)

//www.urionlinejudge.com.br/judge/en/problems/view/1018
int main() {
    int value = 0;
    int notes[] = {100, 50, 20, 10, 5, 2, 1};

	scanf("%d", &value);
	printf("%d\n", value);

	for (int i = 0; i <  7; i++){
		int aux = value / notes[i];
		printf("%d nota(s) de R$ %d,00\n", aux, notes[i]);
		value = value % notes[i];
	}
 
    return 0;
}

