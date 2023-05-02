# Dynamic DNS for google DNS

## Usage
Configuration file **must** put in same directory with binary.

Sample configuration for `config.yaml`:

```yaml
provider: google

dns:
  - name: seproject.makerzone.net
    ip: auto
    username: <username located dynamic dns section>
    password: <password located dynamic dns section>

  - name: makerzone.net
    ip: auto
    username: <username located dynamic dns section>
    password: <password located dynamic dns section>
```

Command line:

`go build .`

`./dynamic-dns-service`