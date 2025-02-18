import requests

def trigger_scrape(job_topic):
    response = requests.post(f"http://localhost:8080/jobs/fetch/{job_topic}")
    print(f"Job Triggered: {response.status_code}, {response.json()}\n")
