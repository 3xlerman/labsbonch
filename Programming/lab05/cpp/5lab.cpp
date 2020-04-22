#include <iostream>
#include <math.h>
using namespace std;

int main()
{
    float x, y, znamSum, chislSum, i, j, k, n, m;
    cout << "n= ";
    cin >> n;
    cout << "m= ";
    cin >> m;
    cout << "x= ";
    cin >> x;
    y = 0;
    znamSum = 0;
    chislSum = 0;
    for(i=1;i<=n;i++)
    {
        for(j=1;j<=m;j++)
        {
            chislSum = chislSum + pow(i+j,2);
        }
        for(k=1;k<=m;k++)
        {
            znamSum = znamSum + k + 1;
        }
        y = y + ((2*x + chislSum) / (x + i * znamSum));
    }
    cout << "y= " << y;
    return 0;

}