# SocialNetwork

Запуск
```bash
  make up
  make migration-up
```

## Postman Collenction
[postman_collection.json](https://github.com/mshmnv/SocialNetwork/blob/main/info/testing/postman_collection.json)


---

### *Highload Architect Course:*
1. MVP
```

```
2. Indexes and Load Testing
```

```
3. Replication
```

```
4. Cache
```

```
5. Sharding
```
    + /dialog/{user_id}/send - send message to user
    + /dialog/list - get message history between users
    
    Horizontal scaling (scaling out) with sharding.
    Resharding possibility.
    
    RESHARDING
    1 - two shard mappings: old and new
    2 - Write - new shard map; Read - both new and old shard map (prioritize new if the same record)
    3 - Copy records from wrong shard num to the new one and delete them on the wrong shard after.
```
6. Queues
 ```
   /post/feed/posted/ - online updated feed using websocket connections.
   Implemetation:
   RabbitMQ is used as a queue between newly created posts and open websocket connections.
   Once new post created, producer sends all new posts to user's friends queues if they have opened connections.
   Consumer gets posts from its own queue and writes to the opened websocket connection.
```

7. In-memory db. Tarantool
8. Microservices 
9. Balancing and Fault tolerance 
10. Distributed Transactions
11. Monitoring and alerting
12. 