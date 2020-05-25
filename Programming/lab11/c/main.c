#include <stdio.h>

int main() {
    int row = 0, column = 0;

    FILE *file_numbers, *file_result;
    char *filename_numbers = "numbers.txt";
    char *filename_result = "result.txt";

    // чтение из файла
    if ((file_numbers = fopen(filename_numbers, "r")) == NULL) {
        perror("Ошибка при открытии файла с данными, проверьте есть ли файл numbers.txt!");
        return 1;
    }
    fscanf(file_numbers, "%d %d\n", &row, &column);
    printf("Строк: %d Колонок: %d\n", row, column);

    if ((file_result = fopen(filename_result, "w")) == NULL) {
        perror("Ошибка при открытии/создании файла result.txt!");
        return 1;
    }

    float t = 0;
    for (int j = 0; j < row; j++) {
        float all[column];
        float max = -100000;
        for (int i = 0; i < column; i++) {
            t = 0;
            fscanf(file_numbers, "%f", &t);
            all[i] = t;
            if (t > max) {
                max = t;
            }
        }

        for (int i = 0; i < column; i++) {
            t = 0;
            all[i] = all[i] / max;
            fprintf(file_result,"%f ", all[i]);
        }
        fprintf(file_result,"\n");
    }
    fclose(file_numbers);
    fclose(file_result);
    printf("Результат сохранен в %s", filename_result);
}