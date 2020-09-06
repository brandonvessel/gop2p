go build -ldflags="-s -w"
upx peer --exact --best --brute
#upx chat --exact --best --ultra-brute --all-methods --all-filters