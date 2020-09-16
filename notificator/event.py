import sys
import json
import os
from gnsq import Producer

event = sys.argv[1]
hash = sys.argv[2]
message = json.dumps({'event': event, 'hash': hash,'full': sys.argv}).encode()

producer = Producer(os.getenv('NSQ_DSN', 'localhost:4150'))
producer.start()
result = producer.publish(os.getenv('NSQ_TOPIC', 'cryptogate-btc-blockchain-events'), message)