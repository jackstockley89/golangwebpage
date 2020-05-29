<h1 align="center" style="font-family: 'Lucida Sans', Arial, sans-serif"><b>golangwebpage</b></h1>

&nbsp;&nbsp;&nbsp;&nbsp;
![GitHub pull-requests](https://img.shields.io/github/issues-pr/jackstockley89/golangwebpage?style=for-the-badge)
![GitHub pull-requests closed](https://img.shields.io/github/issues-pr-closed/jackstockley89/golangwebpage?style=for-the-badge)
![GitHub last-commit](https://img.shields.io/github/last-commit/jackstockley89/golangwebpage?style=for-the-badge)
![GitHub contributors](https://img.shields.io/github/contributors/jackstockley89/golangwebpage?style=for-the-badge)

<br><br/>

 <h2 align="left" style="font-family: 'Lucida Sans', Arial, sans-serif"><b>Prerequisite</b></h2>

---
&nbsp;&nbsp;&nbsp;&nbsp;
![GO](https://img.shields.io/badge/GO-v1.14.3-blue)
![Docker](https://img.shields.io/badge/Docker-v19.03.8-blue)

<br><br/>

<h2 align="left" style="font-family: 'Lucida Sans', Arial, sans-serif"><b>Overview</b></h2>

---
Example of a web server created using GO, running from a Docker container

<br><br/>

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