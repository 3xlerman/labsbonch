#include <stdio.h>

typedef double(*function)(double x);

void tab(function f) {
    for (double i = -1000; i < 1000; i = i + 1) {
        printf("%lf \n", f(i));
    }
}

double f(double x) {
    return x * 2;
}

double f2(double x) {
    return x * x;
}

int main() {
    tab(&f);
    tab(&f2);
}

