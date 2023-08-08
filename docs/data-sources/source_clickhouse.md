---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_clickhouse Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceClickhouse DataSource
---

# airbyte_source_clickhouse (Data Source)

SourceClickhouse DataSource

## Example Usage

```terraform
data "airbyte_source_clickhouse" "my_source_clickhouse" {
  secret_id = "...my_secret_id..."
  source_id = "...my_source_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `source_id` (String)

### Optional

- `secret_id` (String) Optional secretID obtained through the public API OAuth redirect flow.

### Read-Only

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String)
- `workspace_id` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Read-Only:

- `database` (String) The name of the database.
- `host` (String) The host endpoint of the Clickhouse cluster.
- `password` (String) The password associated with this username.
- `port` (Number) The port of the database.
- `source_type` (String) must be one of ["clickhouse"]
- `tunnel_method` (Attributes) Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use. (see [below for nested schema](#nestedatt--configuration--tunnel_method))
- `username` (String) The username which is used to access the database.

<a id="nestedatt--configuration--tunnel_method"></a>
### Nested Schema for `configuration.tunnel_method`

Read-Only:

- `source_clickhouse_ssh_tunnel_method_no_tunnel` (Attributes) Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use. (see [below for nested schema](#nestedatt--configuration--tunnel_method--source_clickhouse_ssh_tunnel_method_no_tunnel))
- `source_clickhouse_ssh_tunnel_method_password_authentication` (Attributes) Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use. (see [below for nested schema](#nestedatt--configuration--tunnel_method--source_clickhouse_ssh_tunnel_method_password_authentication))
- `source_clickhouse_ssh_tunnel_method_ssh_key_authentication` (Attributes) Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use. (see [below for nested schema](#nestedatt--configuration--tunnel_method--source_clickhouse_ssh_tunnel_method_ssh_key_authentication))
- `source_clickhouse_update_ssh_tunnel_method_no_tunnel` (Attributes) Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use. (see [below for nested schema](#nestedatt--configuration--tunnel_method--source_clickhouse_update_ssh_tunnel_method_no_tunnel))
- `source_clickhouse_update_ssh_tunnel_method_password_authentication` (Attributes) Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use. (see [below for nested schema](#nestedatt--configuration--tunnel_method--source_clickhouse_update_ssh_tunnel_method_password_authentication))
- `source_clickhouse_update_ssh_tunnel_method_ssh_key_authentication` (Attributes) Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use. (see [below for nested schema](#nestedatt--configuration--tunnel_method--source_clickhouse_update_ssh_tunnel_method_ssh_key_authentication))

<a id="nestedatt--configuration--tunnel_method--source_clickhouse_ssh_tunnel_method_no_tunnel"></a>
### Nested Schema for `configuration.tunnel_method.source_clickhouse_ssh_tunnel_method_no_tunnel`

Read-Only:

- `tunnel_method` (String) must be one of ["NO_TUNNEL"]
No ssh tunnel needed to connect to database


<a id="nestedatt--configuration--tunnel_method--source_clickhouse_ssh_tunnel_method_password_authentication"></a>
### Nested Schema for `configuration.tunnel_method.source_clickhouse_ssh_tunnel_method_password_authentication`

Read-Only:

- `tunnel_host` (String) Hostname of the jump server host that allows inbound ssh tunnel.
- `tunnel_method` (String) must be one of ["SSH_PASSWORD_AUTH"]
Connect through a jump server tunnel host using username and password authentication
- `tunnel_port` (Number) Port on the proxy/jump server that accepts inbound ssh connections.
- `tunnel_user` (String) OS-level username for logging into the jump server host
- `tunnel_user_password` (String) OS-level password for logging into the jump server host


<a id="nestedatt--configuration--tunnel_method--source_clickhouse_ssh_tunnel_method_ssh_key_authentication"></a>
### Nested Schema for `configuration.tunnel_method.source_clickhouse_ssh_tunnel_method_ssh_key_authentication`

Read-Only:

- `ssh_key` (String) OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
- `tunnel_host` (String) Hostname of the jump server host that allows inbound ssh tunnel.
- `tunnel_method` (String) must be one of ["SSH_KEY_AUTH"]
Connect through a jump server tunnel host using username and ssh key
- `tunnel_port` (Number) Port on the proxy/jump server that accepts inbound ssh connections.
- `tunnel_user` (String) OS-level username for logging into the jump server host.


<a id="nestedatt--configuration--tunnel_method--source_clickhouse_update_ssh_tunnel_method_no_tunnel"></a>
### Nested Schema for `configuration.tunnel_method.source_clickhouse_update_ssh_tunnel_method_no_tunnel`

Read-Only:

- `tunnel_method` (String) must be one of ["NO_TUNNEL"]
No ssh tunnel needed to connect to database


<a id="nestedatt--configuration--tunnel_method--source_clickhouse_update_ssh_tunnel_method_password_authentication"></a>
### Nested Schema for `configuration.tunnel_method.source_clickhouse_update_ssh_tunnel_method_password_authentication`

Read-Only:

- `tunnel_host` (String) Hostname of the jump server host that allows inbound ssh tunnel.
- `tunnel_method` (String) must be one of ["SSH_PASSWORD_AUTH"]
Connect through a jump server tunnel host using username and password authentication
- `tunnel_port` (Number) Port on the proxy/jump server that accepts inbound ssh connections.
- `tunnel_user` (String) OS-level username for logging into the jump server host
- `tunnel_user_password` (String) OS-level password for logging into the jump server host


<a id="nestedatt--configuration--tunnel_method--source_clickhouse_update_ssh_tunnel_method_ssh_key_authentication"></a>
### Nested Schema for `configuration.tunnel_method.source_clickhouse_update_ssh_tunnel_method_ssh_key_authentication`

Read-Only:

- `ssh_key` (String) OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
- `tunnel_host` (String) Hostname of the jump server host that allows inbound ssh tunnel.
- `tunnel_method` (String) must be one of ["SSH_KEY_AUTH"]
Connect through a jump server tunnel host using username and ssh key
- `tunnel_port` (Number) Port on the proxy/jump server that accepts inbound ssh connections.
- `tunnel_user` (String) OS-level username for logging into the jump server host.

