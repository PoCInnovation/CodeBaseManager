# CodeBaseManager

## Description
CodeBaseManager is a command line tool designed to ease the development of projects through multiple, single purpose, language independant, modules.

It will have a backend to hold informations about the codebase, build system and such, of the project for the modules to use.

## Installation

:warning: Requires go 1.14
```
 git clone https://github.com/PoCFrance/CodeBaseManager
 cd CodeBaseManager
 ./install.sh
```
## Quick Start

```
cbm module [command]
# or
cbm help
```

If no command is provided then, depending on the module, you may enter a shell. This shell can execute a command at a time and the module's commands.

## Modules
A module fills up a specific task. In the future they will be able to interact with the backend and with each other (for specific features).

### Current

|        Modules                             | Description | Aliases | Shell |
|:------------------------------------------:|:-----------:|:-------:|:-----:|
|[server](/modules/server/README.md)         | Allows you to interact with CBM's backend | none    | no |
|[repository](/modules/repository/README.md) | Provides helpers for the organization of the repository (as well as project) |none    | no |
|[codebase](/backend/codebase/README.md)     | Allows to find informations about the codebase  |none    | yes |
|[functional-tests](/modules/funcTests/README.md) | Allows to test a binary through the use of [TOML](https://github.com/toml-lang/toml) files to describe tests |"ft" | yes |

### Predicted

| Modules    | Description |
|:----------:|:-----------:|
| unit-tests | Allows to write unit tests for your functions |
| build      | Allows to describe a build system in a unique way |
| debug      | Allows to debug your program or functions |


## Dependencies

|                          Dependency                        |      License       |
|:----------------------------------------------------------:|:------------------:|
| [spf13/Cobra](https://github.com/spf13/cobra)              | Apache License 2.0 |
| [BurntSushi/toml](https://github.com/BurntSushi/toml)      | MIT License        |
| [gin-gonic/gin](https://github.com/gin-gonic/gin)          | MIT License        |
| [logrusorgu/aurora](https://github.com/logrusorgru/aurora) | Unlicense License  |


------------
## Maintainers

 - [Allan Debeve](https://github.com/Gfaim)
 - [Quentin Veyrenc](https://github.com/VrncQuentin)
 - [Damien Bernard](https://github.com/Encorpluptit)
