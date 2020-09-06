program_name="peer"

linux_build () {
    echo "compiling $1 $2"
    env GOOS="$1" GOARCH="$2" go build -ldflags="-s -w" -o "builds/$program_name.$1.$2"
}

linux_pack () {
    upx "builds/$program_name.$1.$2" --exact --best --ultra-brute -o "builds/$program_name.$1.$22"
    mv "builds/$program_name.$1.$22" "builds/$program_name.$1.$2"
}

windows_build() {
    echo "compiling $1 $2"
    env GOOS="windows" GOARCH="$2" go build
}

windows_build64() {
    echo "compiling $1 $2"
    env GOOS="windows" GOARCH="$2" go build
}

mkdir "builds"

linux_build "android" "arm"
linux_build "darwin" "386"
linux_build "darwin" "amd64"
linux_build "darwin" "arm"
linux_build "darwin" "arm64"
linux_build "dragonfly" "amd64"
linux_build "freebsd" "386"
linux_build "freebsd" "amd64"
linux_build "freebsd" "arm"
linux_build "linux" "386"
linux_build "linux" "amd64"
linux_build "linux" "arm"
linux_build "linux" "arm64"
linux_build "linux" "ppc64"
linux_build "linux" "ppc64le" 
linux_build "linux" "mips"
linux_build "linux" "mipsle"
linux_build "linux" "mips64"
linux_build "linux" "mips64le"
linux_build "netbsd" "386"
linux_build "netbsd" "amd64"
linux_build "netbsd" "arm"
linux_build "openbsd" "386"
linux_build "openbsd" "amd64" 
linux_build "openbsd" "arm"
linux_build "plan9" "386"
linux_build "plan9" "amd64"
linux_build "solaris" "amd64" 
#windows_build "386"
#windows_build64 "amd64"

linux_pack "android" "arm"
linux_pack "darwin" "386"
linux_pack "darwin" "amd64"
linux_pack "darwin" "arm"
linux_pack "darwin" "arm64"
linux_pack "dragonfly" "amd64"
linux_pack "freebsd" "386"
linux_pack "freebsd" "amd64"
linux_pack "freebsd" "arm"
linux_pack "linux" "386"
linux_pack "linux" "amd64"
linux_pack "linux" "arm"
linux_pack "linux" "arm64"
linux_pack "linux" "ppc64"
linux_pack "linux" "ppc64le" 
linux_pack "linux" "mips"
linux_pack "linux" "mipsle"
linux_pack "linux" "mips64"
linux_pack "linux" "mips64le"
linux_pack "netbsd" "386"
linux_pack "netbsd" "amd64"
linux_pack "netbsd" "arm"
linux_pack "openbsd" "386"
linux_pack "openbsd" "amd64" 
linux_pack "openbsd" "arm"
linux_pack "plan9" "386"
linux_pack "plan9" "amd64"
linux_pack "solaris" "amd64" 