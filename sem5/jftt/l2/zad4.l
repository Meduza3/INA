%option noyywrap

%{
    #include <stdbool.h>
    #include <math.h>

    #define MAX_STACK_SIZE 420

    typedef struct {
        int top;
        int elements[MAX_STACK_SIZE];
    } Stack; 

    
    /* Global Variables */
    static Stack calculatorStack;
    static int hasErrorOccurred = 0;

    void initStack(Stack* stack) {
        stack->top = -1;
    }

    int isEmpty(Stack* stack) {
        return stack->top == -1;
    }

    int isFull(Stack* stack) {
        return stack->top == MAX_STACK_SIZE - 1;
    }

    bool push(Stack* stack, int value) {
        if (isFull(stack)) {
            return false;
        }
        stack->elements[++stack->top] = value;
        return true;
    }

    bool pop(Stack* stack, int* value) {
        if (isEmpty(stack)) {
            return false;
        }
        *value = stack->elements[stack->top--];
        return true;
    }

    void displayError(const char* message) {
        if (!hasErrorOccurred) {
            printf("Error: %s\n", message);
            hasErrorOccurred = 1;
        }
    }

    void processNumber(Stack* stack, const char* yytext) {
        int value = atoi(yytext);
        if (!push(stack, value)) {
            displayError("Stack overflow");
        }
    }

    void processOperator(Stack* stack, char operatorSymbol) {
        int secondOperand, firstOperand, calculationResult;
        
        /* Check if enough operands are available */
        if (!pop(stack, &secondOperand) || !pop(stack, &firstOperand)) {
            displayError("Insufficient operands for operation");
            return;
        }

        /* Perform the requested operation */
        switch (operatorSymbol) {
            case '+':
                calculationResult = firstOperand + secondOperand;
                break;
            case '-':
                calculationResult = firstOperand - secondOperand;
                break;
            case '*':
                calculationResult = firstOperand * secondOperand;
                break;
            case '/':
                if (secondOperand == 0) {
                    displayError("Division by zero");
                    return;
                }
                calculationResult = firstOperand / secondOperand;
                break;
            case '^':
                calculationResult = pow(firstOperand, secondOperand);
                break;
            case '%':
                if (secondOperand == 0) {
                    displayError("Modulo by zero");
                    return;
                }
                calculationResult = firstOperand % secondOperand;
                break;
            default:
                displayError("Unknown operator");
                return;
        }

        /* Store the result */
        if (!push(stack, calculationResult)) {
            displayError("Stack capacity exceeded");
        }
    }

    void processInvalidInput(const char* yytext) {
        char error[100];
        snprintf(error, 100, "Invalid symbol \"%s\"", yytext);
        displayError(error);
    }

    void processEndOfLine(Stack* stack) {
        int result;
        if (!hasErrorOccurred && stack->top == 0 && pop(stack, &result)) {
            printf("= %d\n", result);
        } else if (stack->top > 0) {
            displayError("Too many operands remaining");
        } else {
            displayError("Insufficient operands");
        }

        hasErrorOccurred = 0;
        initStack(stack);
    }
%}

%%

-?[0-9]+                      { processNumber(&calculatorStack, yytext); }
[\+\-\*/\^\%]                 { processOperator(&calculatorStack, yytext[0]); }
[[:blank:]]+                  {/*ignore whitespace */}
\n                            { processEndOfLine(&calculatorStack); }
.                             { processInvalidInput(yytext); }

%%

int main() {
    initStack(&calculatorStack);

    printf("Please enter equations(one per line):\n");
    yylex();
}