# doll.report

## what is `doll.report`?
doll.report is a communal infrastructure project, providing a platform for independent uptime tracking
(and downtime alerts) of infrastructure and services.


## How does it work?
The rough idea is that anyone can create a merge request, add their own [gatus](https://github.com/TwiN/gatus)
configuration file into the [dashboards/](dashboards/) folder, and have said gatus instance automatically provisioned upon approval.

When pushing to the repo [`yamllint`](https://github.com/adrienverge/yamllint) will be run,
make sure your branch passes this or the merge request will not be approved. the `yamllint` package is available on PyPi


## `<dashboard_name>.yaml` definitions
The `<dashboard_name>.yaml` files contain [gatus](https://github.com/TwiN/gatus) configuration.

The filename (excluding `.yaml`) will be used as the subdomain for the dashboard,
your instance will be available at `https://<name of dashboard>.doll.report`.

Feel free to customize the ui section to your desire within your own defintion file !! :)


## Rules
- Keep your probe intervals above two seconds (this applies to everything apart from domain expiration probes)
- Don't keep excessive history data (leave `storage.maximum-number-of-results` and
  `storage.maximum-number-of-events` at deafult unless absolutely required)
- Domain expiration probe intervals not allowed to be lower than 30 minutes between requests.
  (See more info here: https://github.com/TwiN/gatus?tab=readme-ov-file#monitoring-domain-expiration)

## Alerts?

Alerts are coming very soon. Thermia will shim auth details for an smtp server into your gatus config through the deploy pipeline, which you
can then use to set up email alerts for your own dashboard.

## Issues?

Feel free to open issues on the repo if there are any questions!
