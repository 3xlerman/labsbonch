// RUS Массив содержит сведения о книгах.
// Каждая структура имеет следующие поля:
// автор (авторы), название, год издания, цена и издательство.
// Вывести на экран дисплея список книг,
// изданных в заданном временном интервале.

// ENG An array contains information about books.
// Every structure has these fields:
// author (authors), title, year of publishing, price and publishing office.
// Print a book list, published in entered time interval.

#include <stdio.h>

#define n 5

struct books {
    char Author[40];
    char Title[30];
    int Year;
    float Price;
    char Publish[20];
};

struct books book[n];

void input() {
    for (int i = 0; i < n; i++) {
        printf("book[%d%s", i, "]\n");
        printf("Author: ");
        fgets(book[i].Author, 40, stdin);
        printf("Title: ");
        fgets(book[i].Title, 30, stdin);
        printf("Year: ");
        scanf("%d", &book[i].Year);
        printf("Price: ");
        scanf("%f", &book[i].Price);
        printf("Publish: ");
        fgets(book[i].Publish, 20, stdin);
        printf("\n");
    }
}

void search() {
    int b1, b2;
    printf("Enter first year bound: ");
    scanf("%d", &b1);
    printf("Enter second year bound: ");
    scanf("%d", &b2);
    for (int i = 0; i < n; i++) {
        if (book[i].Year >= b1 && book[i].Year <= b2)       // searching books in time interval
            printf("%s", book[i].Title, "\n");
    }
}

int main() {
    input();
    search();
    return 0;
}