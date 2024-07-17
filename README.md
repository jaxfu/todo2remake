# Todo 2.0

## Overview

Allows users to add and remove items from a todo list.
Currently login/register does not work as I have not implemented database functionality yet.

## Dependencies

- NPM
- Golang

## Setup Instructions

1. Clone GitHub repo

```bash
git clone https://github.com/Furrj/todo2remake.git
```

2. Enter client directory

```bash
cd todo2remake/client
```

3. Install NPM packages and build client

```bash
npm i && npm run build
```

4. Enter server directory

```bash
cd ../server
```

5. Install go packages and build server

```bash
go mod tidy && go build -o server.exe cmd/server.go
```

6. Run project

```bash
./server.exe
```

7. Navigate to http://localhost:5000/ in a browser to view/use the application
