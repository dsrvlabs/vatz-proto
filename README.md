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