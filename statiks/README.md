# Statiks Directory 

This directory contains all static files to be embedded to the `bs` tool.

Each directory is in its own namespace and golang package, such that it will be neatly separated.

List of current embedded dirs:

| Directory        | Namespace | Go Package Name |
|------------------|-----------|-----------------|
| boilerplate      | bproot    | bproot          |
| boilerplate/core | bpcore    | bpcore          |
| boilerplate/lyft | bplyft    | bplyft          |
| boilerplate/tool | bptool    | bptool          |

