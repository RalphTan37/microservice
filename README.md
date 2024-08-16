# microservice - using redis

Redis is an in-memory data structure store which makes it really fast.
The downside is that the data stored is typically ephermeral (lasts for a short time).
It also doesn't usually persist across restarts.
(Can configure redis to be persistent and fault tolerant, but not as safe as using postgres). <br />
*Installed & Utilized Docker: <br />
Docker is an open-source platform that automates the deployment, scaling, and management of applications using containerization (lightweight, standalone, and executable software packages). <br />



**Disclaimer:** <br />
This using-redis branch is based on the tutorial by Net Ninja on YouTube.
All credit for the tutorial content belongs to Net Ninja.
This branch is intended for personal learning purposes to understand the functionality of microservices and how to build them.
It serves as a resource for reference.

**Third-Party Redis Dependency:** <br />
```https://github.com/redis/go-redis``` <br />

**Install Redis:** <br />
```go get github.com/redis/go-redis/v9``` <br />

**Important Command Lines:** <br />
```go run main.go``` --> Runs the main.go file
```curl localhost:3000/``` --> Sends an HTTP GET request to web server at port 3000
```redis-server --version``` --> Starts the Redis Server <br />
```docker ps``` --> Lists the running containers on Docker host <br />
```sudo service redis-server stop``` --> Stops Redis Server <br />
```docker run -p 6379:6379 redis:latest``` --> Download and run the latest Redis image and bind the system 6379 port to the docker container 6379 port <br />
```redis-cli``` --> Interacts w/ Redis Database (Can enter KEYS * - retrieves all keys in selected database) <br />

