#0拷贝技术

- 零拷贝就是一种避免 CPU 将数据从一块存储拷贝到另外一块存储的技术。总之就为了CPU尽量减少拷贝操作，交由DMA去做。

- 参考URL：
    1. https://blog.csdn.net/weixin_42096901/article/details/103017044

- 传统方式
    ![avatar](../img/3211622308476_.pic_hd.jpg)
    
- 使用mmap后
    ![avatar](../img/3221622308516_.pic_hd.jpg)
    
- 使用sendfile
    ```C
    #include<sys/sendfile.h>
    ssize_t sendfile(int out_fd, int in_fd, off_t *offset, size_t count);
    ```
    1. sendfile 的缺点在于 out_fd 必须为 socket
    ![avatar](../img/3231622308556_.pic_hd.jpg)
   
    2. 复制流程
        1. 用户进程发起数据读取请求
        2. 系统调度为该进程分配cpu
        3. cpu向DMA发送io请求
        4. 用户进程等待io完成，让出cpu
        5. 系统调度cpu执行其他任务
        6. 数据写入至io控制器的缓冲寄存器
        7. DMA不断获取缓冲寄存器中的数据（需要cpu时钟）
        8. 传输至内存（需要cpu时钟）
        9. 所需的全部数据获取完毕后向cpu发出中断信号
    
   ![avatar](../img/20180429234354258.png) 
    
- 使用sendfile + DMA代替CPU在内核态拷贝数据
   ![avatar](../img/3201622308364_.pic.jpg)
    
- 题外话
    1. sendfile只适用于将数据从文件拷贝到套接字上，限定了它的使用范围。Linux在2.6.17版本引入splice系统调用，用于在两个文件描述符中移动数据：
    ```C
    #define _GNU_SOURCE         /* See feature_test_macros(7) */
    #include <fcntl.h>
    ssize_t splice(int fd_in, loff_t *off_in, int fd_out, loff_t *off_out, size_t len, unsigned int flags);
    ```
    2. splice调用在两个文件描述符之间移动数据，而不需要数据在内核空间和用户空间来回拷贝。他从fd_in拷贝len长度的数据到fd_out，但是有一方必须是管道设备，这也是目前splice的一些局限性。
    3. 有些时候有些数据非要从内核态拷贝到用户态，Linux 还提供了写时复制技术(copy on write)
     

