BUILDPATH = dist/
EXE = gpioServer

$(EXE): 
	go build -o $(BUILDPATH)$(EXE)

all: $(EXE)

clean:
	rm -f $(BUILDPATH)$(EXE)

run:
	$(BUILDPATH)$(EXE)

deploy:
	scp $(BUILDPATH)$(EXE) arpaeuser@172.16.0.90:/home/app/

size:
	ls -la $(BUILDPATH)
