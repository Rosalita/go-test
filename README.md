# go-test
Testing with Go

Recommended Assertion package
`github.com/stretchr/testify/assert`

To run all tests
`go test`

To run just one test
`go test -run TestNameOfTest`

Coverage Report
`go test -coverprofile cover.out`
`go tool cover -html=cover.out -o cover.html`

Looper tool for TDD
`https://github.com/nathany/looper`

# 00
Download and install Go (https://golang.org/dl/)
Hello Atelier! in Go

# 01 
Testing a single function
Running a coverage report
Adding a missing tests
Run coverage again
Add test for lower case name
Fix the code

# 02 
Test Driven Development
Red Green Clean example
First, define function signature
First test for 0
Test for 1, use if statements
Test for 2, use a for loop define numeral outside the loop decrement and build up numeral
Test for 4, convert logic to a switch statement
Special cases for 1,4,5,9,10, 40 is XL


# 03 
Testing randomness
Making tests deterministic

# 04
Testing around a database

# 05 
Using httptest package to test an API Client

