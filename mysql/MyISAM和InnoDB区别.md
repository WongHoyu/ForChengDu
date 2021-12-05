#MyISAM 和 InnoDB区别

##参考链接
https://www.runoob.com/w3cnote/mysql-different-nnodb-myisam.html

1. InnoDB 支持事务，MyISAM不支持事务
2. InnoDB 支持行锁，MyISAM只支持表级锁
3. InnoDB 不支持全文索引，MyISAM支持全文索引
4. InnoDB count(*) 只能依据索引来查找总行数，MyISAM 会存储总行数，直接返回。
5. 构成文件上也有不同。