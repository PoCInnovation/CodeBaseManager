# CodeBaseManager/cli - Repository

The goal of this module is provide helper regarding the organization of the repository (as well as project).

## Overview

This section will give a brief overview of each command's capabilities.
For more informations, please see [] :warning: TODO: add link to details

### Create

The goal of `create` is to quickly clone a project based on a template, such as:

```bash
    cbm repository create <github link>

    Example:
    cbm repository create https://github.com/PoCInnovation/CodeBaseManager/
```

If the repository contains a `.cbm/repository.toml` file, additional steps to further initialize
the repository based on the given configuration will occur.

#### Predicted

    [x] Adds repo to local API
    [ ] `--path` option to tell where to clone
    [ ] support for more VCS

### Todos

`todos` is inspired by JetBrains' TODOs. Here's the command to enter:

```bash
    cbm repository todos
```

This will display all todos present in the working directory recursively, giving information about
the files & lines where each todo is located as well as content.

#### Predicted

    [x] Save todos in local API
    [ ] :warning: TODO: provide way to handle todos directly via cli (interaction with local API)

