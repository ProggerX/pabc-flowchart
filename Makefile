run:
	go run ./cmd/pabc-parse idk.pas && (env cat ./a.out.md | wl-copy)
deb:
	go run ./cmd/pabc-parse idk.pas 5 && (env cat ./a.out.md | wl-copy)
