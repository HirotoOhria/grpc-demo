<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [gRPC Demo](#grpc-demo)
    - [Directory structure](#directory-structure)
- [Preconditions](#preconditions)
    - [Install grpc_cli](#install-grpc_cli)
- [Launch Application](#launch-application)
    - [run application](#run-application)
- [Call gRPC](#call-grpc)
  - [check service list](#check-service-list)
  - [call Bake](#call-bake)
- [Development](#development)
    - [Generate codes](#generate-codes)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# gRPC Demo

This is [スターティングgRPC](https://www.amazon.co.jp/dp/B087R87L6Z) based sample codes.

### Directory structure

- example
    - codes of Chapter 2
- qiita
    - codes of [qiita example](https://qiita.com/drafts/a4e06a3e7e8c8dfef4df) 
- src
    - codes of Chapter 3 

# Preconditions

### Install grpc_cli

```shell
$ brew tap grpc/grpc

$ brew install grpc
```

# Launch Application

### run application

```shell
$ docker-compose up
```

# Call gRPC

## check service list

```shell
$ grpc_cli ls localhost:50051 pancake.maker.PancakeBakerService -l
filename: pancake.proto
package: pancake.maker;
service PancakeBakerService {
  rpc Bake(pancake.maker.BakeRequest) returns (pancake.maker.BakeResponse) {}
  rpc Report(pancake.maker.ReportRequest) returns (pancake.maker.ReportResponse) {}
}
```

## call Bake

```shell
$ grpc_cli call localhost:50051 pancake.maker.PancakeBakerService.Bake 'menu: 1'
connecting to localhost:50051
pancake {
  chef_name: "ohira"
  menu: CLASSIC
  technical_score: 0.734285235
  create_time {
    seconds: 1648361396
    nanos: 583884176
  }
}
Rpc succeeded with OK status
```

```shell
$ grpc_cli call localhost:50051 pancake.maker.PancakeBakerService.Bake 'menu: 0'
connecting to localhost:50051
Rpc failed with status code 3, error message: パンケーキを選択してください
```

# Development

### Generate codes

```shell
$ docker-compose exec proto bash

$ protoc --go_out=api --go-gropc_out=api --go-grpc_opt=require_unimplemented_servers=false proto/*.proto
```