#include <iostream>
#include <math.h>
using namespace std;
float i, s, a, m, n, x;


int fun(int n1,int n2,int a2,int a1,int a0)
{
    s = 0;
    for (int i = n1; i <= n2; i++)
    {
        s = s + pow((a2*i*i + a1*i + a0),2);
        return s;
    }
}

int main()
{
    cout << "n= ";
    cin >> n;
    cout << "m= ";
    cin >> m;
    cout << "a= ";
    cin >> a;
    cout << "x= ";
    cin >> x;
    float y = (a + fun(1,m,2,1,2)) / (4 + fun(2,n,1,0,3));
    cout << "y= " << y;
    return 0;

}