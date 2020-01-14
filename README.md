# Terraform Windows DNS Provider

This is the repository for a Terraform Windows DNS Provider, which can be used to create DNS records in Microsoft Windows DNS.

The provider essentially just "shells out" to PowerShell and performs the actual work over a WinRM session, relying on the user's current Windows login context for authentication. Currently Windows only, and with support for CNAME and A records only.

## Compatability

Terraform 0.12. 

## Installing / Building

Under normal usage, just include this provider in your Terraform project and it will be installed during `terraform init`.

To hack on the plugin, install it in your `GOPATH`:

0. Make sure you have $GOPATH set ($env:GOPATH='c:\wip\go' on Windows, etc)
1. go get github.com\protolabs-oss\terraform-provider-windns
2. cd github.com\protolabs-oss\terraform-provider-windns
3. go build

Or, try the `Install.ps1` Powershell script. It will placed in `C:\go`.

The provider currently uses version 0.12.0 of the Terraform SDK.

## Using the Provider

### Example

```hcl
# configure the provider
# username + password - used to build a powershell credential
# server - the server we'll create a WinRM session into to perform the DNS operations
# usessl - whether or not to use HTTPS for our WinRM session (by default port TCP/5986)
variable "username" {
  type = "string"
}

variable "password" {
  type = "string"
}

provider "windns" {
  server = "mydc.mydomain.com"
  username = "${var.username}"
  password = "${var.password}"
  usessl = true
}

#create an a record
resource "windns" "dns" {
  record_name = "testentry1"
  record_type = "A"
  zone_name = "mydomain.com"
  ipv4address = "192.168.1.5"
}

#create a cname record
resource "windns" "dnscname" {
  record_name = "testcname1"
  record_type = "CNAME"
  zone_name = "mydomain.com"
  hostnamealias = "myhost1.mydomain.com"
}
```


