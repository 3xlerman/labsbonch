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
#include <ctype.h>

#define row 20
#define pos 128

int n, i, wordC, space1;
char str[row][pos], words[row][65][65];

void input() {
    printf("Number of strokes: ");
    scanf("%d", &n);
    if (n > 20) {
        printf("Out of bounds (20)!");
        exit(1);
    }
    fgets(str[0], pos, stdin);
    printf("Enter strings: \n");
    for (i = 0; i < n; i++) {
        fgets(str[i], pos, stdin);
    }
}

void wordsDivider() {
    for (i = 0; i < n; i++) {
        wordC = 0;
        space1 = 0;
        for (int j = 0; j < strlen(str[i]); j++) {
            if (isspace(str[i][j]) || str[i][j] == '\0') {
                if (j - space1 > 1) {
                    int c1 = 0;
                    for (int c = space1; c < j; c++) {
                        words[i][wordC][c1] = str[i][c];
                        c1++;
                    }
                    space1 = j + 1;
                    wordC++;
                } else {
                    space1 = j + 1;
                }

            }
        }

    }
}

void finalOutput() {
    for (i = 0; i < n; i++) {
        for (int j = 0; j < wordC; j++) {
            if (j % 2 == 0) {
                for (int c1 = 0; c1 < strlen(words[i][j]); c1++) {
                    printf("%c", words[i][j][c1]);
                }
                printf(" ");
            }
        }
        printf("\n");
    }
}

int main() {
    input();
    wordsDivider();
    finalOutput();
    return 0;
}