# CodeBaseManager - Backend

This module is a Dockerized API that allow user to stock information about project added by users.

## Overview

This API use [go-sqlite3](https://github.com/mattn/go-sqlite3) and [Gin](https://github.com/gin-gonic/gin).

It receives requests by CBM-Watcher to keep track of changes in watched Repository.



### Start & Stop

This module is being launch directly during installation, on a Port chosen by the user.

It is not yet started with a service via systemctl, but the user can launch it manually with CLI:

```bash
    cbm server start
    cbm server stop
```

### Tables

---
#### Project
|  Fields   | Description                       | Type                  | Remark                    |
|:---------:|:---------------------------------:|:---------------------:|:-------------------------:|
| Name      | Project name (Base Directory)     | String                | **Mandatory/Unique**      |
| Path      | Project Path (Given by Watcher)   | String                | **Mandatory/Unique**      |
| Modules   | Project Modules                   | List of Module        |                           |

---
#### Module
|  Fields   | Description                       | Type                  | Remark                    |
|:---------:|:---------------------------------:|:---------------------:|:-------------------------:|
| Name      | Module name (Base Directory)      | String                | **Mandatory/Unique**      |
| Path      | Module Path (Given by Watcher)    | String                | **Mandatory/Unique**      |
| Functions | Module's Functions                | List of Functions     |                           |

---
#### Function
|  Fields   | Description                       | Type                  | Remark                    |
|:---------:|:---------------------------------:|:---------------------:|:-------------------------:|
| Name      | Project name (Base Directory)     | String                | **Mandatory/Unique**      |
| Path      | Project Path (Given by Watcher)   | String                | **Mandatory/Unique**      |


## Details
