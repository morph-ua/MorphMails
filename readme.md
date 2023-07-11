# Atomic Emails - Art Module 📬

## Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)
- [Credits](#credits)

## Introduction

### What is Atomic Emails?

Atomic Emails is the service with a simple API that allows you to generate unlimited amount of random emails, connected to your desired [client](https://github.com/AtomicEmails/clients) (Telegram, WhatsApp, Discord, etc.). You can use them to register on any website or service and protect your real email address from spam, bots, and phishing.

### Why Atomic Emails?

Atomic Emails is a free and open-source project. You can host it on your own server and use it for free, or use hosted version of service with a lifetime subscription for only $3.00. We store only required data with retention of a week and don't sell it to third-party companies.

## Features

### Create a temporary email address

You can create a temporary email address by using the `/v2/assign` route or equivalent in your desired client. The service will generate a random email address for you and forward all
incoming emails to your client.

### Persistent email address

One persistent email address is created for each ID (Chat identificator in your client). You can use this email
address to receive emails from friends and family instead of using your main
email address. The service will parse and forward all incoming emails to your client.

### Connect your own domain

You can connect your own domain to the service by placing it on your own hosting to prevent the service from being blocked
by some websites. To do that, you need to install the service on your own machine and
follow the instructions in the [Installation](#installation) section.

## Installation

### Requirements

- VPS/VDS or your own machine with a public IP address.
- PostgreSQL database.
- Docker (Recommended).

### Setup

Using Docker:

- Open your terminal and run the following command:

```bash
$ docker run \
    -e DATABASE_URL=<your_db_url> \ 
    -e SECRET_KEY=<global_service_key_to_parse_emails> \
    -e PORT=8080 -p 8080:8080 -d ghcr.io/AtomicEmails/art:latest
```

- Next steps from [Network Setup](#network-setup) section.

Using Binaries:

- Install the latest version of [API Binary](https://github.com/AtomicEmails/art-module/releases/latest) for your OS.
- Run the following command:

```bash
$ DATABASE_URL=<your_db_url> \
    SECRET_KEY=<global_service_key_to_parse_emails> \
    PORT=8080 \
    ./art
```

- Next steps from [Network Setup](#network-setup) section.

### Network Setup

- Set up a webhook integration in your bot's settings. You can use
  [ngrok](https://ngrok.com/) to expose your local server to the internet.

```bash
$ ngrok http 8080
```

- Login to your mailgun account and add your custom domain. Follow the
  instructions on the mailgun website to verify your domain.
- Open Receiving tab and create a new route. Set this settings:

```diff
+ Expression type: Catch All
+ Forward: <your-webhook-url>/sys/parse?token=SECRET_KEY_FROM_ENV
+ Priority: 0
```

- Host a client for your desired platform. You can find a list of clients
  [here](https://github.com/AtomicEmails/clients).

## Usage

### Routes

- `/v2/register/:id` - Register a new user or connect a new client to existing one with the same ID.
- `/v2/assign/:id` - Generate a random string and make a record in Postgres to forward emails.
- `/v2/forward/:id` - Change forwarding state. This route disables/enables an email forwarding for user with specified ID.
- `/v2/delete/:id/:email` - Delete one email address from a user with specified ID. (Check documentation for examples)
- `/v2/reset/:id` - Delete all email addresses associated with a user with specified ID.
- `/v2/list/:id` - List all email addresses of the specified user.
- `/v2/paid/:id` - Switch the payment status. I use this route for development sponsoring, you can disable it if you want.
- `/sys/parse` - Route connected to Mailgun.
- `/sys/announcement` - Send an announcement to all registered users among all available clients.
- `/create/client` - Register new client. (Check documentation for examples)
- `/html/:id` - Get rendered HTML (Used to preview a letter)
- `/data/:id` - Get data of rendered letter (From, To, HTML). (Used to preview a letter)
- `/clients` - List all available clients.

## Contributing

Contributions are welcome! Please read the
[contribution guidelines](contributing.md) first.

## License

This project is licensed under the Attribution-NonCommercial 4.0 International License - see the
[BY-NC.md](by-nc.md) file for details.

## Credits

- [voxelin 🇺🇦](https://github.com/voxelin) - The creator of this project.
- [Mailgun](https://www.mailgun.com/) - A simple service for sending and
  receiving emails.
- [Go](https://golang.org/) - A great programming language.
- [Docker](https://www.docker.com/) - A great tool for containerizing
  applications.
- [ngrok](https://ngrok.com/) - An easy way for exposing local servers to the
  internet.