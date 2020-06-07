option casemap:none
.686
.model flat, c

.data

extern a_f: real4
extern result_f: real4
extern c_i: real4
extern d_i: dword

.code

;1 (3*c+8-d)/(a-c/4)

;СП - сопроцессор
ASMCalc proc
	;вторая часть
    fild c_i                  ;загрузка в a вершину стека СП

	;fld/fild в masm паботают ТОЛЬКО с памятью (вот сволочи)
	;поэтому ложу в стек и забираю с него
	;вот такой вот костыль
	mov eax, 4d               ;eax = 4
    push eax                  ;eax в стек
	fild dword ptr [esp]      ;берется значение 4 изи памяти и ложится в СП
	pop eax                   ;очистка стека
	fdiv                      ;[c/4]
	fld a_f
	fxch st(1)               ;поменять местами
	fsub

	fstp result_f             ;результат вычисления 2'ой часть ложится в переменную

	;первая часть
	mov eax, 3d               ;eax = 3
    push eax                  ;eax в стек
	fild dword ptr [esp]      ;берется значение 3 изи памяти и ложится в СП
	pop eax                   ;очистка стека
	fild c_i
	fmul

	mov eax, 8d               ;eax = 8
    push eax                  ;eax в стек
	fild dword ptr [esp]      ;берется значение 3 изи памяти и ложится в СП
	pop eax                   ;очистка стека
	fadd
	fild d_i
	fsub

	fdiv result_f

	fstp result_f             ;возврат результата
    ret
ASMCalc endp
end