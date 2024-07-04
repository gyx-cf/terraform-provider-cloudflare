---
page_title: "cloudflare_static_route Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_static_route (Resource)



## Example Usage

```terraform
resource "cloudflare_static_route" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  description = "New route for new prefix 192.0.2.0/24"
  prefix      = "192.0.2.0/24"
  nexthop     = "10.0.0.0"
  priority    = 100
  weight      = 10
  colo_names = [
    "den01"
  ]
  colo_regions = [
    "APAC"
  ]
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Identifier

### Optional

- `description` (String) An optional human provided description of the static route.
- `nexthop` (String) The next-hop IP Address for the static route.
- `prefix` (String) IP Prefix in Classless Inter-Domain Routing format.
- `priority` (Number) Priority of the static route.
- `route_id` (String) Identifier
- `scope` (Attributes) Used only for ECMP routes. (see [below for nested schema](#nestedatt--scope))
- `weight` (Number) Optional weight of the ECMP scope - if provided.

### Read-Only

- `deleted` (Boolean)
- `deleted_route` (String)
- `modified` (Boolean)
- `modified_route` (String)
- `route` (String)
- `routes` (Attributes List) (see [below for nested schema](#nestedatt--routes))

<a id="nestedatt--scope"></a>
### Nested Schema for `scope`

Optional:

- `colo_names` (List of String) List of colo names for the ECMP scope.
- `colo_regions` (List of String) List of colo regions for the ECMP scope.


<a id="nestedatt--routes"></a>
### Nested Schema for `routes`

Required:

- `nexthop` (String) The next-hop IP Address for the static route.
- `prefix` (String) IP Prefix in Classless Inter-Domain Routing format.
- `priority` (Number) Priority of the static route.

Optional:

- `description` (String) An optional human provided description of the static route.
- `scope` (Attributes) Used only for ECMP routes. (see [below for nested schema](#nestedatt--routes--scope))
- `weight` (Number) Optional weight of the ECMP scope - if provided.

Read-Only:

- `created_on` (String) When the route was created.
- `id` (String) Identifier
- `modified_on` (String) When the route was last modified.

<a id="nestedatt--routes--scope"></a>
### Nested Schema for `routes.scope`

Optional:

- `colo_names` (List of String) List of colo names for the ECMP scope.
- `colo_regions` (List of String) List of colo regions for the ECMP scope.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_static_route.example <account_id>/<static_route_id>
```