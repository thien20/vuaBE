from bs4 import BeautifulSoup

import requests
import json
import os

global_id = 0 

def process_p_contents(p_contents):
    return " ".join(p_contents)

def scrape_data_by_cate(url):
    global global_id
    
    # Fetch and parse the main category page
    response = requests.get(url)
    if response.status_code != 200:
        print(f"Failed to fetch URL: {url}, Status Code: {response.status_code}")
        return []

    html_content = response.text
    soup = BeautifulSoup(html_content, 'html.parser')
    h3_tags = soup.find_all('h3')
    
    # Store the scraped data
    scraped_data = []

    for tag in h3_tags:
        a_tag = tag.find('a')
        if a_tag:
            link = a_tag.get('href')
            title = a_tag.get('title')
            
            # Fetch and parse the linked page
            try:
                response_title = requests.get(link)
                if response_title.status_code != 200:
                    print(f"Failed to fetch linked page: {link}, Status Code: {response_title.status_code}")
                    continue
                
                title_soup = BeautifulSoup(response_title.text, 'html.parser')
                
                # Find all <p> tags with class "Normal"
                p_tags = title_soup.find_all('p', class_='Normal')
                p_contents = [p_tag.text.strip() for p_tag in p_tags]
                
                content = process_p_contents(p_contents)

                scraped_data.append({
                    'id': global_id,
                    'link': link,
                    'title': title,
                    'content': content,
                    'category': url.split('/')[-1].replace('-', '_')
                })
                global_id += 1

            except Exception as e:
                print(f"Error fetching or parsing linked page: {link}, Error: {e}")
                continue
       
    return scraped_data

if __name__ == '__main__':
    original_url = 'https://vnexpress.net'
    data_dir = r'D:\MY_FOLDER\Project\vuaBE\data'
    
    # Take the categories from the vuacrawl.json file
    # category = ['thoi-su', 'the-gioi', 'kinh-doanh', 'giai-tri', 'the-thao', 'phap-luat', 'giao-duc', 'suc-khoe', 'doi-song', 'du-lich', 'khoa-hoc', 'so-hoa', 'oto-xe-may', 'y-kien', 'tam-su']    
    with open(r'D:\MY_FOLDER\Project\vuaBE\vuacrawl.json', encoding='utf-8') as f:
        data = json.load(f)
        categories = [item['share_url'] for item in data] # results: ['/thoi-su', '/goc-nhin', '/the-gioi']
    
    categories.remove('https://video.vnexpress.net')
    url_list = [original_url + cate for cate in categories]    
    categories = [cate.split('/')[-1] for cate in categories] # truncate '\' for matching folder name
    
    all_data = []
    for i, temp_url in enumerate(url_list):
        data = scrape_data_by_cate(temp_url)
        all_data.extend(data)
        print("Done scraping data for category:", temp_url)
    
    with open(f"{data_dir}/all_data.json", 'w', encoding='utf-8') as f:
        json.dump(all_data, f, ensure_ascii=False, indent=2)
    
    print("Done scraping data for all categories")
