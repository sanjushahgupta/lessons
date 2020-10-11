## overview
- monolith (albeit very small and thus can be viewed as a microservice)
- `main.go` contains everything needed to start the application, `go run main.go` to start
- the rest of the directories contains reusable codes to be used in `main.go`

## flow
- function calls usually goes in this direction `main.go` -> `handler` -> `application` -> `storage` and back
- `handler` contains codes related to handling request/response
- `application` contains codes related to business logic
- `storage` contains codes to interact with the storage system
- `domain` contains shared codes, used in almost all other parts

## Dependency Injection (DI) - why?
- easier to do unit tests
- easier to extend
- more complex

## best practices
- documentations -> comment on all exported types/functions
- db migration (TBD)
- testings (TBD)

## TODO
- `grep -r TODO` to find the list of TODOs
- `grep -r DETAILS` to find the list of DETAILS/GitHub issues relevant to the TODO
- `git checkout -b <branch_name_related_to_todo>` to checkout and work on that to do
- when done, commit and push to remote branch
