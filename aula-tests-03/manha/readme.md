# Exercise 01

No errors found

# Exercise 02

Report:

coverage.out
coverage.html


# Exercise 03

I see no reason to test main.go

db.go is a Mocked DB, and the untested line is untested because the Update method on the service calls the ReadAll method before, throwing back an error if unable to find, therefore the untested line is currently never actually used. I *could* create a new function that uses that line, but why though?