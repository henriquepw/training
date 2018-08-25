#include <stdio.h>

// www.urionlinejudge.com.br/judge/en/problems/view/1089
void musical_loop();
int is_sign_dif(int sign_current, int sign_next);

int main(){
	musical_loop();
  
	return 0;
}

void musical_loop(){
	while(1){
		int amostras = 0;
		scanf("%d", &amostras);
		
		if (amostras == 0) return;
		
		int list_amostras[amostras];
		for(int n = 0; n < amostras; n++)
			scanf("%d", &list_amostras[n]);
		
		int pulses = 0;
		int f_sign = is_sign_dif(list_amostras[0], list_amostras[1]);
		int sinal = f_sign;
		
		if (amostras > 2){
			for(int n = 1; n <= amostras; n++){
				int sinal_atual = sinal;
						
				if (n == amostras){
					if (list_amostras[n-1] != list_amostras[0])
						sinal_atual = is_sign_dif(list_amostras[n-1], list_amostras[0]);
						
				} else if (list_amostras[n] != list_amostras[n+1])
					sinal_atual = is_sign_dif(list_amostras[n], list_amostras[n+1]);
					
				
				pulses += (sinal != sinal_atual)? 1 : 0;	
				sinal = sinal_atual;
			}
			
			pulses += (sinal != f_sign)? 1 : 0;
			
		} else if ((amostras == 2) && (list_amostras[0] != list_amostras[1])){
			pulses = 2;
		}
			
		printf("%d \n", pulses);
	}  	
}

int is_sign_dif(int sign_current, int sign_next){
	return (sign_current > sign_next)? 1 : 0;
}
