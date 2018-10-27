# You need go (v1.9), npm (v5), nodejs (v6)
# You will also need cross-env (sudo npm install -g cross-env)

# install & update npm packages
rm -rf public pkg
(cd frontend && npm install && npm run build)
mv pkg/web_ui public && rmdir pkg
# report build
echo 'Successfully built ' $(git describe --always --tags)
