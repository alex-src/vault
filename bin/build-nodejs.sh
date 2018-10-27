# You need go (v1.9), npm (v5), nodejs (v6)
# You will also need cross-env (sudo npm install -g cross-env)
ROOT_DIR=$(pwd)

# install & update npm packages
#(cd ${ROOT_DIR} && rm -rf public pkg)
#(cd ${ROOT_DIR}/frontend && npm install && npm run build)
#(cd ${ROOT_DIR} && mv pkg/web_ui public && rmdir pkg)
(cd ${ROOT_DIR}/frontend && npm install && npm run build-web)
(cd ${ROOT_DIR} && rm -rf ${ROOT_DIR}/public && mkdir -p ${ROOT_DIR}/public && cp -Ra ${ROOT_DIR}/frontend/dist ${ROOT_DIR}/public/ && cp ${ROOT_DIR}/frontend/index.web.html ${ROOT_DIR}/public/index.html)
# report build
echo 'Successfully built ' $(git describe --always --tags)
