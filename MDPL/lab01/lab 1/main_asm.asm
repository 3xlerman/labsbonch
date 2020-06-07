.data

	EXTERN s8_a: byte
	EXTERN u8_a: byte
	EXTERN s16_a: word
	EXTERN u16_a: word

	EXTERN s8_c: byte
	EXTERN u8_c: byte
	EXTERN s16_c: word
	EXTERN u16_c: word
	
	EXTERN s8_d: byte
	EXTERN u8_d: byte
	EXTERN s16_d: word
	EXTERN u16_d: word

.code

;var 1 (3*c+8-d)/(a-c/4)
calc_s8 PROC
    ;(a-c/4)
	mov al, s8_c
	cbw           ;al->ax
	mov bx, 4d
	cwd           ;копирование старшего бита ax во все байты dx
	idiv bx

	mov bx, ax
	mov al, s8_a
	cbw           ;al->ax
	sub ax, bx

	mov di, ax    ;результат 2 части

	;(3*c+8-d)
	mov al, s8_d
	cbw           ;al->ax
	mov bx, 8d
	sub bx, ax
	mov si, bx

	mov al, s8_c
	cbw           ;al->ax
	mov bx, ax
	mov ax, 3d
	xor dx, dx
	imul bx
	add ax, si

	cwd           ;копирование старшего бита ax во все байты dx
	idiv di

	ret
calc_s8 endp

calc_u8 PROC
	;(a-c/4)
	xor ax, ax
	mov al, u8_c
	mov bx, 4d
	xor dx, dx
	div bx

	xor bx, bx
	mov bl, u8_a
	sub bx, ax

	mov di, bx    ;результат 2 части

	;(3*c+8-d)
	mov ax, 8d
	xor bx, bx
	mov bl, u8_d
	sub ax, bx
	mov si, ax

	mov ax, 3d
	xor bx, bx
	mov bl, u8_c
	xor dx, dx
	mul bx
	add ax, si

	cwd           ;копирование старшего бита ax во все байты dx
	idiv di

	ret

	ret
calc_u8 endp

calc_s16 PROC
	;(a-c/4)
	mov ax, s16_c
	mov bx, 4d
	cwd           ;копирование старшего бита ax во все байты dx
	idiv bx
	mov bx, s16_a
	sub bx, ax
	mov di, bx	

	;(3*c+8-d)
	mov ax, 8d
	mov bx, s16_d
	sub ax, bx
	mov si, ax

	mov ax, 3
	mov bx, s16_c
	xor dx, dx
	imul bx

	add ax, si

	cwd           ;копирование старшего бита ax во все байты dx
	idiv di


	ret
calc_s16 endp

calc_u16 PROC
	;(a-c/4)
	mov ax, u16_c
	mov bx, 4d
	xor dx, dx
	div bx
	mov bx, u16_a
	sub bx, ax
	mov di, bx	

	;(3*c+8-d)
	mov ax, 8d
	mov bx, u16_d
	sub ax, bx
	mov si, ax

	mov ax, 3
	mov bx, u16_c
	xor dx, dx
	mul bx

	add ax, si

	cwd           ;копирование старшего бита ax во все байты dx
	idiv di
	
	ret
calc_u16 endp

end