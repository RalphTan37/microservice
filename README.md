# microservice - crud application

Crud stands for Create, Read, Update, and Delete.
They are the four basic operations when it comes to data management APIs (Application Programming Interface).
In this crud-app branch, this microservice will store customer order data for an online hypothetical store.

**Disclaimer:** <br />
This crud-app branch is based on the tutorial by Net Ninja on YouTube.
All credit for the tutorial content belongs to Net Ninja.
This branch is intended for personal learning purposes to understand the functionality of microservices and how to build them.
It serves as a resource for reference.

**Important Command Lines:** <br />
```go run main.go``` 
--> Runs the main.go file <br />
```curl -X POST localhost:3000/orders```
--> Sends a post request to localhost:3000/orders; in the log, the create handler is called: create an order <br />
```curl localhost:3000/orders/```
--> Runs a get request to localhost:3000/orders; in the log, the list handler is called: list all orders <br />
```curl localhost:3000/orders/myorders```
--> myorders is a dummy id; in the log, the get by id handler is called: get an order by id <br />
```curl -X PUT localhost:3000/orders/myorders```
--> in the log, the update by id handler is called: update an order by id <br />
```curl -X DELETE localhost:3000/orders/myorders```
--> in the log, the delete by id handler is called: delete an order by id <br />
