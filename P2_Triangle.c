#include <stdio.h>

// www.urionlinejudge.com.br/judge/en/problems/view/1043
void triangle();

int main(){
	triangle();
  
	return 0;
}

void triangle(){
	float a, b, c;

  	scanf("%f", &a);
  	scanf("%f", &b);
  	scanf("%f", &c);
  
  	if((a < b + c) || (b < a + c) || (c < a + b)){
  		printf("Perimetro = %.1f", a + b + c);
  	} else {
  		printf("Area = %.1f", (a + b)* c / 2);
 	}
}

