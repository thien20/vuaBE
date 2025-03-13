import os
import sys
import json
import logging
from kafka import KafkaConsumer
from scrape import *
from utils.db_inserting import save_to_database, update_job_status
from utils.file_utils import write_json
from utils.data_indexing import indexing_only_text_with_bulk
from ingest import ingest

CONFIG_PATH = "/app/config/config.json"
with open(CONFIG_PATH, "r") as config_file:
    config = json.load(config_file)

KAFKA_BROKER = config.get("kafka", "kafka:29092")  # Default to 'kafka:29092' if missing

logging.basicConfig(level=logging.DEBUG)
logging.getLogger('kafka').setLevel(logging.INFO)

def consume_kafka_messages():
    logging.debug(f"Kafka broker: {KAFKA_BROKER}")

    consumer = KafkaConsumer(
        'scrape-news',
        bootstrap_servers=[KAFKA_BROKER],
        auto_offset_reset='latest',
        api_version=(3, 8, 1),
        group_id='scrape-group',
        value_deserializer=lambda x: json.loads(x.decode('utf-8')) if x else None
    )
    
    logging.debug("Kafka Consumer is now listening...")

    while True:
        message_pack = consumer.poll(timeout_ms=1000)
        no_message_count = 0

        if not message_pack:
            no_message_count += 1
            if no_message_count % 10 == 0:
                logging.debug("No messages received in the last poll.")
            continue
        no_message_count = 0

        for tp, messages in message_pack.items():
            for message in messages:
                if message.value is None:
                    logging.warning("Skipping empty message.")
                    continue
                
                logging.debug(f"Raw message received: {message.value}")

                try:
                    action = message.value.get('action')
                    job_id = int(message.value.get('job_id'))

                    print(f"Action: {action}, Job ID: {job_id}")

                    if action == 'scrape':
                        data = scrape_all_categories(job_id=job_id)
                        indexing_only_text_with_bulk(data, 'news')
                        # write_json(data, 'temp.json')
                        logging.info("Data scraped and indexed successfully.")
                        save_to_database(job_id, data)
                        update_job_status(job_id, 'completed')

                    elif action == 'ingest':
                        ingest()

                except Exception as e:
                    logging.error(f"Error processing message: {e}", exc_info=True)

# cd app --> `dlv debug --check-go-version=false main.go
# source simplebe/bin/activate --> python3 -m debugpy --listen 5678  /mnt/d/MY_FOLDER/Project/vuaBE/pipeline/consumer.py
# curl -X POST curl http://0.0.0.0:8080/jobs/fetch/scrape
if __name__ == '__main__':
    consume_kafka_messages()
