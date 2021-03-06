# Word Count Map Reducer based on gRPC
A distributed map reduce program to do word count with gRPC (https://grpc.io) on local machine.

## How to run:
you can start driver program, and then open a few more terminal windows, and start more
worker program, when it finished, you should see output files containing the frequency of each work in the input
files.
### Run driver
```bigquery
go run driver.go
```
### Run worker
```bigquery
go run worker.go mapper_reducer.go
```
## Set protocol buffers
I used protocol buffers for ManageTaskPool microservice

```bigquery
 go get -u github.com/golang/protobuf/protoc-gen-go
```
```bigquery
 protoc --go_out=paths=source_relative:worker_driver -I. worker_driver.proto
```
```bigquery
 protoc --go-grpc_out=paths=source_relative:worker_driver -I. worker_driver.proto
```

## Assumptions
- Both worker and driver access to same file system (they have same path to files).
- Inputs are always in a directory called `inputs`
- Intermediate files, which are results of map tasks are in `intermediates` directory.
- Output files, are in `outputs` directory
- Division of tasks is based on *Count of files*, therefore number of map tasks will be ***minimum of given number of map tasks and given number of files***
- Default number of map tasks is 6
- Default number of reduce tasks is 4

### Note
You can change these settings in `worker_driver/config.go` file.
## Architecture
A map reducer system, where each map or reduce task is done using map reduce methodology

- step1: Driver separates whole task of word counting into some map tasks and reduce tasks to speed up and use available separate computers in different locations.
- step2: A specific map/reduce task is assigned to one and only one worker.
- step3: Worker process the map task using *map reduce methodology* to speed up and use all 8 cores of its system.
- step4: Worker process the reduce task using *map reduce methodology* to speed up and use all 8 cores of its system.

![gRPC based word counter Architecture](https://github.com/mahsirat-atiye/map-reduce-grpc/blob/master/documents/grpc_wc.png)


## Protocol and microservices
You can find about protocols and end points in *ManageTaskPool microservice* in the following diagram.
![Workflow diagram](https://github.com/mahsirat-atiye/map-reduce-grpc/blob/master/documents/grpc_wc_digram.png)
Symbols follow BPM standards.