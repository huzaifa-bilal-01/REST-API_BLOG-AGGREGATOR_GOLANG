# Blog Aggregator

## Overview

This is a Blog Aggregator built in Go that fetches niche blog posts and stores them in a PostgreSQL database. It enables users to aggregate and read blog posts from various sources conveniently.

## Features

- **Fetching:** The aggregator fetches blog posts after every 10s from predefined sources based on specified niches.
- **Storage:** Blog posts are stored in a PostgreSQL database for easy retrieval and management.
- **Scheduled Updates:** Set up a schedule for automatic updates to keep the database current with the latest blog posts.

## Prerequisites

Make sure you have the following installed:

- Go
- PostgreSQL

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/huzaifa-bilal-01/REST-API_BLOG-AGGREGATOR_GOLANG.git
   cd blog-aggregator

2. Create and set the .env file as follow:

   ```
   PORT=8080
   DB_URL=postgres://<username>:<password>@localhost:5432/<database_name>?sslmode=disable

3. Run the following command and throgh postman or thinderclient test the api:

   ```bash
   go build
   ./<name of executable file>

## Contact
For any inquiries or issues, please feel free to conatct me at [huzaifabilal226@gmail.com](mailto:huzaifabilal226@gmail.com)
  

