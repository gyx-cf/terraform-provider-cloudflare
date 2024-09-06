
    resource "cloudflare_zero_trust_access_application" "%[1]s" {
      name       = "%[1]s"
      account_id = "%[3]s"
      domain     = "%[1]s.%[2]s"
    }

    resource "cloudflare_zero_trust_access_service_token" "%[1]s" {
      account_id = "%[3]s"
      name       = "%[1]s"
    }

    resource "cloudflare_zero_trust_access_policy" "%[1]s" {
      application_id = "${cloudflare_zero_trust_access_application.%[1]s.id}"
      name           = "%[1]s"
      account_id     = "%[3]s"
      decision       = "non_identity"
      precedence     = "10"

      include =[ {
        service_token = ["${cloudflare_zero_trust_access_service_token.%[1]s.id}"]
      }]
    }

  