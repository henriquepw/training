#include <stdio.h>
 
// www.urionlinejudge.com.br/judge/pt/problems/view/2297
int main() {
 
    int rodadas = 0;
    int aldo = 0, beto = 0;
    
    int teste = 1;
    while(1){
    	scanf("%d", &rodadas);
    	
    	if(rodadas == 0) break;
    	
    	for(int i = 0; i < rodadas; i++){
	    	int a = 0, b = 0;
	    	scanf("%d", &a);
	    	scanf("%d", &b);
	    	
    		aldo += a;
    		beto += b;
    		
		}
		
		printf("Teste %d\n", teste);
		printf((aldo > beto)? "Aldo\n" : "Beto\n");
		printf("\n");
		teste++;
 	
	}
 
    return 0;
}
