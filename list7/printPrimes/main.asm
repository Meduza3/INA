@ arm-none-eabi-as main.asm -o main.o && arm-none-eabi-gcc -specs=rdimon.specs main.o
@ qemu-arm-static a.out

.data
format_str: .asciz "%d\n"       @ format for printf

.text
.global main
.extern printf
main:
    ldr     r4, =#100000        @ set loop upper limit
    mov     r5, #0              @ set loop counter to 0

main_loop:
    add     r5, r5, #1          @ increment the loop counter
    cmp     r5, r4              @ check if loop counter reached max number
    beq     main_finished       @ if equal, finish the program by branching to main_finished

    mov     r0, r5              @ set current number as the argument for isPrime
    bl      isPrime             @ call isPrime
    cmp     r1, #1              @ check if the isPrime return variable is 1
    beq     main_printNumber    @ if equal, print

    b       main_loop           @ go back to the loop

main_printNumber:
    ldr     r0, =format_str     @ set the format for printing
    mov     r1, r5              @ pass the current number as the value to print
    bl      printf              @ call printf
    b       main_loop           @ go back to the loop

main_finished:
    mov     r7, #1              @ syscall for exiting
    mov     r0, #0              @ exit code
    swi     0                   @ make the syscall


isPrime:
    push    {r5, r6, r7, lr}        @ push the values to the stack to save them

    cmp     r0, #2                  @ compare the number to 2
    blt     notPrime                

    mov     r7, r0                  @ move the number to r7
    lsr     r5, r0, #1              @ r5 = (r0 >> 1), więc r5 = r5/2

    mov     r6, #2                  @ set the loop counter to 2
loop:
    cmp     r6, r5                  @ check if counter > n/2
    bgt     prime

    mov     r0, r7                  @ load the number
    mov     r1, r6                  @ load the counter as a divisor
    bl      __aeabi_uidivmod        @ divide and get the result in r0 and the remainder in r1
    cmp     r1, #0                  @ see if the remainder is zero
    beq     notPrime                @ if so, the number is not prime

    add     r6, r6, #1              @ Increment loop
    b       loop                    
prime:
    mov     r1, #1                  @ set r1 = 1
    b       finished                

notPrime:
    mov     r1, #0                  @ set r1 = 0

finished:
    pop     {r5, r6, r7, lr}        
    bx      lr                      
