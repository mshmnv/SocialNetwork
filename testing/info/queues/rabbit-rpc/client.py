#!/usr/bin/env python
import uuid

import pika
import time

# Клиент для вызова через RPC в RabbitMQ
# Вычислим факториал на удалённом сервере
class FactorialRpcClient(object):

    def __init__(self):
        # установка подключения к брокеру
        self.connection = pika.BlockingConnection(
            pika.URLParameters('amqp://guest:guest@localhost'))

        # создание канала
        self.channel = self.connection.channel()

        # очередь для получения результатов
        # название пустое -> выбирается случайно
        # очередь exclusive - привязана только к этому процессу
        result = self.channel.queue_declare(queue='', exclusive=True)
        # сохраняем название очереди
        self.callback_queue = result.method.queue

        # и начинаем прослушивание ответов на RPC-вызовы
        self.channel.basic_consume(
            queue=self.callback_queue,
            on_message_callback=self.on_response,
            auto_ack=True)

    def on_response(self, ch, method, props, body):
        # при получении ответа от RPC проверяем совпадение correlation_id
        if self.corr_id == props.correlation_id:
            # если совпали - сохраняем результат вызова функции
            self.response = body

    def call(self, n):
        self.response = None
        # создаём случайный correlation_id
        self.corr_id = str(uuid.uuid4())
        # отправляем запрос на вызов функции с указанием свойств:
        # reply_to - название очереди для отправки результата
        # correlation_id - уникальный идентификатор запроса
        self.channel.basic_publish(
            exchange='',
            routing_key='rpc_queue',
            properties=pika.BasicProperties(
                reply_to=self.callback_queue,
                correlation_id=self.corr_id,
            ),
            body=str(n))
        # синхронно ожидаем получения результата
        while self.response is None:
            self.connection.process_data_events()
        # возвращаем целочисленный результат вычисления
        # (в этом случае получен строкой в теле сообщения,
        # но может быть и более сложная двоичная сериализация
        # объектов)
        return int(self.response)


# запрос удаленной процедуры через RPC
factorial_rpc = FactorialRpcClient()

for i in range(1, 21):
    print(f" [x] Requesting factorial({i})")
    response = factorial_rpc.call(i)
    print(" [.] Got %r" % response)
    time.sleep(1)
