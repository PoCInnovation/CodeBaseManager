# CodeBaseManager - Functional Tests

This module is highly inspired of [JenRik](https://github.com/Yohannfra/JenRik).
It can interact with CBM's other modules (such as build) if required.

## Possible Commands

- [x] bin
- [x] refbin -> More or less implemented

- [x] arg
- [x] refArg

- [x] stdout : expected stdout of your program
- [x] stdout_file : compare your program stdout with the content of a given file
- [x] stdout_pipe : redirect your program stderr to a specified shell command before checking it

- [x] stderr : expected stderr of your program
- [x] stderr_file : compare your program stderr with the content of a given file
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
