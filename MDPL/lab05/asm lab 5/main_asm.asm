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

;�� - �����������
ASMCalc proc
	;������ �����
    fild c_i                  ;�������� � a ������� ����� ��

	;fld/fild � masm �������� ������ � ������� (��� �������)
	;������� ���� � ���� � ������� � ����
	;��� ����� ��� �������
	mov eax, 4d               ;eax = 4
    push eax                  ;eax � ����
	fild dword ptr [esp]      ;������� �������� 4 ��� ������ � ������� � ��
	pop eax                   ;������� �����
	fdiv                      ;[c/4]
	fld a_f
	fxch st(1)               ;�������� �������
	fsub

	fstp result_f             ;��������� ���������� 2'�� ����� ������� � ����������

	;������ �����
	mov eax, 3d               ;eax = 3
    push eax                  ;eax � ����
	fild dword ptr [esp]      ;������� �������� 3 ��� ������ � ������� � ��
	pop eax                   ;������� �����
	fild c_i
	fmul

	mov eax, 8d               ;eax = 8
    push eax                  ;eax � ����
	fild dword ptr [esp]      ;������� �������� 3 ��� ������ � ������� � ��
	pop eax                   ;������� �����
	fadd
	fild d_i
	fsub

	fdiv result_f

	fstp result_f             ;������� ����������
    ret
ASMCalc endp
end