from json import loads

from kafka import KafkaConsumer

consumer = KafkaConsumer(
    #название topic для прослушивания
    'test-topic',
    #нужно использовать в точности имя хоста и порт, как
    #в advertised_listeners (может быть либо kafka:9092,
    #(нужно добавить в hosts kafka 127.0.0.1),
    #либо localhost:19092 (тогда в docker-compose.yaml
    #нужно выполнить публикацию на внешний порт 19092,
    #порты у разных listeners должны отличаться)
    bootstrap_servers=['localhost:19092'],
    #сброс смещения на самий раннее из доступных сообщений
    #(если ранее курсор не существовал)
    auto_offset_reset='earliest',
    #автоматическое подтверждение получения
    enable_auto_commit=True,
    #идентификатор группы consumer
    group_id='my-group',
    #метод десериализации значения (может быть также
    #задан метод для десериализации ключа)
    value_deserializer=lambda x: loads(x.decode('utf-8')))

for message in consumer:
    message = message.value
    print(message)
