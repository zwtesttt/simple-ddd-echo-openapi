
MODEL_LIST = user line node

.PHONY: gen-openapi
gen-openapi: $(addprefix gen-openapi-, $(MODEL_LIST))

gen-openapi-%:
	-@oapi-codegen -generate echo,spec -o internal/api/http/v1/$*/openapi.gen.go -package v1 internal/api/http/v1/$*/$*.yaml
	-@oapi-codegen -generate types -o internal/api/http/v1/$*/openapi.types.go -package v1 internal/api/http/v1/$*/$*.yaml


# 定义变量
INPUT_DIRS := $(wildcard  ./internal/api/http/v1/*)
OUTPUT_FILE := ./docs/cosslan.yaml
SWAGGER_MERGER := $(shell command -v go-swagger-merger)

#go get github.com/efureev/go-swagger-merger
#go install github.com/efureev/go-swagger-merger
# 定义合并 OpenAPI 的目标
.PHONY: merge-openapi
merge-openapi:
    # 检查是否存在 go-swagger-merger 命令
	go-swagger-merger -o $(OUTPUT_FILE) $(foreach dir,$(INPUT_DIRS),-i $(dir)/*.yaml)