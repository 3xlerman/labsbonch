#include <stdio.h>
#include <io.h>
#include <fcntl.h>

#include <conio.h>


extern "C"
{
	signed __int16 a_sig;
	signed __int16 b_sig;

	unsigned __int16 a_unsig;
	unsigned __int16 b_unsig;

	signed __int16 x_return;

	void signedF();
	void unsignedF();
}

void signedFT()
{
	int a = a_sig;
	int b = b_sig;

	if (a < b)
	{
		wprintf(L"Тест x: %i\n", (a * b) / 4);
	}
	else if (a == b)
	{
		wprintf(L"Тест x: %i\n", 21);
	}
	else if (a > b)
	{
		wprintf(L"Тест x: %i\n", (a + 2) / b);
	}
}
void unsignedFT()
{
	int a = a_unsig;
	int b = b_unsig;

	if (a < b)
	{
		wprintf(L"Тест x: %i\n", (a * b) / 4);
	}
	else if (a == b)
	{
		wprintf(L"Тест x: %i\n", 21);
	}
	else if (a > b)
	{
		wprintf(L"Тест x: %i\n", (a + 2) / b);
	}
}

int sig()
{
	int data;
	wprintf(L"a: ");
	wscanf(L"%i", &data);
	a_sig = data;

	wprintf(L"b: ");
	wscanf(L"%i", &data);
	b_sig = data;

	signedFT();
	signedF();

	return x_return;
}
int unsig()
{
	int data;
	wprintf(L"a: ");
	wscanf(L"%i", &data);
	a_unsig = data;

	wprintf(L"b: ");
	wscanf(L"%i", &data);
	b_unsig = data;

	unsignedFT();
	unsignedF();

	return x_return;
}

int main()
{
	//рус. яз.
	_setmode(_fileno(stdout), _O_U16TEXT);
	_setmode(_fileno(stdin), _O_U16TEXT);
	_setmode(_fileno(stderr), _O_U16TEXT);

	while (1)
	{
		int mode;
		wprintf(L"1 Знаковый\n");
		wprintf(L"2 Беззнаковый\n");
		wprintf(L"Enter mode: ");
		wscanf(L"%i", &mode);

		switch (mode)
		{
		case(1):
		{
			wprintf(L"x: %i\n", sig());
			break;
		}
		case(2):
		{
			wprintf(L"x: %i\n", unsig());
			break;
		}

		default:
			break;
		}
	}

	_getch();
	return 0;
}