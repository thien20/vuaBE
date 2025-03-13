import os
from scrape import scrape_all_categories
from utils.file_utils import write_json

path = 'https://vnexpress.net'

data = scrape_all_categories()

current_dir = os.path.dirname(os.path.abspath(__file__))
data_dir = os.path.join(current_dir, 'data')
os.makedirs(data_dir, exist_ok=True)  # Create the data directory if it doesn't exist
output_path = os.path.join(data_dir, 'all_data.json')
write_json(data, output_path)