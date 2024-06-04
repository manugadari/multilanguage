#include <stdio.h>
#include <string.h>

void vulnerable_function(char *input) {
    char buffer[10];
    strcpy(buffer, input);
}

int main() {
    char input[100];
    printf("Enter a string: ");
    fgets(input, 100, stdin);
    vulnerable_function(input);
    return 0;
}
