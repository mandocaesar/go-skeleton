PROTOFILES := $(shell find . -name "*.proto" -type f)

generate-proto:
	@echo "Generating protobuf"; \
	for FILE in $(PROTOFILES); do \
		echo "processing" $$FILE; \
		protoc --go_out=grpc:. $$FILE;\
	done;

add-package:
	@echo "Adding Package";\
	dep ensure -add $(name);\
