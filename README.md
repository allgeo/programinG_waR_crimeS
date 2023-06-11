<div align="center">
  <h3>🚧 Frontend still under construction... 🚧</h2>
</div><br>

<img src="https://github.com/allgeo/programinG_waR_crimeS/assets/62227321/3f8a2819-e3ed-4382-b2b5-7fffafaeb825" width="200">

## APIs implemented

<pre>
1. POST API to upload a image
2. POST API to create a post
3. GET API to retreive all posts.
4. Get API for weather 
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

```
API_PORT=8081
DB_NAME=blog-apis
MONGODB_URL=mongodb://localhost:27017
WEATHER_BASE_URL=https://api.openweathermap.org
WEATHER_API_KEY=
AWS_BUCKET_NAME=
AWS_REGION_NAME=
AWS_ACCESS_KEY=
AWS_SECRET_KEY=
```

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
