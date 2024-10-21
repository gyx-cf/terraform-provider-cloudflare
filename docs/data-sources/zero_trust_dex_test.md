---
page_title: "cloudflare_zero_trust_dex_test Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_zero_trust_dex_test (Data Source)




<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `account_id` (String)
- `dex_test_id` (String) API UUID.
- `filter` (Attributes) (see [below for nested schema](#nestedatt--filter))

### Read-Only

- `data` (Attributes) The configuration object which contains the details for the WARP client to conduct the test. (see [below for nested schema](#nestedatt--data))
- `description` (String) Additional details about the test.
- `enabled` (Boolean) Determines whether or not the test is active.
- `interval` (String) How often the test will run.
- `name` (String) The name of the DEX test. Must be unique.
- `target_policies` (Attributes List) Device settings profiles targeted by this test (see [below for nested schema](#nestedatt--target_policies))
- `targeted` (Boolean)

<a id="nestedatt--filter"></a>
### Nested Schema for `filter`

Required:

- `account_id` (String)


<a id="nestedatt--data"></a>
### Nested Schema for `data`

Read-Only:

- `host` (String) The desired endpoint to test.
- `kind` (String) The type of test.
- `method` (String) The HTTP request method type.


<a id="nestedatt--target_policies"></a>
### Nested Schema for `target_policies`

Read-Only:

- `default` (Boolean) Whether the profile is the account default
- `id` (String) The id of the device settings profile
- `name` (String) The name of the device settings profile

