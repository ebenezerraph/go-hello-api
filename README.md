# Go Hello API
A simple API project built with Go.

![](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

## Overview
This project is a simple Go web server that exposes an API endpoint.

> This was one of the tasks assigned to us during the [HNG](https://www.hng.tech/) Internship 11 — no, I didn't make it to the finals. 🙃

## Endpoint
`[GET] /api/hello?visitor_name="Name"`

## Response
```
{
     "client_ip": "127.0.0.1", // The IP address of the requester
     "location": "New York, United States" // The city and country of the requester
     "greeting": "Hello, Name!, the temperature is 11 degrees Celcius in New York, United States"
}
```
