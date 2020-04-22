#include <iostream>
#include <math.h>
using namespace std;
int main()
{
    float a,b,x,p,y,z;
    cout << "x= ";
    cin >> x;
    cout << "a= ";
    cin >> a;
    cout << "b= ";
    cin >> b;
    y = (sqrt(x*x + 16)) / (x + 2);
    p = sqrt(sin(a)+3);
    z = (y + p + b) / (y*y + p);

    cout.precision(3);
    cout << "y= " << y << " " << "z=" << z;
    return 0;
}