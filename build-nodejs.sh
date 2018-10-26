# You need go (v1.9), npm (v5), nodejs (v6)
# You will also need cross-env (sudo npm install -g cross-env)

# install & update npm packages
cd frontend && npm install && npm run build

# report build
echo 'Successfully built ' $(git describe --always --tags)
