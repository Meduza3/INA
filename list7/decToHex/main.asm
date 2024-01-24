%include "../functions.asm"

section .data
welcomeMsg  db "Please enter a number: ", 0

section .bss
userInput:  resb 255

section .text
global _start

_start:
    mov     eax, welcomeMsg ; Load address of welcomeMsg to eax
    call    sprint

    ; invoke SYS_READ
    mov     edx, 255 ; user input size
    mov     ecx, userInput ;
    mov     ebx, 0 ; input descriptor
    mov     eax, 3 ; SYS_READ into eax
    int     80h ; Reads input from user

    mov     eax, userInput ;
    call    stringToInt
    call    printHexLN
    
    mov       ebx, 0 ; return code
    mov       eax, 1 ; exit
    int       80h