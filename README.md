# Go Hello API
A simple API project built with Go.

# Overview
This project demonstrates a basic web server that exposes an API endpoint, written in Go.

This was one of the tasks given to me at [HNG](https://www.hng.tech/) Internship 11.

# Live Demo
Access the live demo [here](https://go-hello-api.up.railway.app/api/hello?visitor_name=Name).

# Endpoint
`[GET] /api/hello?visitor_name="Name"`

# Response
You should get this as a response:

```
{
     "client_ip": "127.0.0.1", // The IP address of the requester
     "location": "New York, United States" // The city and country of the requester
     "greeting": "Hello, Name!, the temperature is 11 degrees Celcius in New York, United States"
}
```
