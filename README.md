# CodeBaseManager

## Description
CodeBaseManager is a command line tool designed to ease the development of projects through multiple, single purpose, language independant, modules.

It will have a backend inspired by **[Unison project](https://www.unisonweb.org/)** to hold informations about the codebase, build system and such, of the project for the modules to use.

## Installation

:warning: Requires go 1.14
```
 $ git clone https://github.com/PoCFrance/CodeBaseManager
 $ cd CodeBaseManager
 $ ./install.sh
```
## Quick Start

```
$ cbm module [command]
or
$ cbm help
```

If no command is provided then, depending on the module, you may enter a shell. This shell can execute a command at a time and the module's commands.

|        Modules    | Aliases | Shell |
|-------------------|:-------:|:-----:|
|codebase           | none    | yes |
|functional-tests   | "ft" | yes |

## Modules
A module fills up a specific task. In the future they will be able to interact with the backend and with each other (for specific features).

### Current

- [Server](/modules/server/README.md): Allows you to interact with CBM's backend.

- [Repository](/modules/repository/README.md): Provides helpers for the organization of the repository (as well as project).

- [CodeBase](/modules/codebase/README.md): Allows to find informations about the codebase. The available commands are `cat` & `find`, they can target functions or files.

- [Functional Tests](/modules/funcTests/README.md): Allows to test a binary through the use of [TOML](https://github.com/toml-lang/toml) files to describe tests. It has one command: `run`

### Predicted

- Unit Tests: Allows to write unit tests for your functions.
- Build: Allows to describe a build system in a unique way.
- Debug: Allows to debug your program or functions.
- More commands for CodeBase & Functional Tests.


------------
## Maintainers

 - Allan Debeve ([GitHub](https://github.com/Gfaim))
 - Quentin Veyrenc ([GitHub](https://github.com/VrncQuentin))
 - Damien Bernard ([GitHub](https://github.com/Encorpluptit))
