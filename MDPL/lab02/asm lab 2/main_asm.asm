;.686
;.model flat, c

.data

extern a_sig: word
extern b_sig: word

extern a_unsig: word
extern b_unsig: word

extern x_return: word


.code

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