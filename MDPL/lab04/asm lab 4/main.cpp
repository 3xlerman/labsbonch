#include <conio.h>
#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>

#define arr_size 256

int64_t c, d;
extern "C"
{
	extern __int64 my_array[arr_size];
	
	uint64_t asmF(int64_t c, int64_t d);
}

void cMode()
{
	uint64_t counter = 0;
	for(uint64_t i = 0; i < arr_size; i++)
	{
		if (c <= my_array[i] && my_array[i] <= d)
		{
			counter++;
		}
	}
	printf("Result: %I64u\n", counter);
}
void asmMode()
{
	uint64_t sum = asmF(c, d);
	printf("Result: %I64u\n", sum);
}
int main()
{
	for (int64_t i = 0; i < arr_size; i++)
	{
		my_array[i] = (int64_t)rand() % 21 - 10;
		printf("%I64d\n", my_array[i]);
	}

	while (1)
	{
		printf("Enter c: ");
		scanf("%I64d", &c);
		printf("Enter d: ");
		scanf("%I64d", &d);
		
		cMode();
		asmMode();
	}
	_getch();
	return 0;
}