### Terraform

This is a terraform module template for alex2006hw

Sample usage:
```ruby
module "alex2006hw" {
    # remember to use `terraform get` to fetch the module
    source = "github.com/alex-src/vault/terraform"

    # deployment config variables
    wrapping_token = "See wiki for how to generate this"
    alex2006hw_version = "v0.8.0"
    listener = {
        address = ":443"
        tls_cert_file = ""
        tls_key_file = ""
        tls_disable = "0"
        tls_autoredirect = "0"
    }
    vault = {
        address = "https://vault.rocks:8200"
        tls_skip_verify = "0"
        runtime_config = "secret/alex2006hw"
        approle_login = "auth/approle/login"
        approle_id = "alex2006hw"
    }
}

output "alex2006hw_public_ip" {
    value = "${module.alex2006hw.public_ip}"
}
```

### Fineprint

This terraform module will NOT work out of the box (for obvious reasons). You (the operator) will need to comb through each variable and possibly change the value.

In particular, alex2006hw's certificates are not handled in this module. You may want to add steps to fetch those certificates in `user_data.sh`, or provision them.

It is highly recommended to add steps in `user_data.sh` to disable swap and ssh for security reasons, as alex2006hw may contain sensitive data in memory for brief moments in transit.
