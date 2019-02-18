name = util
version = `cat ./version`

update_version:
	version

build_linux_amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -o release/linux/amd64/$(name)-$(version)

build_linux_i386:
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -v -a -o release/linux/i386/$(name)-$(version)

test:
	go test -cover -v ./\.\.\.
	go vet

docs:
	swag init