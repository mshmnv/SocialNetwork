#!/usr/bin/env python
import pika

# создать подключение
url = "amqp://guest:guest@localhost"
connection = pika.BlockingConnection(
    pika.URLParameters(url))

# создать канал
channel = connection.channel()

# очередь для получения запросов удаленного вызова
# процедур (rpc_queue)
channel.queue_declare(queue='rpc_queue')


# рекуррентное вычисление факториала
def factorial(n):
    if n <= 1:
        return 1
    else:
        return n * factorial(n - 1)


# при получении сообщения
def on_request(ch, method, props, body):
    # получить аргумент (передан строкой в теле сообщения)
    n = int(body)

    print(" [.] fact(%s)" % n)
    # выполнить вычисление
    response = factorial(n)

    # отправить ответ в очередь, указанную в свойствах сообщения
    # props.reply_to (см. client)
    # для связи запроса и ответа сохраняется свойство
    # correlation_id
    ch.basic_publish(exchange='',
                     routing_key=props.reply_to,
                     properties=pika.BasicProperties(correlation_id= \
                                                         props.correlation_id),
                     body=str(response))
    # подтверждение успешной обработки
    ch.basic_ack(delivery_tag=method.delivery_tag)


# забираем по одному сообщению, чтобы другие consumer'ы могли
# выполнить другие задачи более равномерно
channel.basic_qos(prefetch_count=1)
# подписываемся на очередь
channel.basic_consume(queue='rpc_queue', on_message_callback=on_request, auto_ack=False)

print(" [x] Awaiting RPC requests")
channel.start_consuming()
