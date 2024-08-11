# microservice - using redis

Redis is an in-memory data structure store which makes it really fast.
The downside is that the data stored is typically ephermeral (lasts for a short time).
It also doesn't usually persist across restarts.
(Can configure redis to be persistent and fault tolerant, but not as safe as using postgres).

**Disclaimer:** <br />
This using-redis branch is based on the tutorial by Net Ninja on YouTube.
All credit for the tutorial content belongs to Net Ninja.
This branch is intended for personal learning purposes to understand the functionality of microservices and how to build them.
It serves as a resource for reference.

**Third-Party Redis Dependency:** <br />
```https://github.com/redis/go-redis``` <br />

**Install:**
```go get github.com/redis/go-redis/v9``` <br />


**Important Command Lines:** <br />

