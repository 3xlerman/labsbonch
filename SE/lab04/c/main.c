#include <stdio.h>
#include <math.h>

FILE *file;
double ax, ay, bx, by, cx, cy;

// Функция ввода
void input() {

    file = fopen("points.txt", "r");
    int i = 0;
    while (fscanf(file, "%lf %lf %lf %lf %lf %lf", &ax, &ay, &bx, &by, &cx, &cy) != EOF) {
        printf("%lf %lf %lf %lf %lf %lf", ax, ay, bx, by, cx, cy);
        i++;
    }
}
// Найдём периметр, используя теорему Пифагора
void perimetr(double x1, double y1, double x2, double y2, double x3, double y3) {
    double a = sqrt(pow(x1 - x2, 2) + pow(y1 - y2, 2));
    double b = sqrt(pow(x2 - x3, 2) + pow(y2 - y3, 2));
    double c = sqrt(pow(x3 - x1, 2) + pow(y3 - y1, 2));
    printf("\nP=%lf", a + b + c);
}


// Найти периметр треугольника по трём координатам (x,y)
int main() {
    input();
    perimetr(ax, ay, bx, by, cx, cy);
    return 0;
}
