#include <iostream>
#include <iomanip>
#include <cstdlib>
using namespace std;
#define N 10
int main()
{
    int i, amax,bmax,cmax;
    amax = bmax = cmax = -1;


    float a[N], b[N], c[N], x[N], y[N];

    for (i = 0; i < N; i++)
    {
        a[i]= rand() % 101;
        cout << "a[" << i << "]= "<<a[i]<<"   ";
    }
    cout << endl;

    for (i = 0; i < N; i++)
    {
        b[i]= rand() % 101;
        cout << "b[" << i << "]= "<<b[i]<<"   ";
    }
    cout << endl;

    for (i = 0; i < N; i++)
    {
        c[i]= rand() % 101;
        cout << "c[" << i << "]= "<<c[i]<<"   ";
    }
    cout << endl << endl;


    for (int i = 0; i < N; i++)
    {
        if(a[i]> amax)
            amax = a[i];
        if(b[i]> bmax)
            bmax = b[i];
        if(c[i]> cmax)
            cmax = c[i];
    }

    for (int i = 0; i < N; i++)
    {
        x[i] = ((amax + bmax) - (a[i] + b[i])) / 2;
        y[i] = ((bmax + cmax) - (b[i] + c[i])) / 2;
        cout <<
             "x["<<i<<"]=" <<setw(5)<< x[i]
             <<"   "<<"y["<<i<<"]=" <<setw(5)<< y[i] << endl;
    }
    return 0;
}