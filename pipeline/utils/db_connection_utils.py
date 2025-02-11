import mysql.connector as connection
import mysql.connector

def connect_to_db(db_config):
    
    # Connect to the database
    try:
        conn = connection.MySQLConnection(**db_config)
        cursor = conn.cursor()
    except mysql.connector.Error as e:
        print(f"Failed to connect to MySQL: {e}")
        exit(1)
    return conn, cursor