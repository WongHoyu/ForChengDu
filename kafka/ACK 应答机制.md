# ACK 应答机制

1. Kafka 提供三种可靠性级别
    - ack 参数:
        1. 0: producer 不等待 broker 的ack。
        2. 1: producer 等待 broker 的 ack，当 partition 的 leader 将数据落盘时，即可返回 ACK。
        3. -1: producer 等待 broker 的 ack，当 partition 的 leader 和 follower(ISR 中的 follower) 全部将数据落盘后才返回 ACK。
    - 分别对应的故障:
        1. 0: 对 broker 中的 partition 是否将落盘无感知，broker 故障就会导致数据丢失。
        2. 1: 当 leader 故障时，follower 可能还未同步完数据时，重新从 ISR 中选举 leader，可能会导致数据丢失。
        3. -1: 当 broker 在返回 ACK 前，leader 发生故障，producer 重复发送消息，会导致消息重复。
    - 故障细节：
        1. follower 故障:
            - follower 故障会暂时被踢出ISR，待恢复正常时，follower 会读取上一次的 HW，然后将大于 HW 的日志截取，
            重新从 leader 中的 HW 开始同步消息，当 follower 的 LEO 大于等于 leader 的 HW，将会重新进去 ISR。
        2. leader 故障：
            - 当 leader 故障时，会重新从 ISR 中选取新 leader，然后其余 follower 会将高于 HW 的日志截取，重新从新的 leader 的 HW 开始同步数据，
            此时只能保证副本之间数据的一致性，并不能保证数据的不丢失或不重复。
            
    