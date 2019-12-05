


https://medium.com/tarkalabs/automating-dji-tello-drone-using-gobot-2b711bf42af6


GO111MODULE=off go get -d -u gobot.io/x/gobot/...


brew install ffmpeg $(brew options ffmpeg | grep -v -e '\s' | grep -e '--with-\|--HEAD' | tr '\n' ' ')

export LDFLAGS="-L/usr/local/opt/libffi/lib"
export PKG_CONFIG_PATH="/usr/local/opt/libffi/lib/pkgconfig"


export LDFLAGS=$LDFLAGS" -L/usr/local/opt/openblas/lib"
export PKG_CONFIG_PATH="/usr/local/opt/openblas/lib/pkgconfig":$PKG_CONFIG_PATH
export CPPFLAGS="-I/usr/local/opt/openblas/include"

https://godoc.org/gobot.io/x/gobot/platforms/dji/tello



go build -mod vendor

