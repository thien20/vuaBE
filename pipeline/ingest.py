import os
import json
import mysql.connector

import config

from mysql.connector import connection


db_config = config.Config("config.json").db_config
def ingest():

    # Connect to the database
    try:
        conn = connection.MySQLConnection(**db_config)
        cursor = conn.cursor()
        
    except mysql.connector.Error as e:
        print(f"Failed to connect to MySQL: {e}")
        exit(1)

    # Data files
    data_dir = r'./pipeline/data'

    table = 'news'

    with open(os.path.join(data_dir, 'all_data.json'), 'r', encoding='utf-8') as f:
        data = json.load(f)
        records_to_insert = []
        for record in data:
            record_id = record.get('id')
            link = record.get('link', '')
            title = record.get('title', '')
            content = record.get('content', '')
            category = record.get('category', '')
            date = record.get('date', '')
            if record_id is None or not link or not title or not content:
                print(f"Skipping invalid record: {record}")
                continue
            records_to_insert.append((record_id, link, title, content, category, date))

        try:
            cursor.executemany(f"""
            INSERT INTO `{table}` (id, link, title, content, category, date)
            VALUES (%s, %s, %s, %s, %s, %s)
            """, records_to_insert)
        except mysql.connector.Error as e:
            print(f"Error inserting records into table `{table}`: {e}")

    # Commit and close connection
    conn.commit()
    cursor.close()
    conn.close()