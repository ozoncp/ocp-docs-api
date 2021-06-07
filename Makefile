.PHONY: setup
build: vendor-proto .generate .build
PHONY: .generate
.generate:
	 IF NOT EXIST "pkg\ocp-runner-api" \
	 mkdir "pkg\ocp-docs-api" &&\
	 protoc -I vendor.protogen \
		   --go_out=pkg/ocp-docs-api --go_opt=paths=import \
		   --go-grpc_out=pkg/ocp-docs-api --go-grpc_opt=paths=import \
		   --grpc-gateway_out=pkg/ocp-docs-api \
		   --grpc-gateway_opt=logtostderr=true \
		   --grpc-gateway_opt=paths=import \
		   --validate_out lang=go:pkg/ocp-docs-api \
		   --swagger_out=allow_merge=true,merge_file_name=api:. \
		   api/ocp-docs-api/ocp-docs-api.proto &&\
	 move "pkg\ocp-docs-api\github.com\ozoncp\ocp-docs-api\pkg\ocp-docs-api\*" "pkg\ocp-docs-api" &&\
	 rmdir /s "pkg\ocp-docs-api\github.com" &&\
	 IF NOT EXIST "cmd/ocp-docs-api" \
		mkdir "cmd/ocp-docs-api"
PHONY: .build
.build:
	go build -o bin/ocp-docs-api.exe cmd/ocp-docs-api/main.go
PHONY: install
install: build .install
PHONY: .install
install:
	go install cmd/grpc-server/main.go
PHONY: vendor-proto
vendor-proto: .vendor-proto
PHONY: .vendor-proto
.vendor-proto:
	mkdir "vendor.protogen/api/"