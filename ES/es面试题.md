# ES 面试题

### 倒排索引
  - 由词典(Term Dictionary)，倒排列表(Posting List) 构成
  - 词典: 是存储所有不重复的词条，每个词条对应唯一的 termID
  - 倒排列表: 每个词条都有一个倒排列表，记录了哪些文档包含了该单词。 
      倒排列表由多个 PostingEntry组成，每个 PostingEntry 包含文档 ID 和该单词在文档中出现的位置等信息。
  - 当进行搜索时，ES 会首先根据用户查询从词典种查找对应的单词编号，然后根据编号找到对应的倒排列表，然后获取包含该单词的所有文档列表。
      最后结合 Posting Index中的文档信息进行评分，排序，最终返回匹配的结果
  - FST 在倒排索引中的应用：
    - 词典的FST实现： ES 使用 FST压缩存储分析出的所有唯一词条，借助字符串前缀的公共重叠部分，极大地减少了存储空间，每个终止节点存储了词条的 termID。
    - 倒排列表的 FST 实现：对于每个词条的倒排列表，ES 使用 FST 高效压缩编码文档 ID 和位置信息
    - 查询过程：ES 先在字典的 FST 中找到对应词条的 termID，再根据 termID 找到压缩存储的倒排列表的 FST，遍历解压缩获取文档 ID 列表，最终根据文档 ID
获取存储的文档数据，完成查询。


### Master 选举
- 参考文章：https://blog.csdn.net/qq_35373760/article/details/108974308
1. master 由谁选举，什么时候发起
    - master 选举由 master-eligible（master 备选节点）发起，当一个 master-eligible 节点发现满足以下条件时会发起选举：
      - 该 master-eligible 节点的当前状态不是 master
      - 该 master-eligible 节点通过 ZenDiscovery 模块的 ping 操作询问其已知的集群其他节点，没有任何节点连接到master
      - 包括本节点在内，当前已有超过 minimum_master_nodes 个节点没有链接到 master
      - 总结：当前一个节点发现包括自己在内的多数派的 master-eligible 节点认为集群没有 master 时，就可以发起 master 选举
2. 当需要选举master时，选举谁？
   - 根据节点的 clusterStateVersion排序，值越大优先级越高。clusterStateVersion 相同时，根据节点的 ID 比较（ID 为节点第一次启动时随机生成的）
   - 排序完后，拿到第一个节点，并选举其为 master
3. 什么时候选举成功？
   - 假设 NodeA 选举 NodeB 当 master，NodeA 会向NodeB 发送 join 请求：
     - 如果 NodeB 已经成为 master，NodeB 就会把 NodeA 加入到集群中，然后发布最新的 cluster_state，最新的 cluster_state就会包含 NodeA信息。
        对于 NodeA，等新的 cluster_state 发布到 NodeA 的时候，就会完成 join。
     - 如果 NodeB 在精选 master，那么 NodeB 会把这次 join 当做一张选票。对于这种情况，Node 会等待一段时间，看 NodeB 是否能成为真正的 Master，
        直到超时或者有别的 master 选成功。
     - 如果 NodeB 认为自己不是 master（现在不是，将来也选不上），那么 NodeB 会拒绝这次 join。而 NodeA 会开启下一轮选举。
   - 假设 NodeA 选自己当 Master：
     - 此时 NodeA 会等待别的 Node 的 join 请求（选票），当收集到过半数的选票时，认为自己成为 master，然后更新 cluster_state种的 master node 为自己，
         并向集群发布这一消息


### Elasticsearch 搜索的过程
1. 查询解析
    - 客户端发出查询请求后，节点先将查询字符串解析成对应的查询对象（Query DSL）形式，确定查询的上下文信息，比如查询类型，需要用到的索引
2. 查询重写
    - 查询解析后，ES 会将查询进行规范化，优化及重写处理，比如去除冗余，合并多个相同的字查询等
3. 路由计算
    - ES 需要计算出查询需要走向那些分片，根据文档 ID 哈希计算目标分片列表。对于主分片失效的情况，还需要考虑查询副本分片。
4. 查询执行
    - 确定目标分片后，ES 会在每个涉及的数据节点的每个分片副本上并行执行查询，将查询结果写入位集（bitSets）中，表明满足查询条件的文档。
5. 分片查询结果汇总
    - 所有分片的查询结果会被收集到协调节点，并在协调节点上进行合并去重，排序等操作。
6. 格式化查询结果