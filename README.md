# App run

default run on :8080
```shell
go run ./cmd/sysinfo 
```

run on specific port and print in console sysInfo
```shell
go run ./cmd/sysinfo -port :9090 -verbose
```

# App build and run 

```shell
chmod +x build.sh && ./build.sh
```