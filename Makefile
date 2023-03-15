fpath="/usr/local/share/zsh/site-functions"
binpath="/usr/local/bin"

build: clean prep
	go build -o ./ws main.go

prep:
	go mod tidy

zsh-completions:
	sudo mkdir -p "${fpath}"
	ws --completion zsh > "./_ws"
	sudo mv "./_ws" "${fpath}/_ws"

install:
	mv ./ws "${binpath}/ws"

clean:
	rm -f ./ws

uninstall:
	rm -f "${binpath}/ws"
	rm -f "${fpath}/_ws"
