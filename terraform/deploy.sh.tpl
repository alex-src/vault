# this deployment script assumes you've already prepared vault
# with step 1 from https://github.com/alex-src/vault/wiki/Production-Deployment

# download goldfish binary
curl -L -o goldfish https://github.com/alex-src/vault/releases/download/${GOLDFISH_VERSION}/goldfish-linux-amd64

# start systemd service
systemctl daemon-reload
systemctl start goldfish.service
