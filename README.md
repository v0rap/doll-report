# doll.report

## What is `doll.report`?
doll.report is a communal infrastructure project, providing a platform for independent uptime tracking
(and downtime alerts) of infrastructure and services.


## How does it work?
The rough idea is that anyone can create a merge request, add their own [gatus](https://github.com/TwiN/gatus)
configuration file into the [dashboards/](dashboards/) folder, and have said gatus instance automatically provisioned upon approval.

When pushing to the repo [`yamllint`](https://github.com/adrienverge/yamllint) will be run,
make sure your branch passes this or the merge request will not be approved. the `yamllint` package is available on PyPi

You can validate your config using the go program in `scripts/`
run it like this:

```bash
cd scripts
go run validate-config.go ../dashboards/$your.yaml
```
it exits with error code 3 if the config could not be loaded.


## `<dashboard_name>.yaml` Definition
The `<dashboard_name>.yaml` files contain [gatus](https://github.com/TwiN/gatus) configuration.

The filename (excluding `.yaml`) will be used as the subdomain for the dashboard,
your instance will be available at `https://<name of dashboard>.doll.report`.


## Rules
- Keep your probe intervals above two seconds (this applies to everything apart from domain expiration probes)
- Don't keep excessive history data (leave `storage.maximum-number-of-results` and
  `storage.maximum-number-of-events` at deafult unless absolutely required)
- Domain expiration probe intervals not allowed to be lower than 30 minutes between requests.
  (See more info here: https://github.com/TwiN/gatus?tab=readme-ov-file#monitoring-domain-expiration)

## Globaly Provided Variables
We ship a couple of default secret variables to all gatus instances. These are available no matter if user defined variables are present or not.

### smtp 
We provide the following variables for email alert config:

- `DOLL_REPORT_SMTP_USERNAME`
- `DOLL_REPORT_SMTP_PASSWORD`
- `DOLL_REPORT_SMTP_HOST` 

Usage: (see https://github.com/TwiN/gatus?tab=readme-ov-file#configuring-email-alerts for full usage information)
```yaml
alerting:
  email:
    from: "{any name}@doll.report"
    username: ${DOLL_REPORT_SMTP_USERNAME}
    password: ${DOLL_REPORT_SMTP_PASSWORD}
    host: ${DOLL_REPORT_SMTP_PASSWORD}
    port: 587
    to: "myemail@goes.here"  # use a user-defined secret if you don't want to expose your email publically
    [ ... snip ... ]
```

## User-Defined Secret Variables

### Transparancy

> [!CAUTION]
> The following section is kind of important, so please read it in its entirety.

Thermia is the only one with access to the server and the age private key.
This will never intentionally change, however, incidents can happen. Please
take proper care to ensure that the information that exists within the environment files results in a mild annoyance at most if accessed by a third party.

Also note thermia will have to read the env files to ensure nothing malicious is being loaded into the containers. 

If there are any concerns don't hesitate to reach out to thermia, either through
already established channels or a github issue.

### `age` asymetrical encryption

To keep alert configuration private, users can use
[`age`](https://github.com/FiloSottile/age) to encrypt an environment file and
use the variables defined within it in their dashboard configuration.

There is a script to make this process easier in [`crypt/encrypt-envfile.sh`](crypt/encrypt-envfile). Make sure age is installed on the local system.

> [!NOTE]
> The resulting file needs to start with the name of the dashboard it should
> be attached to, and have the file ending `.env.age`.
>
> i.e `dashboards/myname.env.age` for `dashboards/myname.yaml`

More reading: https://github.com/TwiN/gatus?tab=readme-ov-file#alerting


### Example:
Assuming dashboard configuration is at `dashboards/example.yaml`, create an envfile at the repository root as follows

```bash
cat << EOF > local_envfile  # local_envfile is in .gitignore
DISCORD_WEBHOOK_URL="https://discord.com/api/webhooks/**********/**********"
EOF
```

Encrypt it using the script from the repository root, and redirect the output to the proper location:
```
./crypt/encrypt-envfile.sh local_envfile > ./dashboards/example.env.age
```

Finally, use the environment variable in the dashboard config:

```yaml
# [... snip ...]
alerting:
  discord:
    webhook-url: ${DISCORD_WEBHOOK_URL}  # environment variable
    title: ":ribbon: doll.report alert"
    default-alert:
      description: "health check failed"
      send-on-resolved: true
      failure-threshold: 5
      success-threshold: 5
# [... snip ...]
```

> [!NOTE]
> Make sure to only commit and push your encrypted environment file.
> `local_envfile` is in .gitignore for convenience


## Issues

Feel free to open issues on the repo if there are any bugs or questions!

---
[![doll.report 88x31](https://doll.report/img/doll-report.gif)](https://doll.report) [![thermia 88x31](https://doll.report/img/thermia.gif)](https://girlthi.ng/~thermia)
