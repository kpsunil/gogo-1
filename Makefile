CC=        go
BIN=       ./bin
SRC=       ./src
CLEANDIR=  $(SRC)/goccgen
GCFLAGS=   -ldflags "-w"
DEBUGFLAGS=-gcflags "-N -l"

all:
	make deps
	make gogo

.PHONY: scanner tac gogo clean test

deps: $(SRC)/lang.bnf
	scripts/check-gocc.sh

scanner: $(SRC)/scanner/gentoken.go
	cd $(SRC)/scanner; $(CC) install $(GCFLAGS)

tac: $(SRC)/tac/tac.go
	cd $(SRC)/tac; $(CC) install $(GCFLAGS)

gogo: $(SRC)/main.go
	mkdir -p tmp
	$(CC) build $(GCFLAGS) -o $(BIN)/gogo $(SRC)/main.go

test:
	scripts/run-tests.sh

travis:
	make
	make test
	git diff --exit-code test

clean:
	rm -rf $(CLEANDIR)
	rm -rf $(BIN)
