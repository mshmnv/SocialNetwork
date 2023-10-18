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
5. Sharding [x]
```
    + /dialog/{user_id}/send - send message to user
    + /dialog/list - get message history between users
    
    Horizontal scaling (scaling out) with sharding.
    Resharding possibility.
    
    RESHARDING
    1 - two shard mappings
    2 - write - new shard map; read - both new and old shard map
```
6. Queues
```

```
7. In-memory db 
8. Microservices 
9. Balancing and Fault tolerance 
10. Distributed Transactions