set shell := ["zsh", "-uc"]

build:
    #!/usr/bin/env sh
    mkdir -p publish
    cd publish
    go build -buildmode=c-shared -o libtdl_ffi.dylib ../main.go

gen_grpc:
    #!/usr/bin/env sh
    cd pb
    protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative tdl.proto

clean:
    rm -rf publish