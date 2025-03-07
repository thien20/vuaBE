import os
import json
import urllib.parse


class Config:
    def __init__(self, config_file):
        self.config_file = config_file
        self.data = self.load_config()
        self.db_config = self.parse_db_url(self.data.get("db", ""))
        self.redis_config = self.parse_redis_url(self.data.get("redis", ""))

    def load_config(self):
        with open(self.config_file, "r", encoding="utf-8") as f:
            return json.load(f)

    def parse_db_url(self, db_url):
    # Format: "user:password@tcp(host:port)/database?parseTime=true"
        if not db_url:
            return {}

        # Splitting the base DB URL and optional parameters
        user_pass, address_db = db_url.split("@tcp(")
        username, password = user_pass.split(":")
        address, db_params = address_db.split(")/")

        # Extracting database name and parameters
        if "?" in db_params:
            database, params = db_params.split("?")
            params = "?" + params
        else:
            database, params = db_params, ""

        host, port = address.split(":")

        return {
            "username": username,
            "password": password,
            "host": host,
            "port": int(port),
            "database": database,
            # "params": params or "?parseTime=true",
        }

    def parse_redis_url(self, redis_url):
        if not redis_url:
            return {}

        parsed_url = urllib.parse.urlparse(redis_url)
        return {
            "scheme": parsed_url.scheme,  # Should be "redis"
            "host": parsed_url.hostname,
            "port": parsed_url.port,
            "db": int(parsed_url.path.lstrip("/")) if parsed_url.path else 0,
        }
        
# data_path = r"../vuaBE/app/config/config.json"
# config = Config(data_path)
# print("Database Config:", config.db_config)

# config = Config(r"../vuaBE/app/config/config.json")

# print("Database Config:", config.db_config)
# print("Redis Config:", config.redis_config)
