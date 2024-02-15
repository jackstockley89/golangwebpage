# golangwebpage

## Overview
Example of a web server created using GO, running from a Docker container

## How to Run in Docker Container
From repository firstly have to build the image using the following
```
docker build -t "cycling-blog" .
```

Once the image has completed building, this is now ready to run 
```
docker run -d -p 8080:8080 cycling-blog:latest
```
