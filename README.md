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

Run the following command (`-d` for running container in the background):

`docker compose up --build -d`