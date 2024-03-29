
run-dev:
	@echo "run"
	air -c .air.toml

build:
	go build -o ./dist/main
