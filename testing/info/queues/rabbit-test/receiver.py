import pika, logging
from pika import spec
from pika.adapters.blocking_connection import BlockingChannel


def on_message(
        channel: BlockingChannel,
        method: spec.Basic.Deliver,
        properties: spec.BasicProperties,
        body: bytes):
    # print(channel)
    # print(method)
    # print(properties)
    print(body.decode('utf-8'))
    # если auto_ack=False (ниже), нужно вернуть положительное
    # (ack) или отрицательное подтверждение (nack)
    # во втором случае сообщение отправляется обратно в очередь
    # и будет возвращено с redelivered=True
    # channel.basic_nack(delivery_tag=method.delivery_tag)
    # channel.basic_ack(delivery_tag=method.delivery_tag)


logging.basicConfig()
url = "amqp://guest:guest@localhost/"
params = pika.URLParameters(url)
params.socket_timeout = 5
connection = pika.BlockingConnection(params)
channel = connection.channel()
#установить ограничение на 5 сообщений единовременно
channel.basic_qos(prefetch_count=5)
# channel.tx_select()

exchange = "test"
queue = "testqueue"

channel.queue_declare(queue, durable=True)
channel.queue_bind(queue, exchange)
#здесь можно включить auto_ack для автоматического подтверждения
channel.basic_consume(queue, on_message_callback=on_message, auto_ack=True)
channel.start_consuming()
