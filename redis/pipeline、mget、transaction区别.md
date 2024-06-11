# pipeline、mget、transaction区别

## 参考链接
https://www.jianshu.com/p/48a4d9e1a9a9

### pipeline
1. 使用 pipeline 命令，首先会在客户端本地缓存所有输入的命令，再打包发给服务端执行
2. 优点：减少RTT，客户端只会与服务端进行一次 RTT
3. 缺点：如果一次 pipeline 打包的命令过多，会造成网络堵塞

### mget 原生批量命令 与 pipeline 对比
1. 原生批量命令是原子性的，pipeline 非原子性
2. 原生批量命令是一个命令多多个 key；pipeine是多个命令
3. 原生批量命令是服务器支持实现的，而 pipeline需要服务端和客户端共同实现

### transaction
1. 通过 Multi 开始，Exec 执行，Exec 之前不会影响其他客户端执行命令
2. Exec 执行后，服务端将结果一并发送给客户端，命令太多会阻塞服务端

### pipeline 和 transaction
1. pipeline 关注 RTT，transaction 关注一致性
2. pipeline: 一次请求，服务端顺序执行，一次返回。
3. transaction: 多次请求（MULTI命令＋其他n个命令＋EXEC命令，所以至少是2次请求），服务端顺序执行，一次返回。

