# Architecture for module separation

## Context and Problem Statement

In order to make the code more readable and maintainable, the code should organised in modules.
Module should be separated by the domain they are related to. For example cart subdomain should be placed in the cart module.
Communication style between modules will be decided later.

## Considered Options

* microservice architecture
* modular monolith architecture

## Decision Outcome

Chosen modular monolith architecture, because it is easier to maintain and deploy.
It is also easier to migrate to microservice architecture in the future.