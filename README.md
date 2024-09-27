# Backstage plugin for Steampipe

Use SQL to query namespaces, rules and more from [Backstage][https://backstage.io/].

- **[Get started →](docs/index.md)**
- Documentation: [Table definitions & examples](docs/tables)

## Quick start

Install the plugin with [Steampipe][]:

    steampipe plugin install chussenot/backstage

## Development

To build the plugin and install it in your `.steampipe` directory

    make

Copy the default config file:

    cp config/backstage.spc ~/.steampipe/config/backstage.spc

## License

Apache 2

[steampipe]: https://steampipe.io
[backstage]: https://backstage.io/