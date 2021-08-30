# gRPC server & client applications following DDD architecture

This project attempts to expose how to use DDD and gRPC given a scenario where we need a grpc server and its client to interact as server/client using gRPC and protocol buffers. The client itself exposes also a REST API to communicate with users.


## Design
The design has followed DDD concepts to structure the code, there are 4 separated layers with specific responsibilities:
* Interface: responsibles for the interaction with user, the software exposes or receives data from external users.
* Application: This is a small layer that links interface and domain layers, it may call domain services to be used in application purposes.
* Domain: The heart of our system, this layer holds domain logic and business knowledge.
* Infrastructure: It's the layer that serves as a link for the rest of layers. It may contain external services such as database implementations.

I have tried to make the implementation as reusable as possible, using Repository pattern to easily swap the current in memory implementation by a DB one.
For the REST API, I have used Echo, which a fast and easy option, however the current implementation of the REST API, relies on Bind method for JSON payload processing, which may not be the best option for big payloads.

I have followed TDD and made some mocks to test everything I needed.


## Run the application
* The application could have been run simply using the makefile and Docker, but it still needs refinment, there's a problem related to certificates so currently run the apps in docker is not working.
* However, the whole application server/client has been tested and it's working locally, to do so, simply run the following commands:

From the root of repository:
```sh
$ make up
```


## Testing Client-Server with curls

With both server and client running the complete workflow can be tested with two curls:

Store some Data
``` sh
 curl --request POST \
--url localhost:3001/grpc-ports \
--header 'Content-Type: application/json' \
--data '{"AEAJM":{"name":"Ajman","city":"Ajman","country":"United Arab Emirates","alias":[],"regions":[],"coordinates":[55.5136433,25.4052165],"province":"Ajman","timezone":"Asia/Dubai","unlocs":["AEAJM"],"code":"52000"},"AEAUH":{"name":"Abu Dhabi","coordinates":[54.37,24.47],"city":"Abu Dhabi","province":"Abu Z¸aby [Abu Dhabi]","country":"United Arab Emirates","alias":[],"regions":[],"timezone":"Asia/Dubai","unlocs":["AEAUH"],"code":"52001"}}'
```

Read Data
```sh
curl localhost:3001/grpc-ports
```

## Author
* **Adolfo Rodriguez** - *grpc-ddd-ports* - [adolsalamanca](https://github.com/adolsalamanca)


## References
* [Protocol Buffers](https://developers.google.com/protocol-buffers)
* [Echo](https://echo.labstack.com/)


## License
This project is licensed under MIT License.
 