# go-test
Unit testing in Go

Recommended Assertion package
`github.com/stretchr/testify/assert`

To run all tests
`go test`

To run just one test
`go test -run TestNameOfTest`

Coverage Report
`go test -coverprofile cover.out`
`go tool cover -html=cover.out -o cover.html`

# 01 
Testing a single function
Adding functionality
Running a coverage report
Adding missing tests

# 02 
Test Driven Development
Red Green Clean example

# 03 
Testing randomness
Making tests deterministic

# 04
Testing around a database

# 05 
Using httptest package to test an API Client

