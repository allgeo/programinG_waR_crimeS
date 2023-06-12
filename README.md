<div align="center">
  <h3>🚧 Frontend still under construction... 🚧</h2>
</div><br>

## APIs implemented

<pre>
1. POST API to upload a image to AWS S3
2. POST API to create a posts
3. GET API to retreive all posts
</pre>

## Prerequisite

<pre>
Go should be installed on your machine
Mongo should be installed and up and running
</pre>

## Install go dependency

```
go mod download
```

## Define your environment variables in .env file

<pre>
API_PORT=8081
DB_NAME=blog-apis
MONGODB_URL=mongodb://localhost:27017
AWS_BUCKET_NAME=
AWS_REGION_NAME=
AWS_ACCESS_KEY=
AWS_SECRET_KEY=
</pre>

## Build go package

```
go build main.go
```

## Run the service

```
go run main.go
```

## Post collection

Import post collection from file `blog-apis.json`
