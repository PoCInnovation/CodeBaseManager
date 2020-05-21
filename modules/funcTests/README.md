# CodeBaseManager - Functional Tests

This module is highly inspired by [JenRik](https://github.com/Yohannfra/JenRik).

The goal of this module is too allow the user to describe tests for their binaries through simple [TOML](https://github.com/toml-lang/toml) files.

## Quick Start
WIP...

## Details
Each tests file represent a TestSuite. A test file has up to 3 of the following sections:

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
    args = "-h"
    refBin = "ls"
```
Following is the explanation of all available fields.

#### Commons
These fields are present directly under a `[default]` or `[[Test]]`, they give basic information about the Test(Suite).
|  Fields |           Usage           | Remark |
|:-------:|:-------------------------:|-----------------------------------------------------|
| name    | Test(Suite)'s name        | **Mandatory** for each `[[Test]]`                   |
| desc    | Test(Suite)'s description | **Optional**                                        |
| bin     | Binary being tested       | `[[Test]]` with missing `bin` are **ignored**       |
| refBin  | Binary used as reference  | **Optional**                                        |
| args    | Args passed to `bin`      | **Optional** or **ignored** if given as `[default]` |
| refArgs | Args passed to `refBin`   | **Optional** or **ignored** if given as `[default]` |

> **Reference explained**:
> If a binary is given as a reference then it will go through the same process as our binary and checks will be performed against the reference; therefore the `[Test.expected]` section of each tests can be fully omitted, if it isn't it will be used for checks rather than the reference.

#### Expected
These fields are present as a sub section of either `[default]` or `[[Test]]`, such as:
```TOML
[[Test]]
    ...
    [Test.expected]
        ...
```
As the name suggest, this section is used to check if the test succeeded. The fields are the following:
|   Fields   |           Usage           | Remark |
|:----------:|:-------------------------:|-------------------------------------------------------|
| status     | Expected exit status      | **optional**, if not provided will expect 0.          |
| stdout     | Expected stdout (string)  | **optional**, if provided **do not** use `stdoutFile` |
| stderr     | Expected stderr (string)  | **optional**, if provided **do not** use `stderrFile` |
| stdoutFile | Compare stdout to file    | **optional**, if provided **do not** use `stdout` |
| stderrFile | Compare stderr to file    | **optional**, if provided **do not** use `stderr` |

> **Notes on expectations**: If not provided, then the default test is to check if the binary exited without error.
>
> If a `refBin` was provided it's outputs will be used by default.
> Defining status will override the check on `refBin`'s status. The same logic applies for `stdout` and `stderr`.
>
> If both `stdx` and `stdxFile` are provided, `stdxFile` will be **ignored**

- [x] stdout_pipe : redirect your program stderr to a specified shell command before checking it
- [x] stderr_pipe : redirect your program stderr to a specified shell command before checking it
- [x] stdin : write in the stdin of the process
- [x] stdin_file : write in the stdin of the process from the content of a file
- [x] pre  : Runs a command before the test. Can be used for setup.
- [x] post : Runs a command after the test. Can be used for cleanup.
- [x] env : Define program's env (use default if none provided)
- [x] add_env : change environment variable(s) (append the given value to environment value)

Sprint 2
- [ ] repeat : repeat the test x times
- [ ] time : Time the execution of your program
- [ ] timeout : make the test fail if it times out, after killing it (SIGTERM) (the time is given in seconds)
- [ ] should_fail : make the test success if it fails


Later Sprint
- [ ] build_with : 
- [ ] noBin (gotta look within CBM)

## Predicted

- Build
