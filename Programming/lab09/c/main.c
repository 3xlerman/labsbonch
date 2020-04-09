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
