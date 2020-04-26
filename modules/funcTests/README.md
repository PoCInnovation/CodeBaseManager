# CodeBaseManager - Functional Tests

This module is highly inspired of [JenRik](https://github.com/Yohannfra/JenRik).
It can interact with CBM's other modules (such as build) if required.

## Possible Commands

- [ ] arg

- [ ] stdout : expected stdout of your program
- [ ] stdout_file : compare your program stdout with the content of a given file
- [ ] stdout_pipe : redirect your program stderr to a specified shell command before checking it

- [ ] stderr : expected stderr of your program
- [ ] stderr_file : compare your program stderr with the content of a given file
- [ ] stderr_pipe : redirect your program stderr to a specified shell command before checking it

- [ ] stdin : write in the stdin of the process
- [ ] stdin_file : write in the stdin of the process from the content of a file

- [ ] repeat : repeat the test x times
- [ ] time : Time the execution of your program
- [ ] timeout : make the test fail if it times out, after killing it (SIGTERM) (the time is given in seconds)
- [ ] should_fail : make the test success if it fails

- [ ] pre : run a shell command before executing the test
- [ ] post : run a shell command after executing the test

- [ ] env : change environment variable(s) (replace the value with the given one)
- [ ] add_env : change environment variable(s) (append the given value to environment value)

- [ ] build_with : 
