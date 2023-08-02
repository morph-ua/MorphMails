<p align="center">
  <a href="https://www.decline.live/" rel="noopener" target="_blank"><img width="300" src="/docs/public/logo.svg" alt="Morph Mails Logo"></a>
</p>

<hr/>

**Morph Mails** (*Formerly `Helium`*) is an open-source service that relies on [**Mailgun**](https://mailgun.com) and can integrate with _**any possible client**_. It creates an unlimited amount of random addresses, receives, parses, and sends a minified version of the letter to the user's chat.

<div align="center">

**[`v1.5.0`](https://github.com/morph-ua/MorphMails/releases/latest/)**

[![Latest Tag](https://ghcr-badge.egpl.dev/morph-ua/mails/latest_tag?color=red&label=latest)](https://ghcr.io/MorphMails/mails "Latest Tag")
[![Image Size](https://ghcr-badge.egpl.dev/morph-ua/mails/size)](https://ghcr.io/morph-ua/mails "Image Size")
![license](https://img.shields.io/github/license/morph-ua/MorphMails)
[![Average time to resolve an issue](https://isitmaintained.com/badge/resolution/morph-ua/MorphMails.svg)](https://isitmaintained.com/project/morph-ua/MorphMails 'Average time to resolve an issue')
[![Percentage of issues still open](http://isitmaintained.com/badge/open/morph-ua/MorphMails.svg)](http://isitmaintained.com/project/morph-ua/MorphMails "Percentage of issues still open")

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
  -p 8080:8080 -d ghcr.io/morph-ua/mails:v1.5.0-amd64 
```
> 📝 **Note**: You can generate the secret using this shell command:
> ```shell
> openssl rand -hex 16
> ```
* Next steps from the [**Network Setup**](#network-setup) section

#### Using Pre-built Binaries

* Install the [**latest version**](https://github.com/morph-ua/MorphMails/releases/latest) of MorphMails for your OS and architecture.
* Unpack the tarball.
* Open your Linux **Terminal** and run the following command:
```shell
DATABASE_URL=<your_db_dns> SECRET_KEY=<secret> ./morph_mails
```
* Next steps from the [**Network Setup**](#network-setup) section


### Network Setup

Here is the configuration guide for [Mailgun](https://mailgun.com) and your [VPS or VDS](https://hetzner.com) server.

- Login to your [Mailgun](https://mailgun.com) account and add your custom domain.
  Follow the instructions on the [Mailgun](https://mailgun.com) website to connect and verify your domain.
- Open the **Receiving** tab and create a new route. Set the following settings below:

```diff
+ Expression type: Catch All
+ Forward: <your-webhook-url>/sys/parse?token=SECRET_KEY_FROM_ENV
+ Priority: 0
```

- Host a client for your desired platform. You can find a list of clients
  [here](https://github.com/morph-ua/MailClients).

## Contributing

Read the [contributing guide](/CONTRIBUTING.md) to learn about our development process,
how to propose bug fixes and improvements,
and how to build and test your changes.

## Changelog

The [changelog](https://github.com/morph-ua/MorphMails/releases) is regularly updated
to reflect what's changed in each new release.

## Roadmap

Future plans and high-priority features and enhancements can be found in our [roadmap](https://github.com/orgs/morph-ua/projects/1).

## License

This project is licensed under the terms of [CC0-1.0](/LICENSE.md).

## Security

For details on supported versions and contact details for reporting security issues, please refer to the [security policy](https://github.com/morph-ua/MorphMails/security/policy).