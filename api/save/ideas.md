# CodeBaseManager - Codebase (Repository Management)

Small Api

db pros and cons

## Tables:

Project


---
Langages (to difference way to interpret / compile)

|       Fields          |                   Description                     |               Type                |               Remark                      |
|:---------------------:|:-------------------------------------------------:|:---------------------------------:|-------------------------------------------|
| Name                  | Name of module                                    | String                            |                                           |
| Modules               | ID of module                                      | One to Many relation              |                                           |
| Name                  | Name of module                                    | String                            |                                           |

```json
{
  "name": "",
  "name": ""
}
```

---
Modules (group of things (folder/ packages / classes, types))

|       Fields          |                   Description                     |               Type                |               Remark                      |
|:---------------------:|:-------------------------------------------------:|:---------------------------------:|-------------------------------------------|
| Name                  | Name of module                                    | String                            | Generally folder name                     |
| Path                  | Path of module                                    | String                            | defined in codebase yaml file             |

```json
{
  "name": "",
  "path": ""
}
```

---
Types

|       Fields          |                   Description                     |               Type                |               Remark                      |
|:---------------------:|:-------------------------------------------------:|:---------------------------------:|-------------------------------------------|
| Path                  | Path of the file to find type                     | Array of string                   | Use file extension to find way to include |
| Types ?               | Path of the file to find type                     | Array of string                   | Use file extension to find way to include |

```json
{
  "path": "",
  "name" : "",
}
```

Functions
```json

```

