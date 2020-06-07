.data

    public my_array
    my_array dq 256 DUP(0)
    arr_size dq 256

.code

asmF proc
    mov r13, rdx       ;������� ��������� d (2'��) �  r13

	xor r10, r10       ;counter = 0
    xor r8, r8         ;i = 0

    ;r12 = arr_size
    mov r12, arr_size  ;��������� ������� �������

    @L1whileIf:        ;����� ������ �����
    cmp r8, r12        ;i == arr_size -> ����� �� �����
    je @L2End

    ;if
    mov rax, 8         ;������ ����������
    xor rdx, rdx       ;��������� ����� *
    mul r8             ;8*i
    lea rbx, my_array  ;��������� ������ �������
	add rax, rbx       ;�������� ������ ������� �� ��������� ��������
    mov rbx, [rax]     ;��������� �������� my_array[i]

    cmp rcx, rbx       ;c <= a[i]
    jg @L3EndIf
	cmp rbx, r13       ;a[i] <= d
    jg @L3EndIf

    inc r10            ;l++

    @L3EndIf:
    inc r8             ; --i
    jmp @L1whileIf     ;������ � ������ �����

    @L2End:            ;����� ����� �����
    
    mov rax, r10       ;������� ���������� �������
    ret
asmF endp
end