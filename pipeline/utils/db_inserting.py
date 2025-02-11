import os
import sys

sys.path.append(os.path.abspath("pipeline/config.py"))

from kafka import KafkaConsumer
from config import Config
from utils.db_connection_utils import connect_to_db


data_path = r"../vuaBE/app/config/config.json"
db_config = Config(data_path).db_config

def save_to_database(job_id, scraped_data):
    """Save scraped data to MySQL database."""

    table = "scraped_results"
    connection, cursor = connect_to_db(db_config)

    try:
        with connection.cursor() as cursor:
            records = []
            scraped_data = [scraped_data]
            # print(type(scraped_data))
            for data in scraped_data:
                job_id = int(data.get('job_id'))
                category = data.get('category')
                link = data.get('link')
                title = data.get('title')
                content = data.get('content')
                scraped_at = data.get('scraped_at')
                records.append((job_id, category, link, title, content, scraped_at))
            # cursor.executemany(f"""SHOW COLUMNS FROM {table}""")
            cursor.executemany(f"""
            INSERT INTO {table} (job_id, category, link, title, content, scraped_at)
            VALUES (%s, %s, %s, %s, %s, %s)
            """, records)
        connection.commit()
        print(f"Successfully saved {len(scraped_data)} records to the database for job {job_id}")
    except Exception as e:
        print(f"Error saving data: {e}")
    finally:
        connection.close()
        
def update_job_status(job_id, status):
    """Update job status in the database."""

    connection, cursor = connect_to_db(db_config)
    table = "jobs"
    try:
        with connection.cursor() as cursor:
            cursor.executemany(f"""
            UPDATE {table} SET status = %s WHERE id = %s
            """, [(status, job_id)])
            # sql = f"UPDATE {table} SET status = %s WHERE id = %s"
            # cursor.execute(sql, (status, job_id))
        connection.commit()
        print(f"Successfully updated job {job_id} status to {status}")
    except Exception as e:
        print(f"Error updating job status: {e}")
    finally:
        connection.close()
