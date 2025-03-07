import os
import json

from utils.file_utils import ensure_dir, read_json, write_json
from utils.scrape_utils import scrape_data_by_cate

def scrape_all_categories():
    original_url = 'https://vnexpress.net'

    # Define relative paths
    script_dir = os.path.dirname(os.path.abspath(__file__))
    # print("This is current dir", script_dir)
    data_dir = os.path.join(script_dir, 'data')
    # print("This is data dir", data_dir)
    json_path = r'../vuaBE/vuacrawl.json'

    # Ensure data directory exists
    ensure_dir(data_dir)

    # Load categories
    categories = read_json(json_path)
    categories = [item['share_url'] for item in categories]

    # Remove video category and truncate "/"
    categories.remove('https://video.vnexpress.net')
    url_list = [original_url + cate for cate in categories]    
    categories = [cate.split('/')[-1] for cate in categories] 

    all_data = []
    for temp_url in url_list:
        data = scrape_data_by_cate(temp_url)
        all_data.extend(data)
        print("Done scraping data for category:", temp_url)
    
    # Save all scraped data
    write_json(all_data, os.path.join(data_dir, 'all_data.json'))
    print("Done scraping and writing data for all categories")

    return all_data


# if __name__ == '__main__':
#     scrape_all_categories()
    # consume_kafka_messages()
