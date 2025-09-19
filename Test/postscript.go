import requests
import json

# Load generated posts
with open("test_posts.json") as f:
    posts = json.load(f)

url = "http://localhost:8000/blogs"  # 🔴 Replace with your API endpoint

for i, post in enumerate(posts, 1):
    response = requests.post(url, json=post)
    print(f"Post {i}: Status {response.status_code}, Response: {response.text}")

