#include <stdio.h>
#include <math.h>

double a, b, c, d;

void input() { // функция ввода коэффициентов квадратного уравнения
    printf("\nEnter a: ");
    scanf("%lf", &a);
    printf("\nEnter b: ");
    scanf("%lf", &b);
    printf("\nEnter c: ");
    scanf("%lf", &c);
}

void solve() { // функция ветвления
    if (a == 0) { // линейное уравнение
        if (b == 0) {
            if (c == 0) {
                printf("Answer: X c R"); // Х - любое вещественное число
            } else {
                printf("Answer: X c O/"); // Х принадлежит пустому множеству
            }
        } else {
            double x = -c / b;
            printf("Answer: ");
            printf("%lf", x);
        }
    } else if (b == 0) { // квадратное уравнение
        if (c == 0) {
            int x = 0;
            printf("Answer: ");
            printf("%d", x);
        } else {
            double x = sqrt(-c / a);
            printf("Answers: ");
            printf("%lf%lf", x, -x);
        }
    } else { // вычисляем дискриминант
        d = b * b - 4 * a * c;
        if (d < 0) {
            printf("Answer: X c O/");
        } else {
            double x1 = (-b + sqrt(d)) / (2 * a);
            double x2 = (-b - sqrt(d)) / (2 * a);
            printf("Answers: ");
            printf("%lf%lf", x1, x2);
        }
    }
}

int main() { // главная функция, в которой вызываются остальные
    input();
    solve();
}
