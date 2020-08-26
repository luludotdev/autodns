# AutoDNS
> Simple golang tool to automatically update DNS records for a server.

## 💾 Installation
AutoDNS is built into a single binary. You can build this yourself with Golang tooling or download a prebuilt release from the [Releases](https://github.com/lolPants/autodns/releases) page.

Nightly builds are also available as artifacts on the [Actions](https://github.com/lolPants/autodns/actions?query=workflow%3A%22Golang+Build%22) page.

## 🚀 Usage
AutoDNS is designed to be used in automated environments such as inside `cron`. As such, everything is ran using a CLI.
```
Available Commands:
  cf          Use CloudFlare DNS provider
  help        Help about any command     
  version     Print version information

Flags:
  -h, --help            help for autodns
  -v, --verbose count   verbose output
  -V, --version         print version
```

Additionally, you can instead pass the API tokens using the `AUTODNS_{provider}_TOKEN` environment variables. Note that tokens passed using CLI flags take priority over environment variables. [See below](#providers) for a complete list.

### ⌛ Example with Crontab
This example assumes you have installed AutoDNS into your `PATH`.
```cron
# Run AutoDNS every hour and discard output
0 * * * * autodns cf --token notarealtoken --record server.jackbaron.dev > /dev/null 2>&1
```

## 📡 Providers
These are all the DNS providers currently supported by AutoDNS. If you would like to see another added, submit an issue. Instead if you're feeling generous, shoot me a PR to add a provider and I'll be eternally grateful :smile:

| Provider | Subcommand | Environment Variable |
| - | - | - |
| [CloudFlare](http://cloudflare.com/) | `autodns cf` | `AUTODNS_CF_TOKEN` |
