strlen:
    push      ebx
    mov       ebx,  eax

.nextChar:
    cmp       byte [eax], 0
    jz        .finished
    inc       eax
    jmp       .nextChar

.finished:
    sub       eax, ebx
    pop       ebx
    ret


sprint:
    push      edx
    push      ecx
    push      ebx
    push      eax

    call      strlen

    mov       edx, eax
    pop       eax

    mov       ecx, eax
    mov       ebx, 1
    mov       eax, 4
    int       80h

    pop       ebx
    pop       ecx
    pop       edx
    ret




printLF:
    push    eax
    mov     eax, 0Ah
    push    eax         ; push eax with LF character to the stack
    mov     eax, esp    ; use esp to get the adress(which is required by write syscall)
    call    sprint
    pop     eax
    pop     eax
    ret

sprintLF:
    call    sprint
    call    printLF
    ret


;------------------------------------------------------------------------------------
; void printInt(Integer number)
; prints integer
printInt:
    ; push to stack to preserve the values
    push    eax
    push    ecx
    push    edx
    push    esi
    mov     ecx, 0      ; will count how many bytes to print

.divideLoop:
    inc     ecx
    mov     edx, 0
    mov     esi, 10
    idiv    esi         ; divides eax by esi, eax = result, edx = remainder
    add     edx, 48     ; make the remainder an ASCII int
    push    edx         ; push the string representation of the remainder onto the stack
    cmp     eax, 0      ; check if the result can be divided more
    jnz     .divideLoop

.printLoop:
    dec     ecx         ; decrease the bytes to print counter
    mov     eax, esp
    call    sprint    ; print the digit
    pop     eax         ; remove this digit from the stack
    cmp     ecx, 0
    jnz     .printLoop

    pop     esi
    pop     edx
    pop     ecx
    pop     eax
    ret


;------------------------------------------------------------------------------------
; void printIntLF(Integer number)
; prints integer with newline
printIntLF:
    call    printInt
    call    printLF
    ret





getHexDigit:
    cmp     eax, 9
    jg      .letter 

    add     eax, 48 
    jmp     .number

.letter:
    add     eax, 55 

.number:
    ret












printHex:
    push    eax
    push    ecx
    push    edx
    push    esi
    mov     ecx, 0      ; counter

.divideLoop:
    inc     ecx
    mov     edx, 0
    mov     esi, 16
    idiv    esi         ; eax = result, edx = remainder

    push    eax
    mov     eax, edx
    call    getHexDigit
    mov     edx, eax
    pop     eax

    push    edx         ; pushing ready char
    cmp     eax, 0      ; check if furhter division
    jnz     .divideLoop

.printLoop:
    dec     ecx         ; decrease counter
    mov     eax, esp
    call    sprint    ; print the digit
    pop     eax         ; remove this digit from the stack
    cmp     ecx, 0
    jnz     .printLoop

    pop     esi
    pop     edx
    pop     ecx
    pop     eax
    ret


;------------------------------------------------------------------------------------
; void printHexLF(Integer decimalNumber)
; takes a decimal number and prints it out as a hexadecimal number with newline
printHexLN:
    call    printHex
    call    printLF
    ret


;------------------------------------------------------------------------------------
; int stringToInt(String text)
; takes a string and converts it to an int
stringToInt:
    push    ebx,
    push    ecx,
    push    edx,
    push    esi,
    mov     esi, eax
    mov     eax, 0          ; eax will hold the int
    mov     ecx, 0          ; ecx is the loop counter

.multiplyLoop:
    ; bl is a part of ebx, which only takes 8 bits and this is the size of 1 ASCII character
    xor     ebx, ebx        ; set ebx to 0
    mov     bl, [esi+ecx]
    cmp     bl, 48
    jl      .finished
    cmp     bl, 57
    jg      .finished

    sub     bl, 48          ; bl = bl - 48   ebx's lower half gets converted to decimal representation of ascii value
    add     eax, ebx        ; eax = eax + ebx = eax + digit
    mov     ebx, 10
    mul     ebx             ; eax = eax * ebx = eax * 10
    inc     ecx             ; increment the counter
    jmp     .multiplyLoop 

.finished:
    cmp     ecx, 0          ; check if the counter == 0     that means that no ints were provided
    je      .restore
    mov     ebx, 10
    div     ebx             ; eax = eax / ebx = eax / 10

.restore:
    pop     esi,
    pop     edx,
    pop     ecx,
    pop     ebx,
    ret