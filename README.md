# Airforce

Backend REST API for [SurfTimer](https://github.com/surftimer/SurfTimer), currently
being used in production for my server network [Wolf Airforce](https://wolf.airforce/).

## Requirements

### Basic

- Go 1.18
- MySQL or MariaDB

### Game Server

- [SurfTimer](https://github.com/surftimer/SurfTimer)
- [tVip](https://github.com/Totenfluch/tVip)

## Configuration

You may configure everything you need in the `config.json` file, with an example
shown in [`config.json.template`](./config.json.template).

## Building

### Docker

`docker build . -t ghcr.io/wolfairforce/backend:latest`

Then run with the following.

`docker run -v $(pwd)/config.json:/tmp/config.json -p 8080:8080 ghcr.io/wolfairforce/backend:latest`
