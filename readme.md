Project documentation:

Rest API server on Fiber
Project structure:

1)cmd
Here is the entry point of the main.go.

2)internal
The basic logic of the project.
2.1 The app folder contains the application configuration and the Run method.
2.2 The Endpoints folder contains methods that accept a request and send a response.

2.3 The Services folder stores business logic, the Databases folder contains work with the database (Gorm).
2.4 The Models folder contains models for working with the database (Creating tables and methods)
2.5 The Tests folder contains tests

Bot.py - a bot that sends a link to the application

The application itself is launched via React.