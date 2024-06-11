# GoLang基础

### 参考链接
https://github.com/mao888/golang-guide/blob/main/golang/go-Interview/GOALNG_INTERVIEW_COLLECTION.md

1. New 和 Make 的区别
   1. Make 只用于 channel、slice、map，返回的是结构体本身
   2. New 返回的是指针，但只会对结构体开辟内存，不会对结构体内的指针再开辟内存

2. GC：https://studygolang.com/articles/28643
   - 第一个阶段 gc开始 （stw） 
     - stop the world 暂停程序执行
     - 启动标记工作协程（ mark worker goroutine ），用于第二阶段
     - 启动写屏障
     - 将root 跟对象放入标记队列（放入标记队列里的就是灰色）
     - start the world 取消程序暂停，进入第二阶段
   - 第二阶段 marking（这个阶段，用户程序跟标记协程是并行的）
     - 从标记队列里取出对象，标记为黑色
     - 然后检测是否指向了另一个对象，如果有，将另一个对象放入标记队列
     - 在扫描过程中，用户程序如果新创建了对象 或者修改了对象，就会触发写屏障，将对象放入单独的 marking队列，也就是标记为灰色
     - 扫描完标记队列里的对象，就会进入第三阶段
   - 第三阶段 处理marking过程中修改的指针 （stw）
     - stop the world 暂停程序
     - 将marking阶段 修改的对象 触发写屏障产生的队列里的对象取出，标记为黑色
     - 然后检测是否指向了另一个对象，如果有，将另一个对象放入标记队列
     - 扫描完marking队列里的对象，start the world 取消暂停程序 进入第四阶段
   - 第四阶段 sweep 清除白色的对象
     - 到这一阶段，所有内存要么是黑色的要么是白色的，清楚所有白色的即可
     - golang的内存管理结构中有一个bitmap区域，其中可以标记是否“黑色”

   - golang gc 的触发
     - 触发内存的阈值，内存达到上次gc后的2倍
     - 2分钟
     - 手动触发
3. Channel：https://juejin.cn/post/7037656471210819614
   1. Channel 底层结构体：HChann, 有一个buf缓冲区，是一个循环数组；
   2. 一个lock防止多个线程并发读写；
   3. 有一个写指针指向下一个将要写入的下标；
   4. 有一个读指针指向下一个协程要读的下标； 
   5. 还有写队列和读队列，当满缓冲的时候，写的goroutine会放进写队列中等待，等buf有空间时，会将写队列里的协程放到P队列中，等待CPU调度；当buf为空的时候，会将读的goroutine放进读队列中，等有生产者写入到Channel时直接消费。
4. sync.Mutex 互斥锁：https://juejin.cn/post/7032457568433733639
   1. 底层由 Mutex 结构体来实现
   2. 有两个字段，sema 是信号量，当锁释放时，会释放信号量来唤醒等待队列中的协程来获取锁；另一个字段是 state 状态，int32，不同的比特位代表着不同的含义，比如有记录等待锁的队列中的goroutine数量、是否是饥饿模式、锁是否被持有等
   3. 饥饿模式为0时，代表是正常模式，队列中的goroutine来抢占时，并不会直接获取锁，而是会和新请求进来的goroutine一起抢占锁，往往新请求的goroutine更容易获得锁。
   4. 当队列中的goroutine等待超过1ms时就会切换为饥饿模式，那么新请求的goroutine并不会直接参与竞争锁，而是进入队列排队等待锁的释放。
   5. 重复加锁会panic；未加锁就 unlock 也会panic；不同协程可以调用 unlock 释放别的协程的lock；
5. sync.RWMutex 读写锁：https://juejin.cn/post/7035183181393297422
   1. 底层由 RWMutex 结构体实现
   2. RWMutex 复用了 Mutex 的结构，其次多了 readerSem 读信号量、writerSem 写信号量、readerCount 统计读锁的数量、readerWait 写锁等待读锁的数量
   3. 当加上读锁时，readerCount 会 +1，当 readerCount > 0 时，不允许加写锁
   4. 当加了写锁时，readerCount 会等于 -1，此时不允许再加任何的读锁、写锁
   5. 当readerCount > 0时，写锁需要等待读锁全部释放，等待时会将 readerCount的数量复制到readerWait，当读锁释放时，readerCount、readerWait都会 -1，直到readerWait为0时，会唤醒要加写锁的 goroutine，并加锁
6. Map: https://juejin.cn/post/7029679896183963678
   1. Map底层是由 hmap 组成
   2. hmap 里有 buckets、oldbuckets、B，buckets 是存储数据的一个指针，oldbuckets是重哈希时暂存数据的一个指针
   3. B 代表bucket有 2^B 的长度
   4. bucket指向的数组是由 bmap 这个结构体组成的，这个数组一个有8个长度，当哈希冲突时，用拉链法解决，也就是 bmap 有一个 overflow 的指针指向另一个 bmap数组
   5. 查找方法：将 key 进行 hash，取得 64位的比特值，取后4位（若此时B = 4）的比特，找到 buckets 对应的下标，比如后四位为 0101(10进制为5)，则取下标为5的那个桶，然后用前 8 位（01010101），从桶中的 bmap 找到一样比特位的那一行数据（也是 01010101），再比对完整的 hash 值，如果一致则返回。
   6. 若不一致，则根据 overflow 指针，找到下一个 bmap 数组，重新比对前8位和完整比特位
   7. map 不能并发读写