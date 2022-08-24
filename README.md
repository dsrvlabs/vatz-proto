# Vatz Proto
> This is a Proto Buffer for DSRV's Vatz Node management system.

### Here's simple path 
```
.
├── LICENSE
├── Makefile
├── README.md
├── manager
│   └── v1
│       ├── manager.pb.go
│       └── manager_grpc.pb.go
├── plugin
│   └── v1
│       ├── plugin.pb.go
│       └── plugin_grpc.pb.go
└── proto
    └── vatz
        ├── manager
        │   └── v1
        │       └── manager.proto
        └── plugin
            └── v1
                └── plugin.proto

```

All proto files go to under path proto folder and 
compiled proto files would locate at default-path with their service name as above 
Put any change in shared [excel](https://docs.google.com/spreadsheets/d/1Hq4JwY0Ys9krHcIpcYyoQqi6_IzHl2inTFsdFekGtDA/edit#gid=0) first if there's any update with Vatz-Proto


---

## Configuration for private repo access

Ref1: https://www.digitalocean.com/community/tutorials/how-to-use-a-private-go-module-in-your-own-project
Basic guide.

Ref2: https://stackoverflow.com/questions/69682030/how-to-go-get-private-repos-using-ssh-key-auth-with-password
We have to change the way go get invokes ssh, by disabling batch mode.

```
$ export GOPRIVATE=github.com/hqueue/vatz-secret
$ env GIT_SSH_COMMAND="ssh -o ControlMaster=no -o BatchMode=no" go get github.com/hqueue/vatz-secret
$ make
```

## Submodule
TBD

grpc-gateway.
googleapis by submodule.

