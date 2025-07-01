# authGRPC Project
AuthGRPC is a project that implements authorization and authentication within a microservices architecture using gRPC. The project employs logging through the slogger library and parses configuration files with the help of cleanenv. SQLite is used as the database for data storage.
## Description
This project consists of several key layers that are crucial for its functionality:

### Transport Layer: 
The transport layer is responsible for implementing gRPC to enable communication between different services in the microservices architecture. gRPC is used here to facilitate high-performance, language-neutral communication over HTTP/2, providing efficient and fast communication between the services.
### Service Layer: 
The service layer is where the core business logic resides. It handles the processing of data, including the implementation of authorization and authentication workflows. This layer ensures that the incoming requests are validated and processed correctly before interacting with the database or other services.
### Database Layer: 
SQLite is utilized as the database for storing user credentials and session data. The database layer interacts with SQLite to handle the storage and retrieval of data, ensuring that all actions related to authentication and authorization
### Authorization and Authentication:
The project implements a comprehensive authorization and authentication system. It allows users to register, log in, and manage their sessions in a secure manner. This process is implemented through microservices that communicate with each other using gRPC. 
### Testing: 
To ensure the reliability of the service, a set of tests has been written to verify the functionality of the system. These tests ensure that both the authentication and authorization mechanisms work as expected, and they also test the interaction between the transport, service, and database layers. Running these tests guarantees that the service performs as intended and is resilient to errors or unexpected behaviors.
## Sumary
In summary, AuthGRPC is a secure and scalable solution for managing user authentication and authorization within a microservices-based environment, utilizing gRPC for communication, SQLite for data persistence, and structured layers for maintainable and testable code. The project has been designed with extensibility and performance in mind, making it suitable for various applications requiring secure user access management.
