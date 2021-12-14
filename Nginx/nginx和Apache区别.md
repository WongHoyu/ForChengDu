#Apache和Nginx区别

##参考链接
https://cloud.tencent.com/developer/article/1635326

1. Apache功能丰富、模块多、稳定、适合用于动态资源请求；Nginx有反向代理，负载均衡功能，比较适合做静态资源请求。
2. Apache是select模型；Nginx用的epoll模型。