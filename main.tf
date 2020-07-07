provider "auth0" {
    domain = "abc.eu.auth0.com"
    client_id = "<CLIENT_ID>"
    client_secret = "<CLIENT_SECRET>"
}

resource "auth0_client" "test-client" {
    name = "p-test-dev2"
    is_token_endpoint_ip_header_trusted = true
    is_first_party = false
    description = ""
    cross_origin_auth = false
    sso = true
    token_endpoint_auth_method = "client_secret_post"
    grant_types = [ "authorization_code", "implicit", "refresh_token", "client_credentials" ]
    app_type = "non_interactive"
    custom_login_page_on = true
}