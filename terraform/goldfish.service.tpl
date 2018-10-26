[Unit]
Description=vault
Requires=network-online.target
After=network-online.target
[Service]
User=root
Group=root
WorkingDirectory=/home/ubuntu
ExecStart=/home/ubuntu/vault-linux-amd64 -config=/etc/vault-config.hcl -token=${WRAPPING_TOKEN}
ExecReload=/bin/kill -HUP $MAINPID
KillSignal=SIGINT
[Install]
WantedBy=multi-user.target
