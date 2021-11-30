#Redis 多线程
###参考链接
    https://mp.weixin.qq.com/s/TAPkqk3y3HqNyDL6w7b8Bw
    https://www.baidu.com/s?ie=UTF-8&wd=Lazy%20Free%20%E6%9C%BA%E5%88%B6

####单线程问题点:
    1. 受限于 IO
    2. 受限于大 Key 的删除，会阻塞线程

####Redis 多线程：
    1. 利用多线程来分担 IO 读写的负负荷
    2. 事件处理线程每次获取到可读的事件时，就会将就绪事件分配给 IO 线程，让 IO 线程进行处理，并将命令写到等待队列中。
    3. IO线程处理完后，事件线程就会处理等待队列中的任务。事件线程处理完任务之后，会将任务分配给 IO 线程进行写操作。
    4. IO线程 轮询完读事件之后，才会轮询写事件。读写不会同时进行。
    
####Lazy Free机制
    1. 解决大Key删除导致线程卡住的问题。
    2. 根据配置开启 lazy free
        1. lazyfree-lazy-eviction 
            当 redis 内存达到 maxmeory，并设置有淘汰策略时; 在被动淘汰过期键时，是否开启 lazy free。可能会导致内存超限，
            因为 Unlink 并不是立刻释放内存。
        2. lazyfree-lazy-expire 
            针对设有 TTL 的键，达到过期时间时，被 redis 清理时是否开启 lazy free。
        3. lazyfree-lazy-server-del 
            针对有些指令在处理已存在的键时，会带有一个隐式的DEL键的操作。如rename命令，当目标键已存在,redis会先删除目标键，如果这些目标键是一个big key,
            那就会引入阻塞删除的性能问题。
        4. slave-lazy-flush
            针对slave全量同步master数据时，在加载master的RDB文件时，会运行flushall来清理自己的数据场景。
    3. 比如删除时 （DEL），会看 lazy free 配置是否开启，如果开启就会进入 dbAsyncDelete(c->db,c->argv[j]), dbAsyncDelete 依旧会通过计算删除键的回收收益，
    来决定是否用 background IO进行异步删除，比如：String 类型回收收益就是1，而 Set 类型，回收收益就是集合中元素个数。
    4. 超过一定的回收收益(64)，才会封装成 JOB 加入到异步处理队列中，否则还是会同步回收内存。
    5. 计算回收收益的函数 lazyfreeGetFreeEffort, 主要是计算时间复杂度
    ```C
    size_t lazyfreeGetFreeEffort(robj *obj) {
       if (obj->type == OBJ_LIST) {  
           quicklist *ql = obj->ptr;
           return ql->len;
       } else if (obj->type == OBJ_SET && obj->encoding == OBJ_ENCODING_HT) {
           dict *ht = obj->ptr;
           return dictSize(ht);
       } else if (obj->type == OBJ_ZSET && obj->encoding == OBJ_ENCODING_SKIPLIST){
           zset *zs = obj->ptr;
           return zs->zsl->length;
       } else if (obj->type == OBJ_HASH && obj->encoding == OBJ_ENCODING_HT) {
           dict *ht = obj->ptr;
           return dictSize(ht);
       } else {
           return 1; /* Everything else is a single allocation. */
       }
    }
    ```
    6. dbAsyncDelete 函数
    ```C
    #define LAZYFREE_THRESHOLD 64 //根据FREE一个key的cost是否大于64，用于判断是否进行lazy free调用
    int dbAsyncDelete(redisDb *db, robj *key) {
    /* Deleting an entry from the expires dict will not free the sds of
     * the key, because it is shared with the main dictionary. */
    if (dictSize(db->expires) > 0) dictDelete(db->expires,key->ptr); //从expires中直接删除key
 
    dictEntry *de = dictUnlink(db->dict,key->ptr); //进行unlink处理，但不进行实际free操作
    if (de) {
        robj *val = dictGetVal(de);
        size_t free_effort = lazyfreeGetFreeEffort(val); //评估free当前key的代价
 
        /* If releasing the object is too much work, let's put it into the
         * lazy free list. */
        if (free_effort > LAZYFREE_THRESHOLD) { //如果free当前key cost>64, 则把它放在lazy free的list, 使用bio子线程进行实际free操作，不通过主线程运行
            atomicIncr(lazyfree_objects,1); //待处理的lazyfree对象个数加1，通过info命令可查看
            bioCreateBackgroundJob(BIO_LAZY_FREE,val,NULL,NULL); 
            dictSetVal(db->dict,de,NULL);
        }
    }
    ```

####锁过期
    1.主线程同步删除
    2.lazyFree
    3.定时抽样删除
        首先通过配置或者默认值计算出几个参数，这几个参数直接或间接决定了这些执行的终止条件，分别如下：
            config_keys_per_loop: 每次循环抽样的数据量
            config_cycle_fast_duration: 快速清理模式下每次清理的持续时间
            config_cycle_slow_time_perc: 慢速清理模式下每次清理最大消耗CPU周期数(cpu最大使用率)
            config_cycle_acceptable_stale: 可接受的过期数据量占比，如果本次采样中过期数量小于这个阈值就结束本次清理。
            根据上述参数计算出终止条件的具体值（最大采样数量和超时限制）。
            遍历清理所有的db。
            针对每个db在终止条件的限制下循环清理。