#include <stdio.h>

typedef double(*function)(double x);

void tabulation(function f) {
    for (int i = -10; i < 10; i++) {
        printf("f(%d)= %lf \n", i, f(i));
    }
}

double f(double x) {
    return x * 2 + 1;
}

double f2(double x) {
    return x * x * x;
}

int main() {
    printf("The 1st function: \n");
    tabulation(&f);
    printf("The 2nd function: \n");
    tabulation(&f2);
}

