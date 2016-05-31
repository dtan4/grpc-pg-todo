# TODO application sample w/ gRPC + PostgreSQL

## Protocol Buffers

On Mac:

```bash
$ brew install protobuf --devel
```

## PostgreSQL

Run PostgreSQL Docker container:

```bash
$ make run-pg
$ psql postgres://postgres:password@localhost:5432
```

Create `grpc_pg_todo` database:

```sql
CREATE DATABASE grpc_pg_todo;
```

## Golang

[grpc / gRPC Basics - Go](http://www.grpc.io/docs/tutorials/basic/go.html)

Install gRPC tools:

```bash
$ go get google.golang.org/grpc
$ go get -a github.com/golang/protobuf/protoc-gen-go
```

Build:

```bash
$ make deps
$ make generate-go
$ make build
```

Start server & client:

```bash
$ cd go
$ bin/server &
$ bin/client
```

## Ruby

[grpc / gRPC Basics - Ruby](http://www.grpc.io/docs/tutorials/basic/ruby.html)

Install gRPC tools:

```bash
$ gem install grpc
$ git clone https://github.com/grpc/grpc.git $REPO
$ cd $REPO
$ make grpc_ruby_plugin
$ cp bins/opt/grpc_ruby_plugin /usr/local/bin/
```

Build:

```bash
$ make generate-ruby
$ cd ruby
$ bundle install
```

Start server & client:

```bash
$ bundle exec bin/server &
$ bundle exec bin/client
```
