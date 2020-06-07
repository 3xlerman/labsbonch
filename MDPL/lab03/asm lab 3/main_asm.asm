option casemap:none

includelib legacy_stdio_definitions.lib

extern MessageBoxA: PROC
extern ExitProcess: PROC

extern printf: proc
extern scanf:proc
extern _getch: proc

.DATA

	mode_1_str db "1 signed",10,0
	mode_2_str db "2 unsigned",10,0
	mode_3_str db "3 exit",10,0
	enter_mode_str db "Enter node: ", 0

	mode_scanf_str db "%i"
	mode dword 0

	enter_a_str db "a: ", 0
	enter_b_str db "b: ", 0
	
	enter_sword db "%hd",0
	enter_uword db "%hu",0
	print_sword db "x: %hd",10,0

	a_sig word 0
	b_sig word 0

	a_unsig word 0
	b_unsig word 0

	x_return word 0

.CODE

;mainCRTStartup
main PROC

	sub rsp, 28h		     ; резервируются место в стеке для первых четырех аргументов,
					         ; передаваемых через регистры. +8 байт наверно выравнивание.

@start:
	sub rsp, 8               ;теневая область для printf
	
	;1 Знаковый\n
	lea rcx, mode_1_str
	call printf

	;2 беззнаковый\n
	lea rcx, mode_2_str
	call printf

	;3 выход\n
	lea rcx, mode_3_str
	call printf
	
	;Enter mode
	lea rcx, enter_mode_str
	call printf

	add rsp, 8               ;очвобождение теневой области область для printf

	;Enter
	sub rsp, 8*2
	lea rcx, mode_scanf_str
	lea rdx, mode
	call scanf
	add rsp, 8*2

	cmp mode, 1
	je @m_1
	cmp mode, 2
	je @m_2
	cmp mode, 3
	je @next
	jmp @start
	
@m_1:
	sub rsp, 8*2              ;теневая область для printf
	;a: 
	lea rcx, enter_a_str
	call printf

	;enter a
	lea rcx, enter_sword
	lea rdx, a_sig
	call scanf

	;b: 
	lea rcx, enter_b_str
	call printf

	;enter b
	lea rcx, enter_sword
	lea rdx, b_sig
	call scanf

	call signedF
	
	;print x
	lea rcx, print_sword
	mov dx, x_return
	call printf

	add rsp, 8*2
	jmp @start
@m_2:
	sub rsp, 8*2              ;теневая область для printf
	;a: 
	lea rcx, enter_a_str
	call printf

	;enter a
	lea rcx, enter_uword
	lea rdx, a_unsig
	call scanf

	;b: 
	lea rcx, enter_b_str
	call printf

	;enter b
	lea rcx, enter_uword
	lea rdx, b_unsig
	call scanf

	call unsignedF
	
	;print x
	lea rcx, print_sword
	mov dx, x_return
	call printf

	add rsp, 8*2
	jmp @start
@next:
	
	;call _getch              ; Задержка закрытия

	mov ecx, 0			     ; ExitCode ; Функция ExitProcess ожидает 32-битный параметр
	;add rsp, 28h            ; Очистка стека
	call ExitProcess
	ret

 ;add rsp, 28h		;Это нужно было-бы, если бы  процессор сюда возвращался.

main ENDP


;       signed
;;;;;;;;;;;;;;;;;;;;;;;;
;a == b
signed_s1 PROC
	mov ax, 21d
	mov x_return, ax
	ret
signed_s1 ENDP

;a > b
signed_s2 PROC
	mov ax, a_sig
	add ax, 2
	cwd
	idiv b_sig

	mov x_return, ax
	ret
signed_s2 ENDP

;a < b
signed_s3 PROC
	mov ax, a_sig
	cwd
	imul b_sig
	mov bx, 4d
	idiv bx

	mov x_return, ax
	ret
signed_s3 ENDP

signedF PROC
	mov ax, a_sig
	mov bx, b_sig

	cmp ax, bx
	je signed_s1 ;ax==bx
    jg signed_s2 ;ax>bx
	jl signed_s3 ;ax<bx
signedF ENDP


;       unsigned
;;;;;;;;;;;;;;;;;;;;;;;;
;a == b
unsigned_s1 PROC
	mov ax, 21d
	mov x_return, ax
	ret
unsigned_s1 ENDP

;a > b
unsigned_s2 PROC
	mov ax, a_unsig
	add ax, 2
	xor dx, dx
	idiv b_unsig

	mov x_return, ax
	ret
unsigned_s2 ENDP

;a < b
unsigned_s3 PROC
	mov ax, a_unsig
	xor dx, dx
	imul b_unsig
	mov bx, 4d
	idiv bx

	mov x_return, ax
	ret
unsigned_s3 ENDP

unsignedF PROC
	mov ax, a_unsig
	mov bx, b_unsig

	cmp ax, bx
	je unsigned_s1 ;ax==bx
    jg unsigned_s2 ;ax>bx
	jl unsigned_s3 ;ax<bx
unsignedF ENDP

END