; Assemble:	  nasm -f elf64 -l main_asm_linux.lst  main_asm_linux.asm
; Link:		  gcc -m64 -o main_asm_linux  main_asm_linux.o
; Run:		  ./main_asm_linux

global main
extern printf
extern scanf
extern exit

segment .data

	mode_1_str: db "1 signed",0Dh,0Ah,0
	mode_2_str: db "2 unsigned",0Dh,0Ah,0
	mode_3_str: db "3 exit",0Dh,0Ah,0
	enter_mode_str: db "Enter node: ",0Dh,0Ah,0

	mode_scanf_str: db "%d",0
	mode: dq 0

	enter_a_str: db "a: ", 0
	enter_b_str: db "b: ", 0
	
	enter_sword: db "%hd",0
	enter_uword: db "%hu",0
	print_sword: db "x: %hd",10,0
	;print_uword: db "x: %hu",10,0

	a_sig: dw 0
	b_sig: dw 0

	a_unsig: dw 0
	b_unsig: dw 0

	x_return: dw 0

segment .text

main:
	sub rsp, 28h		     
					
	@start:
	sub rsp, 8               ;теневая область для printf
	
	;1 Знаковый\n
	mov rdi, mode_1_str
	mov rax, 0
	call printf

	;2 беззнаковый\n
	mov rdi, mode_2_str
	mov rax, 0
	call printf

	;3 выход\n
	mov rdi, mode_3_str
	mov rax, 0
	call printf
	
	;Enter mode
	mov rdi, enter_mode_str
	mov rax, 0
	call printf

	add rsp, 8               ;очвобождение теневой области область для printf

	;Enter
	sub rsp, 8*2
	mov rdi, mode_scanf_str
	mov rsi, mode
	mov rax, 0
	call scanf
	add rsp, 8*2

	mov rax, 1
	cmp rax, [mode]
	je @m_1
	mov rax, 2
	cmp rax, [mode]
	je @m_2
	mov rax, 3
	cmp rax, [mode]
	je @next
	jmp @start
	
@m_1:
	sub rsp, 8*2              ;теневая область для printf
	;a: 
	mov rdi, enter_a_str
	call printf

	;enter a
	mov rdi, enter_sword
	mov rsi, a_sig
	call scanf

	;b: 
	mov rdi, enter_b_str
	call printf

	;enter b
	mov rdi, enter_sword
	mov rsi, b_sig
	call scanf

	call signedF
	
	;print x
	mov rdi, print_sword
	mov si, [x_return]
	call printf

	add rsp, 8*2
	jmp @start
@m_2:
	sub rsp, 8*2              ;теневая область для printf
	;a: 
	mov rdi, enter_a_str
	call printf

	;enter a
	mov rdi, enter_uword
	mov rsi, a_unsig
	call scanf

	;b: 
	mov rdi, enter_b_str
	call printf

	;enter b
	mov rdi, enter_uword
	mov rsi, b_unsig
	call scanf

	call unsignedF
	
	;print x
	mov rdi, print_sword
	mov si, [x_return]
	call printf

	add rsp, 8*2
	jmp @start
@next:


	mov esi, 0			   
	call exit
	ret

;main ENDP

;       signed
;;;;;;;;;;;;;;;;;;;;;;;;
;a == b
signed_s1:
	mov ax, 88d
	mov [x_return], ax
	ret
;signed_s1 ENDP

;a > b
signed_s2:
	sub ax, 99d

	cwd
	idiv bx

	mov [x_return], ax
	ret
;signed_s2 ENDP

;a < b
signed_s3:
	xor dx, dx
	mov ax, 64d
	idiv bx

	sub ax, [a_sig]
	mov [x_return], ax
	ret
;signed_s3 ENDP

signedF:
	mov ax, [a_sig]
	mov bx, [b_sig]

	cmp ax, bx
	je signed_s1 ;ax==bx
    jg signed_s2 ;ax>bx
	jl signed_s3 ;ax<bx
;signedF ENDP


;       unsigned
;;;;;;;;;;;;;;;;;;;;;;;;
;a == b
unsigned_s1:
	mov ax, 88d
	mov [x_return], ax
	ret
;unsigned_s1 ENDP

;a > b
unsigned_s2:
	sub ax, 99d

	xor dx, dx
	idiv bx

	mov [x_return], ax
	ret
;unsigned_s2 ENDP

;a < b
unsigned_s3:
	xor dx, dx
	mov ax, 64d
	idiv bx

	sub ax, [a_unsig]
	mov [x_return], ax
	ret
;unsigned_s3 ENDP

unsignedF:
	mov ax, [a_unsig]
	mov bx, [b_unsig]

	cmp ax, bx
	je unsigned_s1 ;ax==bx
    jg unsigned_s2 ;ax>bx
	jl unsigned_s3 ;ax<bx
;unsignedF ENDP