.global _main
.p2align 3

_main:
    adr X0, msg             // Load address of the message into x0
    bl      _puts           // Call puts(message)
    b _exit
msg:
    .asciz "Hola, mundo"    // asciz puts a 0 byte at the end