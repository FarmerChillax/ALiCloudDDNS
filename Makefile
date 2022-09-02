build-win:
	CGO_ENABLE=0 GOOS=windows go build -o bin/win/fddns.exe

build-darwin:
	CGO_ENABLE=0 GOOS=darwin go build -o bin/darwin/fddns

build-linux:
	CGO_ENABLE=0 GOOS=linux go build -o bin/linux/fddns

build-all: build-darwin build-linux build-win

release-linux:
	tar -c bin/linux -f release/fddns_0.2.0_linux.tar

release-darwin:
	tar -c bin/linux -f release/fddns_0.2.0_darwin.tar

release-win:
	zip -r release/fddns_0.2.0_win.zip  bin/win 

release-all: release-linux release-darwin release-win


test:
	go test ./...