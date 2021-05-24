[![Last release](https://img.shields.io/github/v/release/abdollahpour/micro-pdf-generator)](https://github.com/abdollahpour/micro-pdf-generator/releases/)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
![Docker Image Size (latest by date)](https://img.shields.io/docker/image-size/abdollahpour/micro-pdf-generator)
[![Coverage Status](https://coveralls.io/repos/github/abdollahpour/micro-pdf-generator/badge.svg?branch=master)](https://coveralls.io/github/abdollahpour/micro-pdf-generator?branch=master)
[![Build Status](https://secure.travis-ci.org/abdollahpour/micro-pdf-generator.svg?branch=master)](http://travis-ci.org/abdollahpour/micro-pdf-generator)

# micro-pdf-generator

Fast HTTP [microservice](http://microservices.io/patterns/microservices.html) written in Go for PDF generating. micro-pdf-generator can be used as a private or public HTTP service for massive HTML to pdf conversion. For example:

```sh
curl \
  -F template="https://raw.githubusercontent.com/abdollahpour/micro-pdf-generator/master/docs/template.html" \
  -F data="https://raw.githubusercontent.com/abdollahpour/micro-pdf-generator/master/docs/data.json" \
   http://micro-pdf-generator.demo.161.97.186.241.sslip.io/pdf/sample.pdf -o sample.pdf
```

It uses [Go template format](https://golang.org/pkg/text/template/) but you can also use normal standalone html (for template) and don't pass data.
You can also use url, file and string for both data and template parameters (check [configurations](docs/configurations.md))

# More

- [Setup using Serverless (Knative)](docs/knative.md)
- [Setup using Docker](docs/docker.md)
- [Configuration and parameters](docs/configurations.md)

# TODO

- Add more test
- Complete docs
- Complete templates
- Add system template for HTML errors
- Add JSON schema support
