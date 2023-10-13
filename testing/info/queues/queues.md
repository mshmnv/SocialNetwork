## Rabbit


1. docker run -itd -p 5672:5672 -p 15672:15672 --name rabbit rabbitmq:3-management

2. docker exec -it rabbit bash

`rabbit-plugins list`

3. docker network create cluster-network

### Две ноды в один кластер

4. docker run -d --hostname node1.rabbit --net cluster-network --name rabbitNode1 --add-host node2.rabbit:172.24.0.3 -p  "15673:15672" -e "RABBITMQ_USE_LONGNAME=true" -e RABBITMQ_ERLANG_COOKIE="cookie" rabbitmq:3-management

5. docker run -d --hostname node2.rabbit --net cluster-network --name rabbitNode2 --add-host node1.rabbit:172.24.0.2 -p  "15674:15672" -e "RABBITMQ_USE_LONGNAME=true" -e RABBITMQ_ERLANG_COOKIE="cookie" rabbitmq:3-management

6. docker exec -it rabbitNode1 bash

7. rabbitmqctl stop_app

8. rabbitmqctl join_cluster rabbit@node2.rabbit

9. rabbitmqctl start_app

## Kafka

10. cd kafka-test

11. docker-compose up -d

12. docker exec -it kafka-test-kafka-1 bash

13. kafka-topics --bootstrap-server localhost:9092 --topic test --create

14. kafka-topics --bootstrap-server localhost:9092 --list

15. kafka-console-consumer --bootstrap-server localhost:9092 --topic test

16. kafka-console-producer --bootstrap-server localhost:9092 --topic test
