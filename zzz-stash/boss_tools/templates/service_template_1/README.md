# boss_pod_template

boss_pod_template is a project template for web service pod and utility cli  

## WORK IN PROGRESS

### based on code from "Go microservice template for Kubernetes" https://github.com/stefanprodan/podinfo

## Use *minimal* branch for lightweight version

## How to get started

```

cd /tmp
git clone git@github.com:braincorp/boss_pod_template.git --branch minimal

cd /tmp/boss_pod_template

make build

mkdir ~/myservice
cd ~/myservice
boss_pod_template --clone

```

last line will create new service tree under my_service with all "boss_pod_template" replaced with "my_service".

```
myservice$ tree
.
├── LICENSE
├── Makefile
├── README.md
├── cmd
│   ├── clone.go
│   └── main.go
├── go.mod
├── go.sum
└── pkg
    ├── api
    │   ├── http.go
    │   ├── logging.go
    │   ├── mock.go
    │   ├── server.go
    │   ├── version.go
    │   └── version_test.go
    ├── config
    │   ├── config.go
    │   └── config_test.go
    └── version
        └── version.go

```
Then you can try
```
make help
make test
make run

http :8080/version 

or 

curl localhost:8080/version

```

###Specifications (see master branch for all available features)

* Dependency Injection
* Graceful shutdown on interrupt signals
* Instrumented with Prometheus
* Authentication with JWT token
* Structured logging with zap 

see make help for make targets

