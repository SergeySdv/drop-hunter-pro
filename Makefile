
dev := '170.64.160.93'
prod := '128.199.142.47'


build-dev:
	rm -rf build
	mkdir build
	go generate ./...
	go version
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./build/cry ./cmd/service/main.go
	upx ./build/cry --best --lzma
	cp -r internal/server/migrations ./build/
	cp .env.dev ./build/.env
	cp Dockerfile ./build/Dockerfile
	cp docker-compose.yml ./build/docker-compose.yaml

build-prod:
	rm -rf build
	mkdir build
	go generate ./...
	go version
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./build/cry ./cmd/service/main.go
	upx ./build/cry --best --lzma
	cp -r internal/server/migrations ./build/
	cp .env.prod ./build/.env
	cp Dockerfile ./build/Dockerfile
	cp docker-compose.yml ./build/docker-compose.yaml

send-dev:
	rsync --progress -re ssh ./build/ root@$(dev):/root/app/cry/
	ssh -t root@$(dev) "docker-compose -f /root/app/cry/docker-compose.yaml build --force"
	ssh -t root@$(dev) "docker-compose -f /root/app/cry/docker-compose.yaml pull"
	ssh -t root@$(dev) "docker-compose -f /root/app/cry/docker-compose.yaml up -d"

deploy-dev:
	$(MAKE) build-dev
	$(MAKE) send-dev

deploy-prod:
	$(MAKE) build-prod
	$(MAKE) send-prod

send-prod:
	rsync --progress -re ssh ./build/ root@$(prod):/root/app/cry/
	ssh -t root@$(prod) "docker-compose -f /root/app/cry/docker-compose.yaml build --force"
	ssh -t root@$(prod) "docker-compose -f /root/app/cry/docker-compose.yaml pull"
	ssh -t root@$(prod) "docker-compose -f /root/app/cry/docker-compose.yaml up -d"

