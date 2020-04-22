#include <iostream>
#include <iomanip>
#include <math.h>
using namespace std;
int main()
{
    float x,y,xn,xk,dx;
    int i,n, a;
    cout<< "xn= "; cin >> xn;
    cout<< "xk= "; cin >> xk;
    cout << "n= "; cin >> n;
    cout << "a= "; cin >> a;

    dx = (xk-xn) / (n-1);
    x = xn;
    cout << setw(5)<<"i"<<setw(10)<<setprecision(3)<<"x"<<setw(10)<<setprecision(3)<<"y"<<endl;
    for(i=0;i<n;i++,x+=dx)
    {
        y = (exp(-x) + exp(sqrt(a))) / exp(x-a);
        cout<<setw(5)<<i<<setw(10)<<setprecision(3)<<x<<setw(10)<<setprecision(3)<<y<<endl;
    }
    return 0;
}

