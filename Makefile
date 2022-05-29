GENERATOR_ENUMS = src/generators/enums.go

EXECUTOR_ENUMS = src/executor/enums.go

$(GENERATOR_ENUMS): GO_ENUM_FLAGS=--nocase --marshal --names --ptr
$(EXECUTOR_ENUMS): GO_ENUM_FLAGS=--nocase --marshal --names --sqlnullint --ptr

enums: $(GENERATOR_ENUMS) $(EXECUTOR_ENUMS)

# The generator statement for go enum files.  Files that invalidate the
# enum file: source file, the binary itself, and this file (in case you want to generate with different flags)
%_enum.go: %.go $(GOENUM) Makefile
	$(GOENUM) -f $*.go $(GO_ENUM_FLAGS)