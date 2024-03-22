# Test Calculator Exercise

## A note on the exercise

Given that Go already has a testing framework (go test), it wasn't entirely clear to me from the provided description if the goal here was to create a test framework that would be usable instead of `go test`, or, if the goal was to essentially extend the existing test setup to be more flexible and useful for the features requested (handling different types of tests, etc..).  I chose the latter.  Checking with Mark for clarification revealed that this was the right choice.  I was hard pressed to come up with a reason that anyone would want to work around the "testing" package and go test, but "ok", I thought... "for science!", said I. **

to that end....

The way the current calculator tests are setup lacks three primary features that I think would be useful.

1. We lack the ability to run global setup and teardown for the test suite.  In practice, setup could be anything ranging from opening a db connection to setting up test data, etc..., and teardown is fairly self explanatory (imagine for an integration test, we want to kill a running docker container or close a db connection, for instance).  A `TestMain(*m testing.M)` function as provided by the testing library will do the job here.  I've chosen to move the global TestMain into it's own file, `calculator/main_test.go`

2.  Individual tests as written, lack the ability to run test specific setup and teardown as necessary.  Maybe we want individual tests to run unique pre/post hooks, or perhaps we need to set test specific environment variables, etc...  For this, Go's testing package provides a construct known as subtests.  This essentially, is just adding a name to our table tests, so that we can run subsets of tests using basic regex searches and a `-run=testName` flag to our go test commands.  Pre and Post test hooks can be serviced by global functions called within the tests as needed, or these functions could be defined on a test by test basis.  For illustration purposes, I've opted to define them globally in `calculator/main_test.go`.  Also note that I've commented them out in the tests to save a bit of verbose test output.  They can easily be uncommented to see where they execute in the flow of the tests.

3.  Unless we chose to name our tests pre-pended with "unit" and "integration" (or something similar) and then chose to use `go test -run=unit || integration`, then we lack the ability to run these types of tests separately.  Build flags in Go handle this nicely, but it requires some IDE configuration if you're developing in VSCode.  I've created (but not commited) a local .vscode/settings.json that looks like

````
{
    "go.buildFlags": [
        "-tags=unit,integration"
    ],
    "go.testTags": "unit,integration",
}
````
  and then separated the TestCalculate (integration) and TestParseOperand (unit) tests into their own files within the `calculator` package.

  Adding the associated build tags:
  ````
  //go:build unit
  //go:build integration
  ````
   allows us to run `go test ./... -v -tags=integration`  or `-tags=unit`

# Tests
I've added some additonal (albeit trivial) test cases to both `TestCalculator` and `TestParseOperand` and marked my additions with a comment.  What I found interesting about the tests is that, because our tests are testing the calculator logic after our command arguments are converted from string to float64 (in main.go), I was able to generate errors on the command-line that I could not replicate in the tests. Example:
````
command-line:
TestCalculator:$ go run main.go 2 + 2 + 2 + 2
2024/03/22 13:20:27 Usage: main <operand1> <operator> <operand2>
exit status 1
````
in the unit test however, this works, because we are using floats and Go just does the math...
````
{"borked operands", 2 + 2, "+", 2 + 2, 8, nil},
````
I suspect the argument parsing logic could be tested in the main package with some trickery, but given time constraints, I opted not to get into this.***

# CI/CD
I've opted to use github actions, only because it's what I've used most recently, and the fact that it's nicely coupled to Git makes things easy.****
Note that I've included both build flags in the test command in `pr.yml`:
````
with:
  testArguments: '-v ./... -tags unit,integration'
````
to ensure we get all the test output.  I'm using [robherley/go-test-action@v0.2.0](https://github.com/marketplace/actions/go-test-action) to handle pretty test output.  You can view a nice pie chart and foldable test results in the github actions panel of my [fork](https://github.com/teelowe/TestCalculator/actions/runs/8395659679).

### Things I might have done with more time
- a Dockerfile
- a merge.yml file that runs on merge to main and builds and pushes said Dockerfile to a registry
- a Makefile for easy execution of all the things (though honestly, the go toolset is fine)
- a facility for loading text files containing calculations and validating the results - sort of an additional form of user validation testing for the less technical...

footnotes:

###### ** In fact, attempting to create a simple package `test/test_runner.go` to call the test methods in the calculator package, I found that the compiler could not resolve the call, even though the test methods themselves are exported (i.e. function names are capitalized).  As I understand it, this is by design, because naturally, when we compile our code, tests are typically not included. This just isn't a thing that's typically done... or at least I've not come across it yet.  I looked at some other potential approaches, all of which were ugly and struck me as code smell.  
###### *** incidentally, I also noticed that the command-line barfs unless you escape the multiplication operator, i.e. `2 \* 2`.  Presumably this is because the operator gets interpreted as a shell wildcard?
###### **** I also really love Gitlab
