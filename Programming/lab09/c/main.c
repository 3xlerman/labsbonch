// RUS С клавиатуры вводится текст, состоящий из строк (не более 20).
// Длина каждой строки - не более 128. Слова разделены одним или
// более пробелами, в каждой строке не менее 2 слов.
// Задача: удалить из каждой строки слова с чётными номерами.

// ENG Text, consisting of strings (less than 20), is being printed from keyboard.
// Every stroke has a length less than 128. Words divided by one or more
// spaces, number of words - more than 2.
// Objective is delete words with even numbers from every stroke.

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define row 20
#define pos 128

int n, i, spaceC, wordsC, l, space1, space2;
char str[row][pos], words[65][65];

void input() {
    printf("Enter number of strokes: ");
    scanf("%d", &n);
    if (n > 20) {
        printf("Out of bounds (20 strokes)!");
        exit(1);
    } else {
        fgets(str[0], pos, stdin);
        for (i = 0; i < n; i++) {
            fgets(str[i], pos, stdin);
        }
        printf("\n");
    }
}

void divide() {
    int j = 0;
    wordsC = 0;
    while (j < strlen(str[i])) {
        if (str[i][j] != ' ') {
            words[wordsC][j] = str[i][j];
        } else {
            wordsC++;
            j++;
        }
    }

}

int main() {
    spaceC = 0;
    input();
    for (i = 0; i < row; i++) {
        divide();
    }
    //output();
    return 0;
}
