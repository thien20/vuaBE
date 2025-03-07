# Use an official Python runtime as a base image
FROM python:3.11-slim

# Set the working directory
WORKDIR /app

# Copy only requirements first for caching dependencies
COPY requirements.txt ./

# Install dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Copy the rest of the application
COPY . ./

# Command to run the consumer
CMD ["python", "consumer.py"]
