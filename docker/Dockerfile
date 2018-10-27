FROM alpine:3.7

# Version of goldfish to install
ENV GOLDFISH_VERSION=v0.9.0 \
    GOLDFISH_CONFIG= \
    GOLDFISH_PORT=8000 \
    VAULT_TOKEN=

# Directory to put Goldfish binary in
WORKDIR /app

# Install Goldfish binary and clean up
RUN apk --no-cache add \
        --virtual build-dependencies \
          unzip && \
    apk --no-cache add \
          curl \
          ca-certificates && \
          curl -L -o vault https://github.com/alex-src/vault/releases/download/$GOLDFISH_VERSION/vault-linux-amd64 && \
          chmod +x ./vault && \
          apk del build-dependencies

# Default port to expose
EXPOSE $GOLDFISH_PORT

# if env var not defined, run in dev mode
# otherwise, create a local config file, and execute
CMD if [[ -z "${GOLDFISH_CONFIG}" ]]; then \
    /app/vault -dev; \
else \
    echo "$GOLDFISH_CONFIG" > /app/config.hcl && \
    /app/vault -config=/app/config.hcl -token=${VAULT_TOKEN}; \
fi
