
dev := '170.64.160.93'
prod := '128.199.142.47'

deploy-dev:
	rm -rf build
	mkdir build
	cp docker-compose.yml ./build/docker-compose.yaml
	cp prometheus.yml ./build/prometheus.yml
	cp .env.dev ./build/.env
	rsync --progress -re ssh ./build/ root@$(dev):/root/app/cry/infra/
	ssh -t root@$(dev) "docker-compose -f /root/app/cry/infra/docker-compose.yaml up -d"


deploy-prod:
	rm -rf build
	mkdir build
	cp docker-compose.yml ./build/docker-compose.yaml
	cp prometheus.yml ./build/prometheus.yml
	cp .env.prod ./build/.env
	rsync --progress -re ssh ./build/ root@$(prod):/root/app/cry/infra/
	ssh -t root@$(prod) "docker-compose -f /root/app/cry/infra/docker-compose.yaml up -d"




