import os
pid = os.fork()

if pid == 0:
  # 子のプロセス
  print("child process: pid = %d" % pid)
else:
  # 親のプロセス
  print("parent process: pid = %d" % pid)