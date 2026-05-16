connection "backstage" {
    plugin = "local/steampipe-plugin-backstage"

    # Backstage instance URL.
    # Can also be set with BACKSTAGE_HOST environment variable.
    host = "https://demo.backstage.io"

    # Backstage API token for authentication (optional).
    # Required only for Backstage instances with authentication enabled.
    # Can also be set with BACKSTAGE_TOKEN environment variable.
    # token = "your-token-here"
}
