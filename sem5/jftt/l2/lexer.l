%{
    #include <stdio.h>
    #include <string.h>
    #include <stdlib.h>

    int line_count = 0;
    int word_count = 0;
%}

%%

(.*\n)   {
            char *line = yytext;
            int len = yyleng;

            // Remove the newline at the end
            line[len - 1] = '\0';

            // Remove leading whitespace
            while (*line == ' ' || *line == '\t') {
                line++;
            }

            // Remove trailing whitespace
            int line_len = strlen(line);
            while (line_len > 0 && (line[line_len - 1] == ' ' || line[line_len - 1] == '\t')) {
                line[line_len - 1] = '\0';
                line_len--;
            }

            // If the line is empty after trimming, skip it
            if (line_len == 0) {
                // Skip empty lines
            } else {
                // Replace sequences of spaces and tabs with one space
                char *src = line;
                char *dest = line;
                int in_whitespace = 0;
                while (*src != '\0') {
                    if (*src == ' ' || *src == '\t') {
                        if (!in_whitespace) {
                            *dest++ = ' ';
                            in_whitespace = 1;
                        }
                    } else {
                        *dest++ = *src;
                        in_whitespace = 0;
                    }
                    src++;
                }
                *dest = '\0';

                // Count words
                char *token;
                char *line_copy = strdup(line);
                if (line_copy == NULL) {
                    fprintf(stderr, "Memory allocation failed\n");
                    exit(1);
                }
                token = strtok(line_copy, " ");
                while (token != NULL) {
                    word_count++;
                    token = strtok(NULL, " ");
                }
                free(line_copy);

                // Print the processed line
                printf("%s\n", line);

                line_count++;
            }
        }

.|\n     { /* Ignore any other character */ }

%%

int main(int argc, char **argv) {
    yylex();
    printf("\nNumber of lines: %d\n", line_count);
    printf("Number of words: %d\n", word_count);
    return 0;
}

int yywrap() {
    return 1;
}
