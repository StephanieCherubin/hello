#
# Makefile
#
VERSION = snapshot
GHRFLAGS =
.PHONY: build release

default: build

build:
	goxc -d=pkg -pv=$(VERSION)

release:
	ghr  -u StephanieCherubin/hello  $(GHRFLAGS) v$(VERSION) pkg/$(VERSION)
