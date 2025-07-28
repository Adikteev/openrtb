default: test

DEPENDENCIES_FOLDER := .deps
MSGP_BIN := msgp
MSGP_VERSION := v1.3.0
MSGP := $(DEPENDENCIES_FOLDER)/$(MSGP_BIN)_$(MSGP_VERSION)


.minimal.makefile:
	curl -fsSL -o $@ https://gitlab.com/bsm/misc/raw/master/make/go/minimal.makefile

include .minimal.makefile


$(MSGP):
	GOBIN=$(PWD)/$(DEPENDENCIES_FOLDER) go install github.com/tinylib/msgp@$(MSGP_VERSION)
	$(call clean_tool,$(MSGP_BIN),$@)


.PHONY: generate-msgp ## Generate mocks
generate-msgp:$(MSGP)
	go generate ./...
