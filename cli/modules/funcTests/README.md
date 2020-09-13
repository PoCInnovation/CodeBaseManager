# CodeBaseManager/cli - Functional Tests

This module is highly inspired by [JenRik](https://github.com/Yohannfra/JenRik).

The goal of this module is to allow the user to describe tests for their binaries through simple [TOML](https://github.com/toml-lang/toml) files.

## Quick Start
Let's take a simple go program, say `helloworld.go`
```go
func main() {
    if len(os.Args) == 1 {
        fmt.Println("hello world")
    } else {
        fmt.Fprintln(os.Stderr, "yeah.. no")
        os.Exit(1)
    }
}
```
Then you can write tests which can be as simple as this
```TOML
[[Test]]
    name = "OK" # Name of the test
    bin = "helloworld" # The binary being tested
```
Of course you may want to write some more tests and have more control over them. Each file is a TestSuite so you can write multiple tests inside a file.
```TOML
[common] # Parameters for the TestSuite
    name = "My hello world"
    desc = "TestSuite of this special special version of mine"
    bin = "helloworld" # Binary used in all tests
    
[[Test]]
    name = "ok"
    [Test.expected]
        status = 0
        stdout = "hello world"

# This test expect helloworld to exit with 0 after printing "hello world"

[[Test]]
    name = "failure"
    [Test.expected]
        status = 1
        stderrFile = "path/to/expected/file_exp"
        
# This one expect helloworld to exit with 1 after printing on stderr the same thing as "file_exp"'s content
```

## Details
Each test files represent a TestSuite. A test file has up to 2 of the following sections:

### [common]
This section is **optional** as well and can be used to set `commons` `[[Test]]`'s field with a value for the whole TestSuite (except `name` & `desc` which will then describe the TestSuite).

Any field defined in this section will be overriden if redefined in a test. The fields are described in the next section.

### [[Test]]
This section is **mandatory** as it describes the tests. It is a [TOML array of section](https://github.com/toml-lang/toml#array-of-tables).
Here's a brief exemple:
```TOML
[[Test]]
    name = "Test 1"
    desc = "This a useless test"
    bin = "ls"
    refBin = "ls"
    
 # This test will compare ls's output & exit status to ls.

[[Test]]
    name = "Test 2"
    desc = "This is useless & stupid too but will fail"
    bin = "ls"
    args = ["-h"]
    refBin = "ls"

# This test will compare `ls -h`'s output & exit status to ls.
```
Following is the explanation of all available fields.

#### Commons
These fields are present directly under a `[common]` or `[[Test]]`, they give basic information about the Test(Suite).
|  Fields |           Usage           |      Type     | Remark |
|:-------:|:-------------------------:|:-------------:|--------|
| name    | Test(Suite)'s name        | string        | **Mandatory** for each `[[Test]]` |
| desc    | Test(Suite)'s description | string        | **Optional** |
| bin     | Binary being tested       | string (path) | **Mandatory**|
| refBin  | Binary used as reference  | string (path) | **Optional** |
| args    | Args passed to `bin`      | [string]      | **Optional** or **ignored** if given as `[common]` |
| refArgs | Args passed to `refBin`   | [string]      | **Optional** or **ignored** if given as `[common]` |

> :question: **Notes**
>
> If `refBin` is provided then it will go through the same process as `bin` and checks will be performed against the outputs of `refBin`; therefore the `[Test.expected]` section of each `[[Test]]` can be fully omitted, if it isn't it will be used for checks rather than the `refBin`'s outputs.
>
> :warning: If the given binaries aren't located in `$PATH` then the path should be provided.
>
> :warning: If the given binaries can't be found or if no `bin` was provided then the test is **ignored**.
> In future updates, if `bin` isn't provided then we shall use the one the backend knows.

#### Expected
These fields are present as a sub section of either `[common]` or `[[Test]]`, such as:
```TOML
[[Test]]
    desc = "Failing Test"
    bin = "ls"
    args = ["-e"]
    [Test.expected]
        status = 1
        stderrFile = "path/to/ls_dashe_output"

# This test will execute "ls -e" and it's exit status & stderr will be checked against the expected status and the content of `ls_dashe_output`.
```
As the name suggest, this section is used to check if the test succeeded. The fields are the following:
|   Fields   |           Usage         |      Type     | Remark |
|:----------:|:-----------------------:|:-------------:|--------|
| status     | Expected exit status    | int           | **Optional**, if not provided will expect 0 |
| stdout     | Expected stdout         | string        | **Optional**, if provided **do not** use `stdoutFile` |
| stderr     | Expected stderr         | string        | **Optional**, if provided **do not** use `stderrFile` |
| stdoutFile | Compares stdout to file | string (path) | **Optional**, if provided **do not** use `stdout` |
| stderrFile | Compares stderr to file | string (path) | **Optional**, if provided **do not** use `stderr` |

> :question: **Notes**
>
> If not provided, then the default test is to check if the binary exited without error.
>
> If a `refBin` was provided it's outputs will be used by default.
> Defining status will override the check on `refBin`'s status. The same logic applies for `stdout` and `stderr`.
>
> :warning: If both `stdx` and `stdxFile` are provided, `stdxFile` will be **ignored**

### Interactions
These fields are present as a sub section of either `[common]` or `[[Test]]`, such as:
```TOML
[[Test]]
    bin = "cat"
    [Test.interactions]
        stdin = "Hello world"
        stdoutPipe = 'grep "world"'

# This test will execute cat "hello world" as its input. Its output will be redirected to 'grep "world"'
# and it will be checked afterward.
```
This section is used to 
|   Fields   |                         Usage                    |       Type     | Needed |
|:----------:|:------------------------------------------------:|:--------------:|--------|
| stdoutPipe | Redirects stdout to a command before checking it | string (cmd)   | **Optional** |
| stderrPipe | Redirects stderr to a command before checking it | string (cmd)   | **Optional** |
| stdinPipe  | Redirects a command to your stdin                | string (cmd)   | **Optional** |
| stdinFile  | Uses the file's content as your stdin            | string (path)  | **Optional** |
| stdin      | Used as your stdin                               | string         | **Optional** |
| pre        | Runs a command before the test                   | string (cmd)   | **Optional** |
| post       | Runs a command after the test                    | string (cmd)   | **Optional** |
| env        | Serves as the env for your `bin`                 | [string] (key=value) | **Optional** |
| addEnv     | Adds to the current env all of the given pairs   | [string] (key=value) | **Optional** |
> :question: **Notes**
>
> :warning: Only one of the stdin fields should be provided. If multiple stdin fields are provided then CBM will use the first one found in this order: stdin, stdinFile & stdinPipe.

## Predicted

When describing tests, options to expect a failure, a timeout or to time or repeat the test.

When describing tests, possibility to declare variables to ease the writing of a TestSuite under a `[vars]` table.

Usage of CBM's backend to be able to omit the `bin` field & add new fields such as `build`
