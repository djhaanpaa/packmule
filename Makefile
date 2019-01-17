all: compile install
compile:
	go build
install: 
	cp packmule /usr/local/bin
clean:
	rm -rv *.yaml
	rm -rv *.yml
	rm -rv ./tmp
pushtoscm:
	git push origin master
	git push github master
pullfromscm:
	git pull origin master
	git pull github master
