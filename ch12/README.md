# 第12章
## File Descriptor
カーネルは、新し くプロセスを作るたびに、各プロセスでどういった入出力が行われるかの管理テーブル を作ります。そのインデックス値がファイルディスクリプタです。

# 12.1


## BusyBox
  Alpine Linux

  ```
  docker build -t busybox busybox
  docker run -it busybox
  ```

## process group and session group

## Plan9 ?

## dscl
  list users
  ```
  dscl . -list /Users | grep -v '^_'
  ```

## SUID flag
  chownしたら、SUID flagはクリアされる

## capability

# 12.2

## command line args


# 12.3

# 12.4

# 12.5

