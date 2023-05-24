# Dynamic DNS for google DNS

## Usage
Configuration file **must** put in same directory with binary.

Sample configuration for `config.yaml`:

```yaml
provider: google

dns:
  - name: domain or subdomain name
    ip: auto
    username: <username located dynamic dns section>
    password: <password located dynamic dns section>

  - name: domain or subdomain name
    ip: auto
    username: <username located dynamic dns section>
    password: <password located dynamic dns section>
```

Command line:

`go build .`

`./dynamic-dns-service`

## Docker

Pre-built image available on: `registry.binhnguyen.dev/public/dynamic-dns-service`

Or build from command:

`docker build .` or `docker pull registry.binhnguyen.dev/public/dynamic-dns-service`

`docker compose up -d`