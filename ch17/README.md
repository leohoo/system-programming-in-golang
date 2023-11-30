# 17.1
## Names of parentheses by chatgpt
1. **Round Parentheses:**
   - *English:* `()` are called parentheses.
   - *Japanese:* "丸括弧" (maru kakko).

2. **Curly Braces:**
   - *English:* `{}` are called curly braces or curly brackets.
   - *Japanese:* "中括弧" (nakakukko) or "波括弧" (nami kakko).

3. **Square Brackets:**
   - *English:* `[]` are called square brackets or brackets.
   - *Japanese:* "角括弧" (kakukakko).

4. **Angle Brackets:**
   - *English:* `<>` are called angle brackets or chevrons.
   - *Japanese:* "山括弧" (yama kakko).

## hello.s
```assembly
    .global _start          # このラベルが外部公開されるという宣言

    .text                   # これから先はプログラムが始まる宣言
_start:
    mov $1, %rax            # writeシステムコール(1)
    mov $1, %rdi            # 出力先ファイルハンドル(1: stdout)をセット
    mov message(%rip), %rsi # 出力したいメッセージアドレスをセット
    mov $13, %rdx           # 文字数をセット
    syscall                 # writeの実行
    mov $60, %rax           # exitシステムコール
    xor %rdi, %rdi          # リターンコード(ここではエラーなし、0)をセット
    syscall                 # 終了呼び出し

message:                    # データ領域のラベル(いわゆる定数名)
    .ascii "Hello, world\n" # テキストデータ
```
**Correct `mov message(%rip), %rsi`**
- `lea message(%rip), %rsi`
- `mov $message, %rsi`

## Explanation

The `lea` (Load Effective Address) instruction and the `mov` (Move) instruction serve different purposes in x86-64 assembly language.

1. **`lea message(%rip), %rsi`**

   - `lea` calculates the effective address of the source operand (in this case, `message(%rip)`) and loads it into the destination operand (`%rsi`).
   - It doesn't actually load the content at the memory location but computes the address and places it in the specified register.
   - It is often used for address calculations without accessing the memory.

2. **`mov message(%rip), %rsi`**

   - `mov` directly moves the content at the memory location specified by the source operand (`message(%rip)`) into the destination operand (`%rsi`).
   - This is used for copying the actual value from one location to another.
   - However, the syntax `mov message(%rip), %rsi` is incorrect because `mov` doesn't support immediate addressing in this way. The correct syntax would involve using `mov rsi, message(%rip)` or using `lea` for address calculation.

In the corrected code, `lea message(%rip), %rsi` is used because the goal is to load the address of the `message` label into the `%rsi` register, preparing it for the `write` system call, which expects the address of the data to be written.

In summary, `lea` is often used for address calculations, while `mov` is used for moving data between registers and memory. The correct choice depends on the specific task you want to accomplish. In this case, `lea` is the appropriate choice for loading the address of the `message` label.

The `mov $message, %rsi` and `lea message(%rip), %rsi` instructions serve different purposes in x86-64 assembly language.

1. **`mov $message, %rsi`**

   - This instruction directly loads the immediate value (address) represented by the label `message` into the `%rsi` register.
   - It treats `message` as an immediate value and copies that value directly into the register.
   - This is often used when you want to load a specific address into a register.

2. **`lea message(%rip), %rsi`**

   - The `lea` (Load Effective Address) instruction calculates the effective address of the source operand (`message(%rip)`) and loads it into the destination operand (`%rsi`).
   - It does not actually access memory but computes the address, making it useful for performing address calculations without accessing the data.
   - It is often used when you want to obtain the address of a label or variable for use in address arithmetic.

In the context of the original code, either `mov $message, %rsi` or `lea message(%rip), %rsi` can be used to load the address of the `message` label into the `%rsi` register before the `write` system call. Both achieve the same goal in this specific context.

The choice between them often depends on coding conventions, personal preference, or specific requirements. The `lea` instruction is more versatile and is commonly used for address calculations, while `mov` is more straightforward for directly loading values into registers.

## Check binary
```
readelf --syms a.out
readelf --file-header a.out
```


## hello for Apple silicon Mac

https://stackoverflow.com/a/76972251

```
clang -e _start -o hello hello-apple.s
./hello
# list symbols
nm hello
```

https://github.com/below/HelloSilicon

# 17.2
## Runtime
## Linker

# 17.3

# 17.4
```
$ realpath $(which go)
/opt/homebrew/Cellar/go/1.20.6/libexec/bin/go

$ find /opt/homebrew/Cellar/go/1.20.6 -name rt0_darwin_amd64.s
/opt/homebrew/Cellar/go/1.20.6/libexec/src/runtime/rt0_darwin_amd64.s

$ cat -n /opt/homebrew/Cellar/go/1.20.6/libexec/src/runtime/rt0_darwin_amd64.s
     1	// Copyright 2009 The Go Authors. All rights reserved.
     2	// Use of this source code is governed by a BSD-style
     3	// license that can be found in the LICENSE file.
     4
     5	#include "textflag.h"
     6
     7	TEXT _rt0_amd64_darwin(SB),NOSPLIT,$-8
     8		JMP	_rt0_amd64(SB)
     9
    10	// When linking with -shared, this symbol is called when the shared library
    11	// is loaded.
    12	TEXT _rt0_amd64_darwin_lib(SB),NOSPLIT,$0
    13		JMP	_rt0_amd64_lib(SB)
```

# 17.5

```
$ python -h
...
-v     : verbose (trace import statements); also PYTHONVERBOSE=x
         can be supplied multiple times to increase verbosity
...
```

