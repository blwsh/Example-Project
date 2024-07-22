# Services

Services are the core of the application. They are responsible for handling the business logic and data manipulation. 

Prefer primitive inputs for services and DTOs for service inputs and outputs. Services should have dependencies injected via the constructor.

An example dependency would be a repository. The repository would be injected into the service via the constructor. This allows for easy testing and mocking of the repository.
We do this to ensure that the service is only responsible for the business logic and not the data access.

This also makes testing easier because we can mock the repository and test the service in isolation.
