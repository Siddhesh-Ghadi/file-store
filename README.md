# File Store

## Setup & Usage

### Server

Build the server using below instructions or use the [container image](https://hub.docker.com/r/sghadi1203/file-store/tags) to start the server.

```bash
docker run --name file-store-server -p 8080:8080 sghadi1203/file-store:latest
```

### Client

Build the client using below instructions or use the binary from [github release page](https://github.com/Siddhesh-Ghadi/file-store/releases) of this repo.

```bash
# set server address for client
$ export FILE_STORE_SERVER="http://localhost:8080"

$ store help
Help
ls                      List all files from store.
add file1 file2...      Add files to store store.
update file1 file2...   Update files in store store.
rm file1 file2...       Remove files from store store.
freq-words --limit 10|-n 10 --order dsc|asc     List count of each word. limit & order are optional flags.
```

## Build Instructions

#### Server Build

```bash
go build ./cmd/server/server.go
```

This will produce `server` binary. 

#### Client Build

```bash
go build ./cmd/client/store.go
```

This will produce `store` client binary. 