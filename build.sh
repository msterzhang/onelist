appName="onelist"
builtAt="$(date +'%F %T %z')"
goVersion=$(go version | sed 's/go version //')
gitAuthor="msterzhang <zhangguanping18@gmail.com>"
gitCommit=$(git log --pretty=format:"%h" -1)

if [ "$1" = "dev" ]; then
  version="dev"
  webVersion="dev"
else
  version=$(git describe --abbrev=0 --tags)
  webVersion=$(wget -qO- -t1 -T2 "https://api.github.com/repos/msterzhang/onelist-web/releases/latest" | grep "tag_name" | head -n 1 | awk -F ":" '{print $2}' | sed 's/\"//g;s/,//g;s/ //g')
fi

echo "backend version: $version"
echo "frontend version: $webVersion"

ldflags="-w -s"

FetchWebDev() {
  curl -L https://codeload.github.com/msterzhang/onelist-web/tar.gz/refs/heads/dev -o web-dist-dev.tar.gz
  tar -zxvf web-dist-dev.tar.gz
  rm -rf public/dist
  mv -f web-dist-dev/dist public
  rm -rf web-dist-dev web-dist-dev.tar.gz
}

FetchWebRelease() {
  curl -L https://github.com/msterzhang/onelist-web/releases/latest/download/dist.tar.gz -o dist.tar.gz
  tar -zxvf dist.tar.gz
  rm -rf public/dist
  mv -f dist public
  rm -rf dist.tar.gz
}

BuildWinArm64() {
  echo building for windows-arm64
  chmod +x ./wrapper/zcc-arm64
  chmod +x ./wrapper/zcxx-arm64
  export GOOS=windows
  export GOARCH=arm64
  export CC=$(pwd)/wrapper/zcc-arm64
  export CXX=$(pwd)/wrapper/zcxx-arm64
  go build -o "$1" -ldflags="$ldflags" -tags=jsoniter .
}

BuildDev() {
  rm -rf .git/
  xgo -targets=linux/amd64,windows/amd64,darwin/amd64 -out "$appName" -ldflags="$ldflags" -tags=jsoniter .
  mkdir -p "dist"
  mv onelist-* dist
  cd dist
  upx -9 ./onelist-linux*
  upx -9 ./onelist-windows-amd64.exe
  find . -type f -print0 | xargs -0 md5sum >md5.txt
  cat md5.txt
}

BuildDocker() {
  go build -o ./bin/onelist -ldflags="$ldflags" -tags=jsoniter .
}

BuildRelease() {
  rm -rf .git/
  mkdir -p "build"
  muslflags="--extldflags '-static -fpic' $ldflags"
  BASE="https://onelist.top/"
  FILES=(x86_64-linux-musl-cross aarch64-linux-musl-cross arm-linux-musleabihf-cross mips-linux-musl-cross mips64-linux-musl-cross mips64el-linux-musl-cross mipsel-linux-musl-cross powerpc64le-linux-musl-cross s390x-linux-musl-cross)
  for i in "${FILES[@]}"; do
    url="${BASE}${i}.tgz"
    curl -L -o "${i}.tgz" "${url}"
    sudo tar xf "${i}.tgz" --strip-components 1 -C /usr/local
  done
  OS_ARCHES=(linux-musl-amd64 linux-musl-arm64 linux-musl-arm linux-musl-mips linux-musl-mips64 linux-musl-mips64le linux-musl-mipsle linux-musl-ppc64le linux-musl-s390x)
  CGO_ARGS=(x86_64-linux-musl-gcc aarch64-linux-musl-gcc arm-linux-musleabihf-gcc mips-linux-musl-gcc mips64-linux-musl-gcc mips64el-linux-musl-gcc mipsel-linux-musl-gcc powerpc64le-linux-musl-gcc s390x-linux-musl-gcc)
  for i in "${!OS_ARCHES[@]}"; do
    os_arch=${OS_ARCHES[$i]}
    cgo_cc=${CGO_ARGS[$i]}
    echo building for ${os_arch}
    export GOOS=${os_arch%%-*}
    export GOARCH=${os_arch##*-}
    export CC=${cgo_cc}
    export CGO_ENABLED=1
    go build -o ./build/$appName-$os_arch -ldflags="$muslflags" -tags=jsoniter .
  done
  BuildWinArm64 ./build/onelist-windows-arm64.exe
  xgo -out "$appName" -ldflags="$ldflags" -tags=jsoniter .
  upx -9 ./onelist-linux-amd64
  upx -9 ./onelist-windows-amd64.exe
  mv onelist-* build
}

MakeRelease() {
  cd build
  mkdir compress
  for i in $(find . -type f -name "$appName-linux-*"); do
    cp "$i" onelist
    tar -czvf compress/"$i".tar.gz onelist
    rm -f onelist
  done
  for i in $(find . -type f -name "$appName-darwin-*"); do
    cp "$i" onelist
    tar -czvf compress/"$i".tar.gz onelist
    rm -f onelist
  done
  for i in $(find . -type f -name "$appName-windows-*"); do
    cp "$i" onelist.exe
    zip compress/$(echo $i | sed 's/\.[^.]*$//').zip onelist.exe
    rm -f onelist.exe
  done
  cd compress
  find . -type f -print0 | xargs -0 md5sum >md5.txt
  cat md5.txt
  cd ../..
}

if [ "$1" = "dev" ]; then
  FetchWebDev
  if [ "$2" = "docker" ]; then
    BuildDocker
  else
    BuildDev
  fi
elif [ "$1" = "release" ]; then
  FetchWebRelease
  if [ "$2" = "docker" ]; then
    BuildDocker
  else
    BuildRelease
    MakeRelease
  fi
else
  echo -e "Parameter error"
fi