; nasm -f elf -g main.asm && gcc -m32 main.o

section .data
f1Output        db "f1: %lf", 10, 0
f2Output        db "f2: %lf", 10, 0
f3Output        db "f3: %lf", 10, 0
f4Output        db "f4: %lf", 10, 0

a               dq 9.0                  ; double(8 bytes)
b               dd 6.0                  ; float(4 bytes)
c               dq 5.0                 ; double(8 bytes)
d               dd 4                    ; int(4 bytes)
result          dq 0.0

section .text

extern printf
global main


; returns a / b - c * d

f1:
    fild    dword [edx]         ; load integer d
    fmul    qword [ecx]         ; st0 = c * d

    fld     qword [eax]         ; load a
    fdiv    dword [ebx]         ; st0 = a / b

    fsub                        ; st0 = a / b - c * d

    fst     qword [result]      ; load the result to eax
    ret


;------------------------------------------------------------------------------------
; double f2(double a, int b)
; returns log_b(a) in result
; parameters: eax = a, ebx = b
f2:
    ; push 1 * log2(a) to stack
    fld1
    fld         qword [eax]
    fyl2x

    ; push 1 * log2(b) to stack
    fld1
    fild        dword [ebx]
    fyl2x

    fdiv                        ; log2(a) / log2(b) = logb(a)

    fst         qword [result]
    ret


;------------------------------------------------------------------------------------
; double f3(double a, int b)
; returns b^a in result
; parameters: eax = a, ebx = b
; this function uses the fact that 2^(a * log2(b)) = b^a
f3:
    fld         qword [eax]
    fild        dword [ebx]

    fyl2x                       ; st0 = st1 * log2(st0) = a * log2(b)
    fld         st0             ; duplicate

    frndint                     ; get integer part
    fxch                        ; swap st0 with st1
    fsub        st0, st1        ; get fractional part
    f2xm1                       ; 2^frac(a * log2(b))
    fld1
    faddp

    fxch                        ; swap st0 with st1
    fld1
    fscale                      ; 1 * 2^int(a * log2(b))
    fstp        st1             ; pop 1 from the stack
    fmulp                       ; 2^frac(a * log2(b)) * 2^int(a * log2(b)) = 2^(a * log2(b)) = b^a

    fst         qword [result]
    ret


;------------------------------------------------------------------------------------
; double f4(float a, int b)
; returns the b-th root of a in result
; parameters: eax = a, ebx = b
; this function uses the fact that 2^(a * log2(b)) = b^a
f4:
    fld1
    fild        dword [ebx]
    fdiv                        ; b = 1 / b

    fld         dword [eax]

    fyl2x                       ; st0 = st1 * log2(st0) = 1/b * log2(a)
    fld         st0             ; duplicate

    frndint                     ; get integer part
    fxch                        ; swap st0 with st1
    fsub        st0, st1        ; get fractional part
    f2xm1                       ; 2^frac(1/b * log2(a))
    fld1
    faddp

    fxch                        ; swap st0 with st1
    fld1
    fscale                      ; 1 * 2^int(1/b * log2(a))
    fstp        st1             ; pop 1 from the stack
    fmulp

    fst         qword [result]
    ret


main:
    finit

    mov         eax, a
    mov         ebx, b
    mov         ecx, c
    mov         edx, d
    call f1
    ; printf
    push        dword [result + 4]      ; push the lower 4 bytes
    push        dword [result]
    push        f1Output
    call        printf
    add         esp, 12

    mov         eax, a
    mov         ebx, d
    call        f2
    ; printf
    push        dword [result + 4]
    push        dword [result]
    push        f2Output
    call        printf
    add         esp, 12

    mov         eax, a
    mov         ebx, d
    call        f3
    ; printf
    push        dword [result + 4]
    push        dword [result]
    push        f3Output
    call        printf
    add         esp, 12

    mov         eax, b
    mov         ebx, d
    call        f4
    ; printf
    push        dword [result + 4]
    push        dword [result]
    push        f4Output
    call        printf
    add         esp, 12

    mov ebx, 0
    mov eax, 1
    int 80h


