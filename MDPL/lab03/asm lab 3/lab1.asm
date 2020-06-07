
; Assemble:	  nasm -f elf64 -l lab1.lst  lab1.asm
; Link:		  gcc -m64 -o lab1  lab1.o
; Run:		  ./lab1
	global main
    extern printf
	extern exit
section .data
	test: db 1010101b
    msg:    db 'Hello World!',0Dh,0Ah,0
    frm:    db '%s',0
    frm2:   db 'summa = %lld',0x0D,0Ah,0
    var1:   dq 1234
    var2:   dq 99999

section .text
	

main:
	push rbp
	mov rdi, frm
            mov rsi, msg
            mov rax, 0
            call printf

            
            mov rax,[var1]
            mov [var2],rax
            add rax,[var2]
            
            mov rdi, frm2
            mov rsi, rax
            mov rax, 0
            call printf
            
	mov 	rdi,0
	call 	exit
	pop rbp
	ret
	

