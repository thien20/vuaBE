import os
import json
import mysql.connector

from config import *
from mysql.connector import connection

db_config = {
    'user': DB_USER,
    'password': DB_PASSWORD,
    'host': DB_HOST,
    'database': DB_NAME
}

# Connect to the database
try:
    # print("hello 1")
    conn = connection.MySQLConnection(**db_config)
    cursor = conn.cursor()
    
    # conn = mysql.connector.connect(**db_config)
    # cursor = conn.cursor()
except mysql.connector.Error as e:
    print(f"Failed to connect to MySQL: {e}")
    exit(1)

# Data files
data_dir = r'D:\MY_FOLDER\Project\vuaBE\data'
json_files = [f for f in os.listdir(data_dir) if os.path.isfile(os.path.join(data_dir, f))]

if not json_files:
    print(f"No JSON files found in the directory: {data_dir}")
    exit(1)

tables = [os.path.basename(file).replace('data_', '').replace('.json', '').replace('-', '_') for file in json_files]

json_files.sort()
tables.sort()
for file, table in zip(json_files, tables):
    with open(os.path.join(data_dir, file), 'r', encoding='utf-8') as f:
        data = json.load(f)
        records_to_insert = []
        for record in data:
            record_id = record.get('id')
            link = record.get('link', '')
            title = record.get('title', '')
            content = record.get('content', '')
            if record_id is None or not link or not title or not content:
                print(f"Skipping invalid record: {record}")
                continue
            records_to_insert.append((record_id, link, title, content))

        try:
            cursor.executemany(f"""
            INSERT INTO `{table}` (id, link, title, content)
            VALUES (%s, %s, %s, %s)
            """, records_to_insert)
        except mysql.connector.Error as e:
            print(f"Error inserting records into table `{table}`: {e}")

# Commit and close connection
conn.commit()
cursor.close()
conn.close()