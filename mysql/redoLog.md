#redoLog 重做日志

##参考链接
https://blog.51cto.com/u_15015155/2615464
https://www.cnblogs.com/f-ck-need-u/p/9010872.html

1. redoLog 以块存储在磁盘上
2. redoLog 写在redo log group 上，redo log group 至少有两个redo log file，循环写在这个 redo log file上。
3. redo log file 头节点占 2kb，其中记录着 checkpoint
4. innodb_flush_log_at_trx_commit 0 先写到用户空间缓冲区，再由后台线程每秒刷脏日志到磁盘; 1 commit 直接写到磁盘; 2 commit 写到内核空间缓冲区，
    再由后台线程刷脏日志到磁盘。
5. redoLog 两阶段提交：事务commit时，先将 redoLog 写入到缓冲区，并标记为 prepare；将binLog写入缓冲区；最后将prepare标为commit。 
    何时刷盘由 innodb_flush_log_at_trx_commit 决定。binlog同样道理，由 sync_binlog 决定。
6. checkpoint 是最后一次刷脏数据，脏日志到磁盘的标志点，checkpoint 的 lsn 代表这个 lsn 及其之前的数据都已经刷盘成功，不需要重做。
7. 崩溃恢复，比对数据页上的lsn和redoLog上的lsn，如果前者比后者少，则直接恢复checkpoint之后的数据即可。
8. lsn 存在于缓存数据页、redoLog缓冲区、磁盘数据页、磁盘redoLog file、redoLog file的checkpoint上。当 redoLog 刷盘、checkPoint 刷盘等都会更新lsn