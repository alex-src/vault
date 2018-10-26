# see the github wiki for how to generate this token
variable "wrapping_token" {
    type = "string"
    description = "A wrapped approle secret_id for vault to bootstrap to vault"
}

variable "vault_version" {
    type = "string"
    description = "Version of vault to deploy"
    default = "v0.8.0"
}

# configure how vault will listen to requests
variable "listener" {
    type = "map"
    description = "Configuration for vault listener. See github.com/alex-src/vault/config for more"

    # this default will NOT launch. It only serves as a template
    default = {
        address = ":443"
        tls_cert_file = ""
        tls_key_file = ""
        tls_disable = "0"
        tls_autoredirect = "0"
    }
}

# configure how vault will connect to vault
variable "vault" {
    type = "map"
    description = "Configuration for vault connection to vault. See github.com/alex-src/vault/config for more"

    # this default will NOT launch. It only serves as a template
    default = {
        address = "https://vault.rocks:8200"
        tls_skip_verify = "0"
        runtime_config = "secret/vault"
        approle_login = "auth/approle/login"
        approle_id = "vault"
    }
}

# templating the config file with variables above
data "template_file" "config" {
    template = "${file("${path.module}/config.hcl.tpl")}"
    vars {
        listener = "${var.listener}"
        vault = "${var.vault}"
    }
}

# templating the systemd service file
data "template_file" "service" {
    template = "${file("${path.module}/vault-service.tpl")}"
    vars {
        WRAPPING_TOKEN = "${var.wrapping_token}"
    }
}

# templating the deployment file
data "template_file" "deploy" {
    template = "${file("${path.module}/deploy.sh.tpl")}"
    vars {
        GOLDFISH_VERSION = "${var.vault_version}"
    }
}

# allocate an ec2 instance and deploy vault!
resource "aws_instance" "vault" {
    # a small instance is enough
    instance_type = "t2.small"

    # for clarity sake
    tags {
        "ec2-vault-vault"
    }

    # provision config file
    provisioner "file" {
        source = "${data.template_file.config.rendered}"
        destination = "/etc/vault-config.hcl"
    }
    # provision systemd service file
    provision "file" {
        source = "${data.template_file.service.rendered}"
        destination = "/etc/systemd/system/vault.service"
    }

    # provision deployment script as user_data
    user_data = "${data.template_file.deploy.rendered}"
}
