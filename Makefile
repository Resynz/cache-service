SHELL=/bin/bash

EXE = cache-service

all: $(EXE)

cache-service:
	@echo "building $@ ..."
	$(MAKE) -s -f make.inc s=static

clean:
	rm -f $(EXE)

