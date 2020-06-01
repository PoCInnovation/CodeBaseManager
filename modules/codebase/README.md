# CodeBaseManager - Codebase (Repository Management)

The goal of this module is to allow the user to describe repository for their Project through simple [TOML](https://github.com/toml-lang/toml) files.

Furthermore, Basic Cat and Find Tools (for files and functions) are implemented.

## Repository
In .cbm folder, there is a toml file called `repository.toml`

This file allow you to configure some features such as:
```TOML
Language = []
Modules = []
Tests = []
``` 

### Details
Each line represent a feature that can be specified:

These fields are mandatory for now.

|  Fields   |           Usage               |      Type         | Remark        |
|:---------:|:-----------------------------:|:-----------------:|---------------|
| Language  | Project language list         | Array of string   | **Mandatory** |
| Modules   | Project sources dependencies  | Array of string   | **Mandatory** |
| Tests     | Folders that contains tests   | Array of string   | **Mandatory** |


## Codebase Tools

Theses tools allow the user to quickly find some information.

| Fields    | Usage                         | Description                               |
|:---------:|:-----------------------------:|:-----------------------------------------:|
| Find      | cbm codebase find `[args]`    | Find requested files or functions         |
| Cat       | cbm codebase cat `[args]`     | Display requested files or functions      |


| Language      | Remark                                        |
|:-------------:|:---------------------------------------------:|
| C             | Normal, static and static inlines functions   |
| Go            | Functions                                     |
| Python        | Functions and methods, with decorators        |
