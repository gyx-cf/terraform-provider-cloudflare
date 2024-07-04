---
page_title: "cloudflare_account_member Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_account_member (Resource)



## Example Usage

```terraform
resource "cloudflare_account_member" "example" {
  account_id    = "f037e56e89293a057740de681ac9abbe"
  email_address = "user@example.com"
  role_ids = [
    "68b329da9893e34099c7d8ad5cb9c940",
    "d784fa8b6d98d27699781bd9a7cf19f0"
  ]
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Account identifier tag.
- `email` (String) The contact email address of the user.

### Optional

- `policies` (Attributes List) Array of policies associated with this member. (see [below for nested schema](#nestedatt--policies))
- `roles` (String) Array of roles associated with this member.

### Read-Only

- `id` (String) Membership identifier tag.
- `status` (String) A member's status in the account.

<a id="nestedatt--policies"></a>
### Nested Schema for `policies`

Required:

- `access` (String) Allow or deny operations against the resources.
- `permission_groups` (Attributes List) A set of permission groups that are specified to the policy. (see [below for nested schema](#nestedatt--policies--permission_groups))
- `resource_groups` (Attributes List) A list of resource groups that the policy applies to. (see [below for nested schema](#nestedatt--policies--resource_groups))

Read-Only:

- `id` (String) Policy identifier.

<a id="nestedatt--policies--permission_groups"></a>
### Nested Schema for `policies.permission_groups`

Required:

- `id` (String) Identifier of the group.


<a id="nestedatt--policies--resource_groups"></a>
### Nested Schema for `policies.resource_groups`

Required:

- `id` (String) Identifier of the group.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_account_member.example <account_id>/<member_id>
```