# Usecases

Usecase ties services together to perform a specific task. They are responsible for orchestrating the services and
handling the flow of data between them.

One could argue that the logic could sit in the handler however there are few disadvantages to this approach. Testing is
far more difficult because you have to do handler setup in addition to set up for your usecase.

The second reason is that it makes the handler more complex. The handler should be responsible for handling the request
and response and not the business logic.

By separating the handler and the usecase, we also make it much easier to offer the same usecase via different handlers.
For example, we could have a REST handler and a gRPC handler that both use the same usecase. Because the handlers
responsibility is only to handle the request and response mapping we can easily swap out the handler without changing
the business logic.
