variable "AUTH0_CLIENT_ID" {}
variable "AUTH0_CLIENT_SECRET" {}

provider "cyral" {
    auth0_domain = "aday-cyral.auth0.com"
    auth0_audience = "cyral.com"
    control_plane = ""
}

resource "cyral_repository" "some_new_name" {
    name = "some_name"
}
