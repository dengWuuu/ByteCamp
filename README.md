## 抖音项目服务端简单示例

具体功能内容参考飞书说明文档
https://bytedancecampus1.feishu.cn/docx/ZrYxdIIYRonQe4xTD45caEvznXd



## TODO
1. redis mq之间更新
2. mq 防止消费者重复消费
3. mq 存入数据库防止消息丢失
4. 启动ack机制
5. 各种参数校验
6. redis分片实现videoId为key存的userId set， 解决大key问题