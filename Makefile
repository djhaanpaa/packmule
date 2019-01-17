all: compile install
compile:
	go build
install: 
	cp packmule /usr/local/bin
clean:
	rm -rv *.yaml
	rm -rv *.yml
	rm -rv ./tmp
