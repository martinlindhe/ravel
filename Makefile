watch:
	gorazor -watch views views

deps:
	npm install

serve:
	go run ./cli/server/server.go

ping:
	curl "http://localhost:8080/ping"
