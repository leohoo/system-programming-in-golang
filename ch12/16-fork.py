import os
pid = os.fork()

if pid == 0:
  # 子のプロセス
  print("in child process: pid = %d" % pid)
else:
  # 親のプロセス
  print("in parent process: pid = %d" % pid)