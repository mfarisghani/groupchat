run-nsq: 
	@echo " >> running nsq"
	@go run cmd/nsq/app.go

run-nats: 
	@echo " >> running nats"
	@go run cmd/nats/app.go

run-client:
	@go run cmd/client/app.go
