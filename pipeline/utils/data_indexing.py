from elasticsearch import Elasticsearch, helpers

client = Elasticsearch("http://elastic:9200")


def indexing_only_text_with_bulk(data, index_name):
    if not client.indices.exists(index=index_name):
        mappings = {
            "mappings": {
                "properties": {
                    "job_id": {"type": "keyword"},
                    "category": {"type": "text"},
                    "link": {"type": "keyword"},
                    "title": {"type": "text"},
                    "content": {"type": "text"},
                    "scraped_at": {"type": "date", "format": "yyyy-MM-dd HH:mm:ss"}
                }
            }
        }
        client.indices.create(index=index_name, body=mappings)

    actions = [
        {
            "_index": index_name,
            "_id": doc.get("id"),
            "_source": doc
        }
        for doc in data
    ]

    # Perform bulk indexing
    helpers.bulk(client, actions)

    # Refresh index
    client.indices.refresh(index=index_name)

    print(f"Successfully indexed {len(actions)} documents into {index_name}")


def indexing_with_imgs_vid():
    pass

