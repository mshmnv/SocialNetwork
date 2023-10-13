from json import dumps
from time import sleep

from kafka import KafkaProducer

#важно не забыть создать topic (через UI или
#через kafka-topics, например внутри контейнера:
#kafka-topics --create --topic test-topic --bootstrap-server kafka:9092

producer = KafkaProducer(
    # нужно использовать в точности имя хоста и порт, как
    # в advertised_listeners (может быть либо kafka:9092,
    # (нужно добавить в hosts kafka 127.0.0.1),
    # либо localhost:19092 (тогда в docker-compose.yaml
    # нужно выполнить публикацию на внешний порт 19092,
    # порты у разных listeners должны отличаться)
    bootstrap_servers="localhost:19092",
    # метод сериализации значения (может быть также
    # задан метод для сериализации ключа)
    value_serializer=lambda x: dumps(x).encode('utf-8'))

for e in range(100):
    data = {'number': e}
    producer.send('test-topic', value=data)
    sleep(1)
producer.flush()
