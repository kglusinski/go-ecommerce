# Use HTTP Framework

## Context and Problem Statement

In order to save time and effort, we should use a HTTP framework.

## Considered Options

* Echo
* Gin
* Gorilla

## Decision Outcome

Chosen option is "Echo", because it saves time on binding HTTP request bodies to specific structs.
Solves most common problems with HTTP requests.