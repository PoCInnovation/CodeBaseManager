# CodeBaseManager - Functional Tests

This module is highly inspired by [JenRik](https://github.com/Yohannfra/JenRik).

The goal of this module is to allow the user to describe tests for their binaries through simple [TOML](https://github.com/toml-lang/toml) files.

## Quick Start
WIP...

## Details
Each test files represent a TestSuite. A test file has up to 3 of the following sections:

### [vars]
This section is **optional** and can be used to declare variables which can later be used **in a literal string** like `'...$myvar...'`.
Here's a brief exemple:
```TOML
 [vars]
    longPath = "path/to/smth/very/deep/somewhere/here"
  
 [default]
    bin = '$longPath/cbm' # -> Becomes bin = 'path/to/smth/very/deep/somewhere/here/cbm'
```

### [default]
This section is **optional** as well and can be used to set most of `[[Test]]`'s field with a value for the whole TestSuite (except name & description which will then describe the TestSuite).

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
    
[[Test]]
    name = "Test 2"
    desc = "This is useless & stupid too but will fail"
    bin = "ls"
    args = ["-h"]
    refBin = "ls"
```
Following is the explanation of all available fields.

#### Commons
These fields are present directly under a `[default]` or `[[Test]]`, they give basic information about the Test(Suite).
|  Fields |           Usage           |      Type     | Remark |
|:-------:|:-------------------------:|:-------------:|--------|
| name    | Test(Suite)'s name        | string        | **Mandatory** for each `[[Test]]` |
| desc    | Test(Suite)'s description | string        | **Optional** |
| bin     | Binary being tested       | string (path) | **Mandatory**|
| refBin  | Binary used as reference  | string (path) | **Optional** |
| args    | Args passed to `bin`      | [string]      | **Optional** or **ignored** if given as `[default]` |
| refArgs | Args passed to `refBin`   | [string]      | **Optional** or **ignored** if given as `[default]` |

> :question: **Notes on `refBin`**
>
> If `refBin` is provided then it will go through the same process as `bin` and checks will be performed against the outputs of `refBin`; therefore the `[Test.expected]` section of each `[[Test]]` can be fully omitted, if it isn't it will be used for checks rather than the `refBin`'s outputs.
>
> :warning: If the given binaries aren't located in `$PATH` then the path should be provided.
>
> :warning: If the given binaries can't be found or if no `bin` was provided then the test is **ignored**.
> In future updates, if `bin` isn't provided then we shall use the one the backend knows.

#### Expected
These fields are present as a sub section of either `[default]` or `[[Test]]`, such as:
```TOML
[[Test]]
    ...
    [Test.expected]
        ...
```
As the name suggest, this section is used to check if the test succeeded. The fields are the following:
|   Fields   |           Usage        |      Type     | Remark |
|:----------:|:----------------------:|:-------------:|--------|
| status     | Expected exit status   | int           | **optional**, if not provided will expect 0 |
| stdout     | Expected stdout        | string        | **optional**, if provided **do not** use `stdoutFile` |
| stderr     | Expected stderr        | string        | **optional**, if provided **do not** use `stderrFile` |
| stdoutFile | Compare stdout to file | string (path) | **optional**, if provided **do not** use `stdout` |
| stderrFile | Compare stderr to file | string (path) | **optional**, if provided **do not** use `stderr` |

> :question: **Notes on expectations**
>
> If not provided, then the default test is to check if the binary exited without error.
>
> If a `refBin` was provided it's outputs will be used by default.
> Defining status will override the check on `refBin`'s status. The same logic applies for `stdout` and `stderr`.
>
> :warning: If both `stdx` and `stdxFile` are provided, `stdxFile` will be **ignored**

### Interactions
|   Fields   | Usage | Type | Remark |
|:----------:|:-----:|:----:|--------|
| stdoutPipe || string (cmd)   ||
| stderrPipe || string (cmd)   ||
| stdinPipe  || string (cmd)   ||
| stdinFile  || string (path)  ||
| stdin      || string         ||
| pre        || string (cmd)   ||
| post       || string (cmd)   ||
| env        || [string] (kv¹) ||
| addEnv     || [string] (kv¹) ||
> :question: **Notes**
>
> ¹kv: key/value format like `hello=world`
- [x] stdout_pipe : redirect your program stderr to a specified shell command before checking it
- [x] stderr_pipe : redirect your program stderr to a specified shell command before checking it
- [x] stdin : write in the stdin of the process
- [x] stdin_file : write in the stdin of the process from the content of a file
- [x] pre  : Runs a command before the test. Can be used for setup.
- [x] post : Runs a command after the test. Can be used for cleanup.
- [x] env : Define program's env (use default if none provided)
- [x] add_env : change environment variable(s) (append the given value to environment value)

### Options
|   Fields   |                    Usage              |      Type      |    Remark    |
|:----------:|:-------------------------------------:|:--------------:|--------------|
| repeat     | Repeats the `[[Test]]`                | int            | **optional** |
| timeout    | Kills the `[[Test]]` if reached       | float(seconds) | **optional** |
| time¹      | Times the execution of the `[[Test]]` | bool           | **optional** |
| shouldFail | `[[Test]]` succeed if it fails        | bool           | **optional** |
> :question: **Notes**
>
> ¹time: for more precise results the `[[Test]]` should be as bare as possible.
## Predicted

Usage of CBM's backend to be able to omit the `bin` field & add new fields such as `build`
