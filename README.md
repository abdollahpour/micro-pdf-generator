[![Last release](https://img.shields.io/github/v/release/abdollahpour/micro-pdf-generator)](https://github.com/abdollahpour/micro-pdf-generator/releases/)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
![Docker Image Size (latest by date)](https://img.shields.io/docker/image-size/abdollahpour/micro-pdf-generator)
[![Coverage Status](https://coveralls.io/repos/github/abdollahpour/micro-pdf-generator/badge.svg?branch=master)](https://coveralls.io/github/abdollahpour/micro-pdf-generator?branch=master)
[![Build Status](https://secure.travis-ci.org/abdollahpour/micro-pdf-generator.svg?branch=master)](http://travis-ci.org/abdollahpour/micro-pdf-generator)

# micro-pdf-generator

Fast HTTP [microservice](http://microservices.io/patterns/microservices.html) written in Go for PDF generating. micro-pdf-generator can be used as a private or public HTTP service for massive HTML to pdf conversion. You can use query param, string, and URL as an input and go template engine to update input data as well. Ex

Here are some examples
```sh
# This is a serverless so you may have 5 secnods daly in your first call (could start)
# Also you PDF is gone after you get it once and you cannot use any external dependency in your HTML filesd (single HTML file format)
SERVER=http://micro-pdf-generator.demo.1.1.1.1.xio.io/pdf/sample.pdf

curl -F template="http://to-html-file" $SERVER
curl -F template=@local_html_file $SERVER
curl -F $SERVER?template=<html><body>Some_HTML</body></html>
curl -F template="http://raw.gitttttttttt -F data=@sample.json -F download=true -F waitFor=body $SERVER
```

More detils in here:

* [Setup using Serverless (Knative)](docs/knative.md)
* [Setup on Kubernetes](docs/kubernetes.md)
* [Setup using Docker](docs/docker.md)
* [Setup using Binary](docs/binary.md)
* [Configuration and parameters](docs/configurations.md)
* [Build and test](docs/build.md)

TODO
===
* Add more test
* Complete docs
* Complete templates
* Add system template for HTML errors
* Add JSON schema support
