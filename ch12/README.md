# 第12章
- File Descriptor

  カーネルは、新し くプロセスを作るたびに、各プロセスでどういった入出力が行われるかの管理テーブル を作ります。そのインデックス値がファイルディスクリプタです。

# 12.1

- 実行ファイル名とシンボリックリンク
## BusyBox
  Alpine Linux

  ```
  docker build -t busybox busybox
  docker run -it busybox
  ```

## process group and session group

## Plan9
[Plan 9 from Bell Labs](https://ja.wikipedia.org/wiki/Plan_9_from_Bell_Labs)
## dscl
  list users
  ```
  dscl . -list /Users | grep -v '^_'
  ```

## SUID flag
  chownしたら、SUID flagはクリアされる

## capabilities
https://ja.wikipedia.org/wiki/Capability-based_security

# 12.2

## command line args


# 12.3

```
cat /proc/{process id}/cmdline
```

## github.com/shirou/gopsutil/process not found
> no required module provides package github.com/shirou/gopsutil/process: go.mod file not found in current directory or any parent directory; see 'go help modules'
  ```
  go mod init golang.sys/ch12
  go get github.com/shirou/gopsutil/process
  ```

# 12.4

# 12.5

# 12.6

# 12.7

fork

Go 1.9からは、Linuxでamd64アーキテクチャの場合にのみ、CLONE_VFORKと CLONE_VM を付与した clone() システムコールを使っています。


posix_spawn is a function provided by the POSIX (Portable Operating System Interface) standard, and it's used to spawn (or create) processes. It's an alternative to the traditional fork() and exec() combination commonly used in UNIX-like systems to create new processes. The primary advantage of posix_spawn over fork() and exec() is its efficiency, especially in situations where the parent process consumes a large amount of memory.

Memoryのcopy-on-writeとGC

daemon()

# 12.8

- vfork()
- rfork()
- clone()

# 12.9
