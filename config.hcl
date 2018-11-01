disable_cache = true
disable_mlock = true

ui = true

listener "tcp" {
    address = "127.0.0.1:8200"
}

storage "file" {
    location = "./vault-storage"
}

max_lease_ttl = "10h"
default_lease_ttl = "10h"
cluster_name = "testcluster"
pid_file = "./pidfile"
raw_storage_endpoint = true
disable_sealwrap = true
disable_printable_check = true

