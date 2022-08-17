BINARY=engine
engine:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ${BINARY} ./cmd/server/*.go

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t order-service .

run:
	docker-compose up -d

stop:
	docker-compose down

.PHONY: test engine unittest clean docker run stop lint sqlc