# Exec

## install grpc_cli

```shell
$ brew tap grpc/grpc

$ brew install grpc
```

## run application

```shell
$ docker-compose up
```

# Call gRPC

## service list

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