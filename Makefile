
api:
	go run cmd/main.go -port 9000

fe:
	cd public/fe && bun run dev

format:
	goimports-reviser -rm-unused -use-cache -set-alias -format ./...