#ID生成器

###订单ID
1. 算法unix 秒级时间戳 << 14 | redis->incr(unix时间戳) << 13 | mt_rand(0, 1)
2. redis 出问题时，用mt_rand(0, 2^13)随机数替代redis, 但并发肯定没法保证。

###箱子ID
1. 订单ID + 序列号（订单上第几箱）