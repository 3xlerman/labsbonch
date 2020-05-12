#include <stdio.h>
#include <math.h>


// Найти периметр треугольника по трём координатам (x,y)
int main() {
    FILE *file;
    double x1, y1, x2, y2, x3, y3;
    file = fopen("points.txt", "r");
    int i = 0;
    while (fscanf(file, "%lf %lf %lf %lf %lf %lf", &x1, &y1, &x2, &y2, &x3, &y3) != EOF) {
        printf("%lf %lf %lf %lf %lf %lf", x1, y1, x2, y2, x3, y3);
        i++;
    }
    double a = sqrt(pow(x1 - x2, 2) + pow(y1 - y2, 2));
    double b = sqrt(pow(x2 - x3, 2) + pow(y2 - y3, 2));
    double c = sqrt(pow(x3 - x1, 2) + pow(y3 - y1, 2));
    printf("\nP=%lf", a + b + c);
}
