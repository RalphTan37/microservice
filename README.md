# basic-microservice

A microservice is essentially a small and independent unit of an application that does one specific job. <br />

**Disclaimer:** <br />
This basic-microservice branch is based on the tutorial by Net Ninja on YouTube. All credit for the tutorial content belongs to Net Ninja. This branch is intended for personal learning purposes to understand the functionality of microservices and how to build them. It serves as a resource for reference.

**Chi, a third-party dependency** <br />
chi is a lightweight, idiomatic and composable router for building Go HTTP services. <br />
```https://github.com/go-chi/chi``` <br />
+conforms to the HTTP interface of the standard library <br />
+in addition, the middleware package <br />

**Chi Installation** <br />
```go get -u github.com/go-chi/chi/v5```

**Important Command Lines:** <br />
```go run main.go```
--> Runs the main.go file <br />
```curl localhost:3000/Hello```
--> Sends an HTTP get request from the /Hello path
```curl -X POST localhost:3000```
--> Sends an HTTP post request
