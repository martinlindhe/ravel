watch:
	gorazor -watch views views

deps:
	npm install

serve:
	go run app.go

ping:
	curl "http://localhost:8080/ping"
