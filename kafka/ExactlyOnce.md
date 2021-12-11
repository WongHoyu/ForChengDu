#ExactlyOnce

##参考链接
https://www.zybuluo.com/tinadu/note/949867

1. 通过 transactionId 找到 broker 端持久化的 producerID + epoch（版本号）。
2. epoch 版本号低的生产者发送的消息，broker会拒绝掉。相同 epoch 而版本号低的生产者将会不再工作。
3. broker 会维护一个 <pid, topic, partition> 的序列号，一个pid发送的 topic+partition的消息，序列号是持续+1的。如果该序列号比上一个序列号 > 2,
    broker 会认为乱序，前一个消息未被提交，而拒绝本次消息的提交。

4. 个人认为，还是在消费者那边根据业务来做幂等更靠谱，宁愿生产者发送了多条相同的消息。