# multiproxy: multiple environments per application

[![Go Report Card](https://goreportcard.com/badge/github.com/esnunes/multiproxy)](https://goreportcard.com/report/github.com/esnunes/multiproxy)

Multiproxy is an HTTP reverse proxy that allows and simplifies the creation of multiple environments (e.g. staging environments) for a single Facebook App / Shopify App / other platforms app.

## Overview

Platforms like Facebook and Shopify require a single endpoint to start the authentication handshake or a base URL to embed/open your application. This requirement makes tedious the process of creating multiple environments as they usually require manual setup of multiple application on those platforms.

Multiproxy provides a hassle-free solution for both users and developers. Using multiproxy users are able to easily select what environment they want to access and developers have to configure a single application on those platforms. Multiproxy also allows developers to create dynamic environments as they won't need to create an additional application for each environment.

## Concept

Multiproxy relies on a browser cookie to identify what environment is active at the moment and route requests accordingly. The authentication handshake and/or the app launch happens on a browser level, there is no direct communication between the app platform and the app during this phase.

<p align="center">
  <img width="400" src="https://user-images.githubusercontent.com/961874/48678015-b1217780-eb7d-11e8-906c-36c079e786c0.png" />
</p>

Some platforms like Shopify send events back to the applications using Webhooks, in this case, you have two options:

1. Set a specific webhook endpoint per environment (if possible);
1. Configure multiproxy to broadcast those Webhooks to all environments;

### Web Integrated Tool

Multiproxy comes with a built-in web interface for users to browse apps and select the active environment of each app. The web interface is just a nice interface to create the cookie used by multiproxy to route requests and set the proper value.

#### List of applications

<p align="center">
  <img width="700" alt="List of applications" src="https://user-images.githubusercontent.com/961874/48678068-a3202680-eb7e-11e8-8398-c84c0c4cd167.png" />
</p>

#### List of environments

<p align="center">
  <img width="700" alt="List of environments" src="https://user-images.githubusercontent.com/961874/48678070-ab786180-eb7e-11e8-8299-76baec572a1d.png" />
</p>

As you can configure the web interface in a different domain, the web interface will communicate with the app endpoint using XHR + Cross-Origin Request Sharing (CORS). Every time a user clicks to select an environment, the web interface makes an XHR to `app endpoint/_multiproxy` with the given environment key, this request creates/updates multiproxy cookie accordingly.

## Install

### Packages and Binary

```bash
# If you want to extend or use multiproxy packages
go get -u github.com/esnunes/multiproxy

# if you want to install multiproxy binary
go install -i github.com/esnunes/multiproxy/cmd/multiproxy
```

### Docker

```bash
docker run --rm -ti -p 8080:8080 -v $PWD/config.json:/app/config.json esnunes/multiproxy:latest
```

Check [Docker Hub](https://hub.docker.com/r/esnunes/multiproxy/) for the list of available tags.

## Getting Started

Due to the fact that multiproxy is focused on solving one problem, its configuration is simple and centered in a single JSON file.

Example (available also in [examples/config.json](https://github.com/esnunes/multiproxy/blob/master/examples/config.json)):

```js
{
  // URL of the Web Integrated Tool, scheme + host + path (required).
  admin: "http://admin.localhost:8080/",

  // Name of the cookie used by multiproxy to identify selected environment
  // (required).
  cookie: "multiproxy",

  // List of apps.
  apps: [
    {
      // ID of the application used by Web tool (required).
      id: 1,

      // Name of the app, used only by the Web tool to improve user experience
      // (required).
      name: "My app 1",

      // A brief description of the app, used only by the Web tool, you can
      // leave it blank (required).
      description: "This is my fancy app 1",

      // The base URL to access the app, this URL is used by multiproxy to map
      // requests to environments (required).
      addr: "http://app1.localhost:8080",

      // List of paths used by servers to communicate with your app, these
      // paths are used by multiproxy to send matched requests to all
      // environments, you can leave an empty array (required).
      broadcast: ["webhooks/"],

      // List of available environments (required).
      envs: [
        // Name of the environment (Web tool), unique key to identify
        // environment (cookie value), URL to proxy requests to.
        { name: "Environment 1", key: "env1", upstream: "http://localhost:8001" },
        { name: "Environment 2", key: "env2", upstream: "http://localhost:8002" },
        { name: "Environment 3", key: "env3", upstream: "http://localhost:8003" }
      ]
    },
    {
      id: 2,
      name: "My app 2",
      description: "Another fancy description",
      addr: "http://app2.localhost:8080",
      broadcast: ["webhooks/", "api/"],
      envs: [
        { name: "Environment 1", key: "env1", upstream: "http://localhost:8004" },
        { name: "Environment 2", key: "env2", upstream: "http://localhost:8005" }
      ]
    }
  ]
}
```

> The example above uses Javascript syntax just because it allows comments, however, the config file must be in JSON format.

## Authentication

Multiproxy does not provide any kind of authentication. There are projects like [Bit.ly OAuth2 Proxy](https://github.com/bitly/oauth2_proxy) that provide an easy to use/setup authentication layer.

## Roadmap

- [ ] Convert to Go 11 modules;
- [ ] Cover packages with tests;
- [ ] Better handle broadcast body;
- [ ] Reload config when file changed;
- [ ] Instrument (Prometheus);

## Contributing

We are open to, and grateful for, any contributions made by the community.

1. Fork it
1. Download your fork to your computer (`git clone https://github.com/your_username/multiproxy && cd multiproxy`)
1. Create your feature branch (`git checkout -b my-new-feature`)
1. Make changes and add them (`git add .`)
1. Commit your changes (`git commit -m 'Add some feature'`)
1. Push to the branch (`git push origin my-new-feature`)
1. Create new pull request

## License

Multiproxy is released under the MIT license. See [LICENSE](https://github.com/esnunes/multiproxy/blob/master/LICENSE)

## Disclaimers

The multiproxy tool, solutions, and opinions expressed here belong solely to the author, and not necessarily to the author's employer, organization, committee or other group or individual.

Multiproxy is not sponsored by any of the mentioned platforms.
