# alex2006hw Vault
--------------------

### Why
- Build a local vault to hold secrets

### Usage

- install  source
```
  go get github.com/alex-src/vault
```

- build everything
```
  cd $GOPATH/src/github.com/alex-src/vault
  ./runme
```

- build frontend
```
  cd $GOPATH/src/github.com/alex-src/vault
  ./bin/getFrontend
  ./bin/build-nodejs.sh
```

- build backend
```
  ./bin/build-golang.sh
```

- binaries in build directory
