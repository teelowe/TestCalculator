# Test Calculator Exercise

## A note on the exercise

Given that Go already has a testing framework (go test), it wasn't entirely clear to me if the goal here was to create a test framework that would completely replace go test, or, if the goal was to essentially extend the existing test setup to be more flexible and useful for the features requested (handling different types of tests, etc..).  I chose the latter.  If the goal was in fact the former, then I'm happy to try again, though I might suggest that such an exercise would be better served using something other than Go, which as I'm sure you are aware, has an existing, ubiquitous and tightly coupled test framework in `go test`.**

to that end...

The way the current calculator tests are setup lacks three primary features that I think would be useful.

1. We lack the ability to run global setup and teardown for the test suite.  In practice, setup could be anything ranging from opening a db connection to setting up test data, etc..., and teardown is fairly self explanatory (imagine for an integration test, we want to kill a running docker container or close a db connection, for instance).  A `TestMain(*m testing.M)` function as provided by the testing library will do the job.

2.  Individual tests as written, lack the ability to run test specific setup and teardown as necessary.  Maybe we want individual tests to run unique pre/post hooks, or perhaps we need to set test specific environment variables, etc...  For this, Go's testing package provides a construct known as subtests.  This essentially, is just adding a name to our table tests, so that we can run subsets of tests using basic regex searches and a `-run=testName` flag to our go test commands.  Pre and Post test hooks can be serviced by global functions called within the tests as needed.

3.  Unless we chose to name our tests pre-pended with "unit" and "integration" (or something similar) and then chose to use `go test -run=unit || integration`, then we lack the ability to run these types of tests separately.  Build flags in Go handle this nicely, but it requires some IDE configuration if you're developing in VSCode.  I've created (but not commited) a local .vscode/settings.json that looks like

````
{
    "go.buildFlags": [
        "-tags=unit,integration"
    ],
    "go.testTags": "unit,integration",
}
````
  and then separated the TestCalculate (integration) and TestParseOperand (unit) tests...

  Adding the associated build tags allows us to run `go test ./... -v -tags=integration`  or `-tags=unit`

###### ** In fact, attempting to create a simple package `test/test_runner.go` to call the test methods in the calculator package, I found that the compiler could not resolve the call, even though the test methods themselves are exported (i.e. function names are capitalized).  As I understand it, this is by design, because naturally, when we compile our code, tests are typically not included. This just isn't a thing that's typically done... or at least I've not come across it yet.  I looked at some other potential approaches, all of which were ugly and struck me as code smell.  If I've missed something obvious, then it's possible I lack the skills you seek at this time.

