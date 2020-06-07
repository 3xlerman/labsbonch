#include <stdio.h>
#include <conio.h>

extern "C"
{
	signed   __int8  s8_a;
	unsigned __int8  u8_a;
	signed   __int16 s16_a;
	unsigned __int16 u16_a;

	signed   __int8  s8_c;
	unsigned __int8  u8_c;
	signed   __int16 s16_c;
	unsigned __int16 u16_c;

	signed   __int8  s8_d;
	unsigned __int8  u8_d;
	signed   __int16 s16_d;
	unsigned __int16 u16_d;
}

extern "C" signed __int16 calc_s8();
extern "C" signed __int16 calc_u8();
extern "C" signed __int16 calc_s16();
extern "C" signed __int16 calc_u16();

//1 (3*c+8-d)/(a-c/4)
signed __int16 test_calc_s8()
{
	return (3 * s8_c + 8 - s8_d) / (s8_a - s8_c / 4);
}
signed __int16 test_calc_u8()
{
	return (3 * u8_c + 8 - u8_d) / (u8_a - u8_c / 4);
}
signed __int16 test_calc_s16()
{
	return (3 * s16_c + 8 - s16_d) / (s16_a - s16_c / 4);
}
signed __int16 test_calc_u16()
{
	return (3 * u16_c + 8 - u16_d) / (u16_a - u16_c / 4);
}

void m1()
{
	int buff;
	printf("Enter a: ");
	scanf("%i", &buff);
	s8_a = buff;
	printf("Enter c: ");
	scanf("%i", &buff);
	s8_c = buff;
	printf("Enter d: ");
	scanf("%i", &buff);
	s8_d = buff;

	int r = calc_s8();

	printf("Result: %i\n", r);
	r = test_calc_s8();
	printf("Test:   %i\n", r);
}
void m2()
{
	int buff;
	printf("Enter a: ");
	scanf("%i", &buff);
	u8_a = buff;
	printf("Enter c: ");
	scanf("%i", &buff);
	u8_c = buff;
	printf("Enter d: ");
	scanf("%i", &buff);
	u8_d = buff;

	int r = calc_u8();

	printf("Result: %i\n", r);
	r = test_calc_u8();
	printf("Test:   %i\n", r);
}
void m3()
{
	int buff;
	printf("Enter a: ");
	scanf("%i", &buff);
	s16_a = buff;
	printf("Enter c: ");
	scanf("%i", &buff);
	s16_c = buff;
	printf("Enter d: ");
	scanf("%i", &buff);
	s16_d = buff;

	int r = calc_s16();

	printf("Result: %i\n", r);
	r = test_calc_s16();
	printf("Test:   %i\n", r);
}
void m4()
{
	int buff;
	printf("Enter a: ");
	scanf("%i", &buff);
	u16_a = buff;
	printf("Enter c: ");
	scanf("%i", &buff);
	u16_c = buff;
	printf("Enter d: ");
	scanf("%i", &buff);
	u16_d = buff;

	int r = calc_u16();

	printf("Result: %i\n", r);
	r = test_calc_u16();
	printf("Test:   %i\n", r);
}

int main()
{
	printf("Mode 1: signed 8\n");
	printf("Mode 2: unsigned 8\n");
	printf("Mode 3: signed 16\n");
	printf("Mode 4: unsigned 16\n");

	int mode;
	while (1)
	{
		printf("Enter mode: ");
		scanf("%i", &mode);
		printf("Select mode %i\n", mode);

		switch (mode)
		{
		case (1):
		{
			m1();
			break;
		}
		case (2):
		{
			m2();
			break;
		}
		case (3):
		{
			m3();
			break;
		}
		case (4):
		{
			m4();
			break;
		}

		default:
			break;
		}
	}
	_getch();
	return 0;
}