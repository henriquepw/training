#include <stdio.h>

// www.urionlinejudge.com.br/judge/en/problems/view/1043

int main(){
	float a, b, c;

  	scanf("%f %f %f", &a, &b, &c);
  
  	if((a < b + c) && (b < a + c) && (c < b + a)){
  		printf("Perimetro = %.1f\n", a + b + c);
  	} else {
  		printf("Area = %.1f\n", (a + b)* c / 2);
 	}
  
	return 0;
}

