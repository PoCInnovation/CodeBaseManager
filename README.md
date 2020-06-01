# CodeBaseManager

## Description
CodeBaseManager is a command line tool designed to ease the development of projects through multiple, single purpose, language independant, modules.

It will have a backend inspired by **[Unison project](https://www.unisonweb.org/)** to hold informations about the codebase, build system and such, of the project for the modules to use.

## Build

:warning: Require go 1.14
```
 $ go build main.go -o cbm
 $ ./cbm [command]
```
|Available commands | Description |
|-------------------|------------|
  |codebase         | Simple shell to navigate through your codebase. |
  |functional-tests | Helps you deal with your functional tests. |
  |help             | Help about any command |

## Modules
A module fills up a specific task. In the future they will be able to interact with the backend and with each other (for specific features).

### Current

- [CodeBase](): Allows to find informations about the codebase.

- [Functional Tests](): Allows to test a binary through the use of [TOML](https://github.com/toml-lang/toml) files to describe tests.

### Predicted

- Unit Tests: Allows to write unit tests for your functions.
- Build: Allows to describe a build system in a unique way.
- Debug: Allows to debug your program or functions.


------------
## Maintainers

 - Allan Debeve ([GitHub](https://github.com/Gfaim))
 - Quentin Veyrenc ([GitHub](https://github.com/VrncQuentin))
 - Damien Bernard ([GitHub](https://github.com/Encorpluptit))
