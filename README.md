# AutoDNS
> Simple golang tool to automatically update DNS records for a server.

## 💾 Installation
AutoDNS is built into a single binary. You can build this yourself with Golang tooling or download a prebuilt release from the [Releases](https://github.com/lolPants/autodns/releases) page.

Nightly builds are also available as artifacts on the [Actions](https://github.com/lolPants/autodns/actions?query=workflow%3A%22Golang+Build%22) page.

## 🚀 Usage
AutoDNS is designed to be used in automated environments such as inside `cron`. As such, everything can be configured through CLI flags.
```
 -T --api-token string    CloudFlare API Token
 -h --help                Prints this help information
 -r --record string       DNS Record
 -v --version             Print version information
```

Additionally, you can instead pass the API token using the `CLOUDFLARE_TOKEN` environment variable. This takes priority over the CLI flag.

### ⌛ Example with Crontab
This example assumes you have installed AutoDNS into your `PATH`.
```cron
# Run AutoDNS every hour and discard output
0 * * * * autodns --api-token notarealtoken --record server.jackbaron.dev > /dev/null 2>&1
```

## ❗ Limitations
Currently AutoDNS is only designed to work with CloudFlare DNS. If you would like to add other DNS providers, shoot me a Pull Request!
