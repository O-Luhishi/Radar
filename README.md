# Radar : A Restful API Written In Go for Port Scanning

## Overview

Small project in Go, for creating a port scanner using Rest. A potential use case is dropping it on a server 
(EC2/Droplet etc) and allow it to scan your cloud infrastructure. 

## Install
```
go get github.com/O-Luhishi/Radar
go build
./Radar
```
## Endpoints
```
METHOD: GET
/api/healthcheck/
```

````
METHOD: GET
/api/scan/
````

````
METHOD: POST
/api/scan/{ip_addr}/
````



Big thanks to [@Anvie](https://github.com/anvie) for the Port Scanner Go Module!


