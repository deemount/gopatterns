# Go, NodeJS & gRPC

## Voraussetzungen

* go muss installiert sein
* protoc muss installiert sein
* protoc-gen-go muss installiert sein (Code Generierung)
* protoc-gen-go-grpc muss installiert sein (Code Generierung)
* goctl kann installiert sein
* nodejs muss installiert sein

## Verweise

## Terminal

### Prüfen, ob Voraussetzungen gegeben sind

```bash
goctl env check -i -f -v
```

Der oben aufgeführte Befehl muss diese Aufgabe für eine erfolgreiche Ausführung aufweisen

```bash
[goctl-env]: preparing to check env

[goctl-env]: looking up "protoc"
[goctl-env]: "protoc" is installed

[goctl-env]: looking up "protoc-gen-go"
[goctl-env]: "protoc-gen-go" is installed

[goctl-env]: looking up "protoc-gen-go-grpc"
[goctl-env]: "protoc-gen-go-grpc" is installed

[goctl-env]: congratulations! your goctl environment is ready!
```

### Befehl zur Generierung er *.pb.go-Dateien

```bash
cd rpc-server
protoc --go-grpc_out=require_unimplemented_servers=false:. --go_out=. chat.proto
```

### Befehl zur Erstellung des RPC Clienten

```bash
cd rpc-client
npm init
npm i node-grpc-client
```

## Anwendung starten

Zwei Terminalfenster öffnen

```bash
cd rpc-server
go run server.go
```

```bash
cd rpc-client
node client.js
```