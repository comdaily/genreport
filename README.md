# genreport

## Based on

- github.com/go-rod/rod `v0.116.2`
- chromium revision `1321438`
- linux playwright `1124`

Prebuilt Docker base image: `pilinux/go:1.22.5-bookworm`

## Sample Docker Compose

```yaml
name: comdaily
services:
  genreport_dev:
    image: pilinux/go:1.22.5-bookworm
    container_name: genreport_dev
    working_dir: /genreport_dev
    restart: unless-stopped
    command: /genreport_dev/app
    environment:
      - TZ=Europe/Berlin
    ports:
      - '127.0.0.1:51011:8999'
    volumes:
      - ./app:/genreport_dev
      - /path/to/reports:/reports
```

## CURL

```shell
curl "http://ip:port/api/v1/create-pdf?orgID={orgID}&brandID={brandID}&reportID={reportID}&filename={filename.pdf}"
```
