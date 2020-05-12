#include <stdio.h>

int main(void) {
    FILE *mf;
    char str[50];
    char *estr;
    printf("Opening file...");
    mf = fopen("/home/lerman/Projects/Programming/lab11/c/numbers.txt", "r");

    if (mf == NULL) {
        printf("Error1\n");
        return -1;
    } else printf("Completed\n");

    printf("String readed:\n");
    while (1) {
        estr = fgets(str, sizeof(str), mf);

        if (estr == NULL) {
            if (feof(mf) != 0) {
                printf("\nReading file ended\n");
                break;
            } else {
                printf("\nError reading file\n");
                break;
            }
        }
        printf("     %s", str);
    }

    printf("Closing file");
    if (fclose(mf) == EOF) printf("Error\n");
    else printf("Completed\n");

    return 0;
}