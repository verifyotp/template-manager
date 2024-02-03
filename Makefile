run:
	go run cmd/main.go -port 9000

format:
	goimports-reviser -rm-unused -use-cache -set-alias -format ./...