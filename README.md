# Grid - a decentralized version control system

## Technology

- golang
- Linux-like file system
- download & upload in HTTP
- compared with git：
  - tech: grid runs as microservices
  - decentralized: any local Grid can become server anytime anywhere
  - easy: more explicit in usage and reminders
  - extensible: can use extensions to see the change between different submit history

## Key Features

     server runs grid. Clients also run `grid`. It is a decentralized deployment. Server is nothing different from clients (they all run grid). Since they all run grid, they can be seen as a **grid node** on the internet. So you can actually synchronize things to everywhere(not just Gridle server). Everyone can be either Grid server/ grid client.

the only thing that server is different from client, is server runs Gridle Service but clients don’t

---

### Download

```bash
grid -d url
grid download url

#if download before
grid sync
```

### Upload

```bash
grid add ./hello.cpp
grid add all #if you are lazy
grid submit "title" "description"
grid sync
```

imagine that there is a local server for each project they download

after using `grid sync`, local server will send a request to synchronize to the server

Every single submit has its own ID stored in file system to be identified

### History

```bash
grid trace <historyID> # in this step, local server will just trace back into the history
grid sync # this step is irrversible. cuz it will affect to cloud server directly
```

### Connect to the Server (other grid nodes)

uses HTTP

```bash
grid connect <ip address>
grid test # to test if the sync runs well
grid chat "test" #talk to server
```

### Branch

```bash
grid branch create "name"

grid branch delete "name" #this needs 2-step authentication

grid branch "main" #switch to the branch

```

### Conflict (when sync)

- if there are already new changes in the project, but it will not be conflicted with the submit that I am syncing with, it can just choose to **`merge`**

```bash
$ grid sync
- sync starts!
- Fatal Error: conflict between two submits
- but you can choose to use "grid merge"

$ grid merge
- sync starts!
- [=================] 100%
- Succeeded.

```

- also can choose to restore files to become the status before it changed

### Decentralized Storage of Projects

     Since everything in Grid is decentralized, so does the storage. Grid is integrated a go server. We can see Grid as microservices which communicates with each other by http

### Version

     You can actually divide the project into multiple stages (like different versions). Every single version in Grid has its own `release` and `code`
    
     Version is actually the special version of branch. But we can choose a branch as a specific version. And we can also add release into version like how we did in branch

It is also available to use `grid branch` to switch our workflow into specific version

```bash
$ grid version create
name: Gridle_1.0
version:  1.0.1
description: It is an official version of Gridle

$ grid branch "Gridle_1.0"
$ grid version "Gridle_1.0"

$ grid version update 1.1.2 #updating
$ grid version history #to see the history of updating in version
$ grid history #it is different from the last command. grid history is to see the history of submits in this branch
$ grid submit history 1213 #This is used to see the exact detail of selected submit.

$ grid sync release #it will automatically update "release" folder into the release of this version
```


### Hash Used in Grid

    Grid will allocate everything with an identification: `SHA-1` value.
    
    Every file in Grid will be along with a `hash` value, which can be used to see whether the files have been changed by comparing their values.
