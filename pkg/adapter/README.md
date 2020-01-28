# Ports & Adapters (Hexagonal Architecture)
I thought I explain what is the purpose of this Domain `adapter`. In order to follow
the clean architecture design (The onion layers with having Dependency Inversion in
mind). I created this domain which could act as Adapter and Port to our Application.
Any other domains inside the `adapter` could be considered the **Infrastructure layer**
in context of Domain Driven Design.


## HTML (GoQuery package)
It uses GoQuery to read HTML content and identify the DOM elements and find the target 
elements recursively.


## HTTP _(Fast HTTP package)_
It uses fast http to make all the GET request calls. This library make the calls much faster than 
built-in library due it's use of concurrency.


## Use Case
If we wish to swap `GoQuery` with another HTML DOM library or the built-in library `net/html`, 
we don't need to modify the code inside the core domain of our application, we could simply just
write a new adapter and inject that adapter to the main HTML package. This reduces
the hours of development and promotes decoupling from third party dependencies hence the
cleaner and more maintainable code.


## Other Scenarios e.g. Databases, Queues, and ...
For example if we wish to use a queue system we could create a new domain `queue` inside
the `adapter` and create a new adapter for AWS SQS. If we wish to use another provider 
you could create a new adapter for new client for example Azure Service Bus and inject 
it to main `queue` package.