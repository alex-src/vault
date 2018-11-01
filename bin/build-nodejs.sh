# You need go (v1.9), npm (v5), nodejs (v6)
# You will also need cross-env (sudo npm install -g cross-env)
ROOT_DIR=$(pwd)

# install & update npm packages
#(cd ${ROOT_DIR} && rm -rf public pkg)
#(cd ${ROOT_DIR}/frontend && npm install && npm run build)
#(cd ${ROOT_DIR} && mv pkg/web_ui public && rmdir pkg)
(cd ${ROOT_DIR} && rm -rf ${ROOT_DIR}/build/frontend public && mkdir -p ${ROOT_DIR}/build/frontend)
(cd ${ROOT_DIR}/frontend && npm install && npm run build-web)
(cd ${ROOT_DIR}/frontend && ${ROOT_DIR}/node_modules/pkg/lib-es5/bin.js -t node8-alpine,node8-linux,node8-mac,node8-win -o ${ROOT_DIR}/build/frontend/ui .)
(cp -Ra ${ROOT_DIR}/frontend/dist ${ROOT_DIR}/build/frontend/ && cp ${ROOT_DIR}/frontend/index.web.html ${ROOT_DIR}/build/frontend/)
(cd -Ra ${ROOT_DIR}/build/frontend ${ROOT_DIR}/public)

#
# report build
echo 'Successfully built ' $(git describe --always --tags)
