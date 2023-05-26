BUILDPATH = dist/
EXE = gpioServer

$(EXE): 
	go build -o $(BUILDPATH)$(EXE)

all: $(EXE)

clean:
	rm -f $(BUILDPATH)$(EXE)

run:
	$(BUILDPATH)$(EXE)

size:
	ls -la $(BUILDPATH)
