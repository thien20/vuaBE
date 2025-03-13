import requests
from bs4 import BeautifulSoup

global_id = 0  # Maintain global counter

def fetch_html(url):
    """Fetch HTML content from a URL."""
    response = requests.get(url)
    if response.status_code == 200:
        return response.text
    else:
        print(f"Failed to fetch {url} - Status code: {response.status_code}")
        return None

def parse_html(html, tag, class_name=None):
    """Parses HTML and finds elements by tag and optional class."""
    soup = BeautifulSoup(html, 'html.parser')
    return soup.find_all(tag, class_=class_name)

def scrape_data_by_cate(url, job_id):
    """Scrapes article links from a category page."""
    global global_id
    html_content = fetch_html(url)
    if not html_content:
        return []
    
    soup = BeautifulSoup(html_content, 'html.parser')
    h3_tags = soup.find_all('h3')
    
    scraped_data = []
    for tag in h3_tags:
        a_tag = tag.find('a')
        if a_tag:
            link = a_tag.get('href')
            title = a_tag.get('title')
            content_data = scrape_article_content(link)
            if content_data:
                scraped_data.append({
                    'job_id': job_id,
                    'id': global_id,
                    'link': link,
                    'title': title,
                    'content': content_data['content'],
                    'category': url.split('/')[-1].replace('-', '_'),
                    'date': content_data['date']
                })
                global_id += 1

    return scraped_data

def scrape_article_content(link):
    """Scrapes content from an individual article page."""
    html_content = fetch_html(link)
    if not html_content:
        return None
    
    soup = BeautifulSoup(html_content, 'html.parser')
    p_tags = soup.find_all('p', class_='Normal')
    p_contents = [p_tag.text.strip() for p_tag in p_tags]

    date_tag = soup.select_one('body > section:nth-of-type(5) > div > div:nth-of-type(2) > div:nth-of-type(1) > span.date')
    date = date_tag.text.strip() if date_tag else 'N/A'

    return {
        'content': " ".join(p_contents),
        'date': date
    }
