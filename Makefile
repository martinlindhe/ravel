watch:
	gorazor -watch views views

deps:
	npm install

serve:
	go run ./cli/server/server.go

ping:
	curl "http://localhost:8080/ping"

lint:
	find ./db/migrations -name '*.yml' -print0 | xargs -0 -n1 js-yaml
