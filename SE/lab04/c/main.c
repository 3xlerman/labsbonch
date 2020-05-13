#include <stdio.h>
#include "functions.h"

FILE *file;
double ax, ay, bx, by, cx, cy;

// Найти периметр треугольника по трём координатам (x,y)
int main() {
    input();
    perimetr(ax, ay, bx, by, cx, cy);
    return 0;
}
