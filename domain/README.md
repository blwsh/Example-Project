# Domain

This is the domain module. It represents entities which are native and universal to our application.

There is a distinction to be made between service entities and domain entities. Domain entities like YearMonth and Money
are universal and can be used in any application. Service entities like Product and Order are specific to the
application
and should be placed in the appropriate service module.

To make things easier you may want to promote an entity from a service entity to a domain entity, but you should only do
so after careful consideration. Some questions to ask yourself are:

* Am I frequently mapping the service entity attributes from service to service?
* Is the entity used in multiple services?

If the answer is yes to either of these questions then you may want to consider promoting the entity to a domain entity.

However, you should be cautious when expanding your domain. The larger the domain the more complex the application
becomes. This is primarily because you have to consider the interactions between the entities and the business rules
that govern them. Preferring service entities over domain entities can help keep the domain small and focused and helps
keep your services decoupled.

A good example of a service entity that is likely to be promoted to a domain entity is a User entity. This is because
users are a common entity in many applications and are likely to be used in multiple services.
