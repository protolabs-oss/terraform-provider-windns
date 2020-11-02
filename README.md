# Terraform Windows DNS Provider

![](https://github.com/protolabs-oss/terraform-provider-windns/workflows/Go-Win/badge.svg)

This is the repository for a Terraform Windows DNS Provider, which can be used to create DNS records in Microsoft Windows DNS.

The provider essentially just "shells out" to PowerShell and performs the actual work over a WinRM session, relying on the user's current Windows login context for authentication. Currently Windows only, and with support for CNAME and A records only.

## Compatability

Terraform 0.12. 

## Installing / Building

Build and install the provider with `Go`:

0. Make sure you have `Go` installed and `$GOPATH` set ($env:GOPATH='c:\go' for example)
1. Run `go get github.com/protolabs-oss/terraform-provider-windns`
2. Run `go build`
3. Copy the file created at `$GOPATH/bin/terraform-provider-windns.exe` into your Terraform plugins directory, as described in  [Terraform's plugin instructions](https://www.terraform.io/docs/plugins/basics.html#installing-plugins). 

Or, try the `Install.ps1` Powershell script. It assumes your `$GOPATH` is `C:\go`.

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


