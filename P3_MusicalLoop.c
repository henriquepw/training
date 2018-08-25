#include <stdio.h>

// www.urionlinejudge.com.br/judge/en/problems/view/1089
int get_sign(int sign_current, int sign_next);

int main(){
	while(1){
		int amostras = 0;
		scanf("%d", &amostras);
		
		if (amostras == 0) break;
		
		int list_amostras[amostras];
		for(int n = 0; n < amostras; n++)
			scanf("%d", &list_amostras[n]);
		
		int pulses = 0;
		int f_sign = get_sign(list_amostras[0], list_amostras[1]);
		int sinal_atual = f_sign;
		int sinal = f_sign;
		
		if (amostras > 2){
			for(int n = 1; n < amostras; n++){
				printf("%d\n", sinal_atual);
				sinal = sinal_atual;
				
				if (n == amostras-1) {
					if (list_amostras[n] != list_amostras[0])
						sinal_atual = get_sign(list_amostras[n], list_amostras[0]);
				} else if (list_amostras[n] != list_amostras[n+1]){
					sinal_atual = get_sign(list_amostras[n], list_amostras[n+1]);
				}
							
				pulses += (sinal != sinal_atual)? 1 : 0;
			}
			
			printf("fsing: %d\n", f_sign);
			printf("sing final: %d\n", sinal);
			pulses += (f_sign != sinal_atual)? 1 : 0;
			
		} else if ((amostras == 2) && (list_amostras[0] != list_amostras[1])){
			pulses = 2;
		}
			
		printf("%d\n", pulses);
	}  	
  
	return 0;
}

int get_sign(int sign_current, int sign_next){
	return (sign_current > sign_next)? 1 : 0;
}
