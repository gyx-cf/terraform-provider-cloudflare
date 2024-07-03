---
page_title: "cloudflare_access_identity_provider Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_access_identity_provider (Resource)



~> It's required that an `account_id` or `zone_id` is provided and in
   most cases using either is fine. However, if you're using a scoped
   access token, you must provide the argument that matches the token's
   scope. For example, an access token that is scoped to the "example.com"
   zone needs to use the `zone_id` argument.

## Example Usage

```terraform
# one time pin
resource "cloudflare_access_identity_provider" "pin_login" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "PIN login"
  type       = "onetimepin"
}

# oauth
resource "cloudflare_access_identity_provider" "github_oauth" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "GitHub OAuth"
  type       = "github"
  config = {
    client_id     = "example"
    client_secret = "secret_key"
  }
}

# saml
resource "cloudflare_access_identity_provider" "jumpcloud_saml" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "JumpCloud SAML"
  type       = "saml"
  config = {
    issuer_url      = "jumpcloud"
    sso_target_url  = "https://sso.myexample.jumpcloud.com/saml2/cloudflareaccess"
    attributes      = ["email", "username"]
    sign_request    = false
    idp_public_cert = "MIIDpDCCAoygAwIBAgIGAV2ka+55MA0GCSqGSIb3DQEBCwUAMIGSMQswCQ...GF/Q2/MHadws97cZg\nuTnQyuOqPuHbnN83d/2l1NSYKCbHt24o"
  }
}

# okta
resource "cloudflare_access_identity_provider" "okta" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "Okta"
  type       = "okta"
  config = {
    client_id     = "example"
    client_secret = "secret_key"
    api_token     = "okta_api_token"
    okta_account  = "https://example.com"
  }
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `config` (Attributes) The configuration parameters for the identity provider. To view the required parameters for a specific provider, refer to our [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/). (see [below for nested schema](#nestedatt--config))
- `name` (String) The name of the identity provider, shown to users on the login page.
- `type` (String) The type of identity provider. To determine the value for a specific provider, refer to our [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).

### Optional

- `account_id` (String) The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
- `id` (String) UUID
- `scim_config` (Attributes) The configuration settings for enabling a System for Cross-Domain Identity Management (SCIM) with the identity provider. (see [below for nested schema](#nestedatt--scim_config))
- `zone_id` (String) The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.

<a id="nestedatt--config"></a>
### Nested Schema for `config`

Optional:

- `claims` (List of String) Custom claims
- `client_id` (String) Your OAuth Client ID
- `client_secret` (String) Your OAuth Client Secret
- `conditional_access_enabled` (Boolean) Should Cloudflare try to load authentication contexts from your account
- `directory_id` (String) Your Azure directory uuid
- `email_claim_name` (String) The claim name for email in the id_token response.
- `prompt` (String) Indicates the type of user interaction that is required. prompt=login forces the user to enter their credentials on that request, negating single-sign on. prompt=none is the opposite. It ensures that the user isn't presented with any interactive prompt. If the request can't be completed silently by using single-sign on, the Microsoft identity platform returns an interaction_required error. prompt=select_account interrupts single sign-on providing account selection experience listing all the accounts either in session or any remembered account or an option to choose to use a different account altogether.
- `support_groups` (Boolean) Should Cloudflare try to load groups from your account


<a id="nestedatt--scim_config"></a>
### Nested Schema for `scim_config`

Optional:

- `enabled` (Boolean) A flag to enable or disable SCIM for the identity provider.
- `group_member_deprovision` (Boolean) A flag to revoke a user's session in Access and force a reauthentication on the user's Gateway session when they have been added or removed from a group in the Identity Provider.
- `seat_deprovision` (Boolean) A flag to remove a user's seat in Zero Trust when they have been deprovisioned in the Identity Provider.  This cannot be enabled unless user_deprovision is also enabled.
- `secret` (String) A read-only token generated when the SCIM integration is enabled for the first time.  It is redacted on subsequent requests.  If you lose this you will need to refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
- `user_deprovision` (Boolean) A flag to enable revoking a user's session in Access and Gateway when they have been deprovisioned in the Identity Provider.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_access_identity_provider.example <account_id>/<identity_provider_id>
```