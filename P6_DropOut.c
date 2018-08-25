#include <stdio.h>

// www.urionlinejudge.com.br/judge/en/problems/view/1327

struct jogador{
	char nome[16];
	int carta;
};

int main(){
 	int players = 2;
 	int baralho[4][13];
 	
 	while(players != 0){
 		scanf("%d", &players);
 	
	 	jogador ativos[players];
	 	jogador inativos[players];
	 	
	 	for(int i = 0; i < players; i++)
	 		scanf("%s", ativos[i].nome);	
	 		
	 	for(int i = 0; i < 4; i++)
	 		for(int j = 0; j < 13; j++)
	 			scanf("%d", &baralho[i][j]);
	 	
		while(){
			
			int cont = 0;
		 	if( 52-cont > 4){
		 		for(int i = n; i < 4; i++)
			 		for(int j = m; j < 13; j++){
						ativos.carta = baralho[i][j];								
					}
					
			} else {
				for(int i = 0; i < sizeof(ativos); i++)
		 			printf("%s ", ativos[i].nome);
			}		
			
		}
		
	}
 		
	return 0;
}
