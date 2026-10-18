Go Tests

Procedure

1. Run all tests from the project root with `go test ./...`.
2. To target a specific package or individual test, use `go test` with the relevant package and optional `-run` arguments.
3. If tests fail, analyze the error and fix its root cause.
4. After modifying any production code or tests, run the relevant tests and verify the results before finishing.

Examples

Run all tests:

go test ./...

Run all tests in a specific package:

go test ./app/shared/types/primitive

Run a specific test:

go test ./app/shared/types/primitive -run TestNotifierStartsInactive
