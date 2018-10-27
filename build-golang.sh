# You need go (v1.9), npm (v5), nodejs (v6)
# You will also need cross-env (sudo npm install -g cross-env)

# code-ify static assets
go get github.com/GeertJohan/go.rice
go get github.com/GeertJohan/go.rice/rice
rm -f rice-box.go
rice embed-go

# compile goldfish binary
#go get github.com/caiyeon/goldfish
go get github.com/hashicorp/vault/helper/mlock
go get github.com/alex-src/vault
mkdir -p build
env GOOS=linux GOARCH=amd64 go build -o build/vault-linux-amd64 -v github.com/alex-src/vault &
env GOOS=windows GOARCH=amd64 go build -o build/vault-windows-amd64.exe -v github.com/alex-src/vault &
env GOOS=darwin GOARCH=amd64 go build -o build/vault-osx-amd64 -v github.com/alex-src/vault &

# report build
wait
echo 'Successfully built ' $(git describe --always --tags)
