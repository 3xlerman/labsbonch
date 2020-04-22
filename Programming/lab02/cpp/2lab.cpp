#include <iostream>
#include <math.h>
using namespace std;

int main()
{
    float x, y;
    cout << "x=";
    cin >> x;
    if (x < -3)
        y = cos(x) / (x+10);
    else if (x < 4)
        y = exp(x / 10);
    else if (x < 6)
        y = log10(x);
    else y = pow(sin(x),2);
    cout.precision(3);
    cout << "y= " << y;
    return 0;
}

