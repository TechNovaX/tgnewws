CACHE = false
# cache 代表是否需要缓存
# 生成 rpc 代码 windows
gen-api:
	goctl api go -api tgnewws.api -dir ./ -style gozero

build:
	go env -w CGO_ENABLED=0 && go env -w GOOS=linux && go env -w GOARCH=amd64 && go build -o news-api

rebuild:
	git pull && go build -o news-api && pm2 restart news-api