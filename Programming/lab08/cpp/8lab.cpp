#include <iostream>
#include <cstdlib>
#include <iomanip>
using namespace std;
#define N 5
#define M 5
int i,j,m,n,s;
float a[M][N];
int fun()
{

    for (int i = 0; i < N ; i++)
    {
        for (int j = 0; j < M ; j++)
        {
            a[i][j]= rand() % 101;
            cout <<  "a["<< i << "][" << j << "]= " << setw(2) << a[i][j] << " ";

        }
        cout << endl;
    }
    return 0;
}

int main ()
{

    s= 0;
    fun();
    cout << endl;
    for (int i = 0; i < N; i++)
    {
        for (int j = 0; j < M; j++)
            s = s + a[i][j];
        cout << "s["<< i <<"]= " << s << endl;
        s = 0;
    }

    return 0;
}