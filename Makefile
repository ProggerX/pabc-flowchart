run:
	go run ./cmd/pabc-parse idk.pas
debug:
	go run ./cmd/pabc-parse idk.pas 5
test:
	go run ./cmd/pabc-parse idk.pas 5 && (env cat ./a.out.md | wl-copy)
