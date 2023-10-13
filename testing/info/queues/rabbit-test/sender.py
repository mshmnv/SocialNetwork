import logging
import pika


logging.basicConfig()
url = "amqp://guest:guest@localhost/"
params = pika.URLParameters(url)
params.socket_timeout = 5
connection = pika.BlockingConnection(params)
channel = connection.channel()

exchange = "test"
queue = "testqueue"

# channel.tx_select()
channel.exchange_declare(exchange, durable=True)
for i in range(0, 100):
    channel.basic_publish(exchange, queue, str(i).encode("utf-8"), mandatory=True)
# channel.tx_rollback()
connection.close()
