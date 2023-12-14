# project-a

## Build & Run

Using `Taskfile.yml` <https://taskfile.dev/>

> Build

```sh
❯ task build
task: [build] go build -o bin/project-a main.go
```

> Run dev

```sh
❯ task run-dev
task: [run-dev] go run main.go
```

> Run with binary

```sh
❯ task run
task: [build] go build -o bin/project-a main.go
task: [run] bin/project-a
```

## API Calls

> Query

```sh
❯ curl http://127.0.0.1:3001/api/version | jq
```

> Answer

```json
{ "version": "16.7.0-pre" }
```
