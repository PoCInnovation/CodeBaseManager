# CodeBaseManager/cli - Server

This module is one of the most important as it allows to directly interact with CBM's server

## Overview

This section will give a brief overview of each command's capabilities.
For more informations, please see [] :warning: TODO: add link to details

### Start & Stop

These two commands are pretty self explanatory, so here's a quick exemple of both:

```bash
    cbm server start
    cbm server stop
```

### Add & Drop

These two commands allow you to respectively drop or add a repository to CBM's watch-list.

```bash
    cbm server add 'path/to/repository'
    cbm server drop 'repository-name'
```

### List

This command will list some information's what CBM's server.

By default, if no arguments are provided, it will display CBM's watch-list
(and is currently this only behavior of this command).

```bash
    cbm server list
```

## Details
