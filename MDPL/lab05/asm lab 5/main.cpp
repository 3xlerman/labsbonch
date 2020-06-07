#include <stdio.h>
#include <conio.h>


//(3*c+8-d)/(a-c/4)
extern "C"
{
	float a_f, result_f;
	int c_i, d_i;

	void ASMCalc();
}

void CPP()
{
	float t1 = (3 * c_i + 8.0 - d_i);
	float t2 = (a_f - c_i / 4.0);

	printf("C++: %f\n", t1 / t2);
}
void asmCall()
{
	ASMCalc();
	printf("ASM: %f\n", result_f);
}

int main()
{
	while (1)
	{
		printf("Enter a: ");
		scanf("%f", &a_f);
		printf("Enter c: ");
		scanf("%i", &c_i);
		printf("Enter d: ");
		scanf("%i", &d_i);

		CPP();
		asmCall();
	}
	return 0;
}