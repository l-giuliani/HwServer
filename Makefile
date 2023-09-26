BUILDPATH = dist/
EXE = gpioServer

$(EXE): 
	go build -o $(BUILDPATH)$(EXE)
	cp extlib/libNP6118BaseLib.so $(BUILDPATH)

all: $(EXE)

clean:
	rm -f $(BUILDPATH)$(EXE)

run:
	$(BUILDPATH)$(EXE)

size:
	ls -la $(BUILDPATH)
