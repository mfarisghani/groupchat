run-nsq: 
	@echo " >> running nsq"
	@go run cmd/nsq/main.go

run-nats: 
	@echo " >> running nats"
	@go run cmd/nats/main.go

run-client:
	@go run cmd/client/main.go

run-test:
	@go run cmd/test/main.go
