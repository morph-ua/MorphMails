<p align="center">
  <a href="https://www.decline.live/" rel="noopener" target="_blank"><img width="300" src="/docs/public/logo.svg" alt="Atomic Emails Logo"></a>
</p>

<hr/>

**Atomic Emails** is an open-source temporary email service that relies on [**Mailgun**](https://mailgun.com) and can integrate with _**any possible client**_.

- [_Atomic Emails Core_](https://github.com/AtomicEmails/AtomicEmails/) is a service itself which contains **Parser**, **Uploader**, and synchronizes all connected clients.

- [_AtomicEmails/telegram_](https://github.com/AtomicEmails/telegram/) is the first available client which I develop by myself written on [**Deno**](https://deno.land).

<div align="center">

**[`v2.0.1-h2o`](https://github.com/AtomicEmails/AtomicEmails/releases/latest/)**

[![Latest Tag](https://ghcr-badge.egpl.dev/AtomicEmails/app/latest_tag?color=red&label=latest)](https://ghrc.io/AtomicEmails/app "Latest Tag")
[![Image Size](https://ghcr-badge.egpl.dev/AtomicEmails/app/size)](https://ghrc.io/AtomicEmails/app "Image Size")
![license](https://img.shields.io/github/license/AtomicEmails/AtomicEmails)
[![Average time to resolve an issue](https://isitmaintained.com/badge/resolution/AtomicEmails/AtomicEmails.svg)](https://isitmaintained.com/project/AtomicEmails/AtomicEmails 'Average time to resolve an issue')
[![Percentage of issues still open](http://isitmaintained.com/badge/open/AtomicEmails/AtomicEmails.svg)](http://isitmaintained.com/project/AtomicEmails/AtomicEmails "Percentage of issues still open")

</div>

## Installation

### Requirements

* [**VPS/VDS**](https://hetzner.com) with public IP
* **Docker** (recommended)
* **PostgreSQL** database

### Setup

#### Using Docker

* Open your Linux **Terminal** and run the following command:
```shell
docker run \
  -e DATABASE_URL=<your_db_dsn> \
  -e SECRET_KEY=<secret> \
  -p 8080:8080 -d ghcr.io/atomicemails/app:v2.0.1-h2o-amd64 
```
> 📝 **Note**: You can generate the secret using this shell command:
> ```shell
> openssl rand -hex 16
> ```
* Next steps from [**Network Setup**](#network-setup) section.

#### Using Pre-built Binaries

* Install the [**latest version**](https://github.com/AtomicEmails/AtomicEmails/releases/latest) of Atomic Emails binary for your OS and Architecture
* Unpack the tarball
* Open your Linux **Terminal** and run the following command:
```shell
DATABASE_URL=<your_db_dns> SECRET_KEY=<secret> ./helium
```
* Next steps from [**Network Setup**](#network-setup) section.


### Network Setup

Here is the configuration guide for [Mailgun](https://mailgun.com) and your [VPS/VDS](https://hetzner.com) server.

- Login to your [Mailgun](https://mailgun.com) account and add your custom domain.
  Follow the instructions on the [Mailgun](https://mailgun.com) website to connect and verify your domain.
- Open the **Receiving** tab and create a new route. Set the following settings below:

```diff
+ Expression type: Catch All
+ Forward: <your-webhook-url>/sys/parse?token=SECRET_KEY_FROM_ENV
+ Priority: 0
```

- Host a client for your desired platform. You can find a list of clients
  [here](https://github.com/AtomicEmails/clients).

## Contributing

Read the [contributing guide](/CONTRIBUTING.md) to learn about our development process,
how to propose bug fixes and improvements,
and how to build and test your changes.

## Changelog

The [changelog](https://github.com/AtomicEmails/AtomicEmails/releases) is regularly updated
to reflect what's changed in each new release.

## Roadmap

Future plans and high-priority features and enhancements can be found in our [roadmap](https://github.com/orgs/AtomicEmails/projects/1).

## License

This project is licensed under the terms of the
[CC0-1.0](/LICENSE.md).

## Security

For details of supported versions and contact details for reporting security issues,
please refer to the [security policy](https://github.com/AtomicEmails/AtomicEmails/security/policy).
