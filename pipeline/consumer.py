import os
import sys
import json

# sys.path.append(os.path.abspath("pipeline/scrape.py"))

from scrape import *
from kafka import KafkaConsumer
from utils.db_inserting import save_to_database, update_job_status
# from cache.cache import save_to_cache


def consume_kafka_messages():
    consumer = KafkaConsumer(
        'scrape-news',
        bootstrap_servers=['localhost:9092'],
        group_id='scrape-group',
        value_deserializer=lambda x: json.loads(x.decode('utf-8')) # convert bytes to dict
    )
    temp_data = {
        "job_id": "1",
        "category": "Politics",
        "link": "https://vnexpress.net/chinh-tri",
        "title": "Chính trị",
        "content": "Chính trị",
        "scraped_at": "2021-07-20 12:00:00"
    }

    for message in consumer:
        action = message.value.get('action')
        job_id = message.value.get('jobID')
        print(action, job_id)
        print(f"Consumed message: {message.value}")
        if action == 'scrape':
            # Trigger the scraping process
            # data = scrape_all_categories()
            data = temp_data
            with open('temp.json', 'w', encoding='utf-8') as f:
                json.dump(data, f, ensure_ascii=False, indent=2)
            save_to_database(job_id, data)
            # job_id = int(job_id)
            # save_to_cache(job_id, data)
            update_job_status(job_id, 'completed')
                
        # elif action == 'ingest':
        #     ingest()
        
        
if __name__ == '__main__':
    consume_kafka_messages()