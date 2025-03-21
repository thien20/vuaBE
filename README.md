# Project Overview

This project consists of two main components: the pipeline and the app.

## Pipeline

The pipeline is responsible for scraping data and loading it into a MySQL database. It is implemented in Python and consists of two primary files:

1. **scrape.py**: This script scrapes data from the source.
    - To run: 
    ```
    # activate PVM
    source/simplebe/bin 
    
    cd pipeline
    
    python scrape.py
    ```
2. **ingest.py**: This script loads the scraped data into a MySQL database.
    - To run:
    ```
    # activate PVM
    source/simplebe/bin 
    
    cd pipeline
    
    python ingest.py
    ```

## App

The app is responsible for providing a CRUD API to interact with the data stored in the MySQL database. It is implemented in Go.

- To run the app:
    ```
    cd app
    
    go run .
    ```

Now, go to Postman or browser to take data: `http://localhost:8080/news/the_gioi/`

