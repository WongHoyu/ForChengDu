# panic 和 recovery

## 参考链接
    1. https://www.cnblogs.com/tortoise512/articles/15308339.html
    2. https://www.qiyacloud.cn/2021/09/2021-09-02/

### panic 步骤
1. goroutine有两个重要字段
    ```go
    type g struct {
        // ...
        _panic         *_panic // panic 链表，这是最里的一个
        _defer         *_defer // defer 链表，这是最里的一个；
        // ...
    }
    ```

2. 当发生 panic 时，会调用 gopanic函数
   1. 遍历 goroutine 的 defer 链表，获取到一个 _defer 延迟函数；
   2. 获取到的 _defer 延迟函数，设置标识 d.started，绑定当前 d._panic
   3. 执行 _defer 延迟函数
   4. 摘掉执行完的 _defer 延迟函数
   5. 判断 _panic.recovered 是否设置为 true，进行相应的操作;
       - 如果是 true，那么重置 pc，sp 寄存器（一般从 deferreturn 指令前执行）, goroutine 投递到调度队列，等待执行
       - 否则，重复以上步骤 

3. 为什么 recover() 要搭配上 defer
   1. 因为不用 defer，recover() 要不在 panic 前执行，是无法获取到所在 goroutine 的 _panic链表的值，所以无法捕获 panic；
   2. 又或者定义在 panic 之后，是无法执行 recover()
   3. recover() 还得搭配 defer func() {} 匿名函数使用，因为 defer 入队列前就确认好值，如果 defer recover() 入链表时就确认好当前
goroutine 的 _panic 的值了，也无法捕获 panic