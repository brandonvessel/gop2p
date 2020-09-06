go build -ldflags="-s -w"
upx client --exact --best --brute
#upx chat --exact --best --ultra-brute --all-methods --all-filters