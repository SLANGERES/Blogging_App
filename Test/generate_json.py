import random
import json
import requests

# Sample data
titles = [
    "My First Blog", "Learning Go", "Introduction to APIs", "Scaling with Docker",
    "Understanding REST", "Working with Redis", "Exploring gRPC", "CI/CD Basics",
    "System Design 101", "Database Transactions"
]

descriptions = [
    "This is a sample blog post to test the API.",
    "A beginner-friendly blog post.",
    "Exploring advanced concepts in backend development.",
    "Quick thoughts on software design.",
    "Sharing some tips and tricks for new developers.",
]

tags_pool = ["general", "introduction", "go", "python", "backend", "api", "docker", "redis", "cloud", "system-design"]

categories = ["General", "Backend", "DevOps", "Databases", "Cloud"]

# Generate 100 posts
posts = []
for i in range(100):
    post = {
        "title": random.choice(titles) + f" #{i+1}",
        "description": random.choice(descriptions),
        "content": f"This is the detailed content for blog post number {i+1}.",
        "tags": random.sample(tags_pool, k=random.randint(2, 4)),
        "metadata": {
            "likes": random.randint(0, 50),
            "category": random.choice(categories)
        }
    }
    posts.append(post)

# Save to file (optional)
with open("test_posts.json", "w") as f:
    json.dump(posts, f, indent=2)

print("✅ Generated 100 test blog posts in test_posts.json")
