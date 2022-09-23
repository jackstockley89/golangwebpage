<h1 align="center" style="font-family: 'Lucida Sans', Arial, sans-serif"><b>golangwebpage</b></h1>

&nbsp;&nbsp;&nbsp;&nbsp;
![GitHub pull-requests](https://img.shields.io/github/issues-pr/jackstockley89/golangwebpage?style=for-the-badge)
![GitHub pull-requests closed](https://img.shields.io/github/issues-pr-closed/jackstockley89/golangwebpage?style=for-the-badge)
![GitHub last-commit](https://img.shields.io/github/last-commit/jackstockley89/golangwebpage?style=for-the-badge)
![GitHub contributors](https://img.shields.io/github/contributors/jackstockley89/golangwebpage?style=for-the-badge)
<br/>

<h2 align="left" style="font-family: 'Lucida Sans', Arial, sans-serif"><b>Actions</b></h2>

[![Pull Request Authentication Check](https://github.com/jackstockley89/golangwebpage/actions/workflows/auth-check.yml/badge.svg)](https://github.com/jackstockley89/golangwebpage/actions/workflows/auth-check.yml)
[![Terraform Plan infra/EKS](https://github.com/jackstockley89/golangwebpage/actions/workflows/terraform-plan.yml/badge.svg)](https://github.com/jackstockley89/golangwebpage/actions/workflows/terraform-plan.yml)
[![localstack-action-example](https://github.com/jackstockley89/golangwebpage/actions/workflows/localstack-testing.yml/badge.svg)](https://github.com/jackstockley89/golangwebpage/actions/workflows/localstack-testing.yml)
[![Go Test](https://github.com/jackstockley89/golangwebpage/actions/workflows/go-test.yml/badge.svg)](https://github.com/jackstockley89/golangwebpage/actions/workflows/go-test.yml)

[![Publish Docker image](https://github.com/jackstockley89/golangwebpage/actions/workflows/docker-build-and-push.yml/badge.svg)](https://github.com/jackstockley89/golangwebpage/actions/workflows/docker-build-and-push.yml)
[![Continuous Deployment](https://github.com/jackstockley89/golangwebpage/actions/workflows/cd.yml/badge.svg)](https://github.com/jackstockley89/golangwebpage/actions/workflows/cd.yml)
[![CodeQL](https://github.com/jackstockley89/golangwebpage/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/jackstockley89/golangwebpage/actions/workflows/codeql-analysis.yml)
[![Open Issues](https://github.com/jackstockley89/golangwebpage/actions/workflows/open-issues.yml/badge.svg)](https://github.com/jackstockley89/golangwebpage/actions/workflows/open-issues.yml)
[![Pull Request Checker](https://github.com/jackstockley89/golangwebpage/actions/workflows/pr-checker.yml/badge.svg)](https://github.com/jackstockley89/golangwebpage/actions/workflows/pr-checker.yml)
<br/>

<h2 align="left" style="font-family: 'Lucida Sans', Arial, sans-serif"><b>Prerequisite</b></h2>

---
&nbsp;&nbsp;&nbsp;&nbsp;
![GO](https://img.shields.io/github/go-mod/go-version/jackstockley89/golangwebpage)
![Docker](https://img.shields.io/badge/Docker-v20.10.2-blue)

<br/>

<h2 align="left" style="font-family: 'Lucida Sans', Arial, sans-serif"><b>Overview</b></h2>

---
Example of a web server created using GO, running from a Docker container

<br/>

<h2 align="left" style="font-family: 'Lucida Sans', Arial, sans-serif"><b>How to Run in Docker Container</b></h2>
 
---

From repository firstly have to build the image using the following
```
docker build -t "cycling-blog" .
```

Once the image has completed building, this is now ready to run 
```
docker run -d -p 8080:8080 cycling-blog:latest
```

<br/>

<h2 align="left" style="font-family: 'Lucida Sans', Arial, sans-serif"><b>How to Run in Docker Compose</b></h2>
 
---

The docker-compose.yml contains a volume mounted to the local drive to prevent data lost within the database if shutdown. This would required edit if to be used on local machine.

To run the application with database this can be done using docker compose command. The following command build and start the containers in the background
```
docker-compose up --build -d
```

To stop the application with and database with
```
docker-compose down
```
