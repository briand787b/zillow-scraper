provider "digitalocean" {
    token = var.do_token
}

# terraform {
#     // access_key comes from env var AWS_ACCESS_KEY_ID
#     // secret_key comes from env var AWS_SECRET_ACCESS_KEY
#     backend "s3" {
        
#     }
# }

terraform {
  backend "s3" {
    endpoint                    = "nyc3.digitaloceanspaces.com"
    key                         = "terraform.tfstate"
    bucket                      = "briand787b"
    region                      = "us-west-1"
    skip_requesting_account_id  = true
    skip_credentials_validation = true
    skip_get_ec2_platforms      = true
    skip_metadata_api_check     = true
  }
}



module "registry" {
    source = "../../modules/container_registry"
}