ALL_PROTO_FILES=$(shell find api -name '*.proto')

.PHONY: api
# generate grpc code
api:
	protoc --proto_path=. \
  		   --proto_path=./third \
 		   --go_out=paths=source_relative:. \
           --go-grpc_out=paths=source_relative:. \
           --go-errors_out=paths=source_relative:. \
           $(ALL_PROTO_FILES)

format:
	 @ for i in $(ALL_PROTO_FILES); \
 		do \
 		  clang-format $$i --style='{BasedOnStyle: Google, IndentWidth: 4, ColumnLimit: 0, AlignConsecutiveAssignments: true, AlignConsecutiveAssignments: true}'>bak.proto; \
 		  mv bak.proto $$i;\
 		done

