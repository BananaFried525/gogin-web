default:
	docker build -t gogin-web:v.1.0.0 .

up: default
	docker-compose up -d

logs:
	docker-compose logs -f

down:
	docker-compose down

test:
	go test -v -cover ./...

clean: 
	down
	rm -f api
	docker system prune -f
	docker volume prune -f
prune:
	docker system prune -f
	docker volume prune -f
run:
	go run server.go