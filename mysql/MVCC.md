#MVCC

##参考链接
https://www.cnblogs.com/jelly12345/p/14889331.html

1. mvcc 就是多版本并发控制，利用readView + undoLog 解决读写冲突。
2. readView 就是快照读时，对系统的数据生成快照，以及存储当前数据库活跃的事务ID。
3. RC级别下，每一次快照读都会生成一个readView，RR 只会在第一次快照读生成 readView。
4. 每一行数据都会有四个隐藏列，自增ID rowID，修改这行数据的事务ID DB_TRX_ID，delete_flg 以及回滚指针 DB_ROLL_PTR。
5. 事务在快照读时，生成了readView，就会根据readView对该行数据做可见性判断：
   1. 如果该行数据的 DB_TRX_ID 小于 readView 中活跃事务list的所有事务ID，则认为该数据在当前事务生成readView时已提交，当前事务可读。
   2. 否则，判断 DB_TRX_ID 是否大于 readView 中活跃事务list的所有事务ID，如果是，则该行数据是生成readView后才有的，则不可读。
   3. 否则，判断 DB_TRX_ID 是否在活跃事务中，如果不在，则证明事务已提交，可读。
   4. 否则，不可读。则根据回滚指针，读下一条undoLog，继续可读性判断，找到可读数据为止。