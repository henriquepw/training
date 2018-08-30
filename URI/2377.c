#include <stdio.h>
 
// www.urionlinejudge.com.br/judge/pt/runs/add/2377
int main() {
 	int comprimento = 0, distancia = 0;
 	int custo_km = 0, valor_pedagio = 0;
    scanf("%d %d", &comprimento , &distancia);
	scanf("%d %d", &custo_km , &valor_pedagio);
	
	printf("%d\n", ((comprimento/distancia) * valor_pedagio) + (custo_km * comprimento));
	
	return 0;
}
