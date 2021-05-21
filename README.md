[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
![Docker Image Size (latest by date)](https://img.shields.io/docker/image-size/abdollahpour/micro-pdf-generator)
[![Coverage Status](https://coveralls.io/repos/github/abdollahpour/micro-pdf-generator/badge.svg?branch=master)](https://coveralls.io/github/abdollahpour/micro-pdf-generator?branch=master)
[![Build Status](https://secure.travis-ci.org/abdollahpour/micro-pdf-generator.svg?branch=master)](http://travis-ci.org/abdollahpour/micro-pdf-generator)

# micro-pdf-generator

Fast HTTP [microservice](http://microservices.io/patterns/microservices.html) written in Go for PDF generating. micro-pdf-generator can be used as a private or public HTTP service for massive HTML to pdf conversion. You can use query param, string, and URL as an input and go template engine to update input data as well. Ex

Here are some examples
```sh
# This is a serverless so you may have 5 secnods daly in your first call (could start)
# Also you PDF is gone after you get it once
SERVER=http://micro-pdf-generator.demo.161.97.186.241.xip.io/pdf/sample.pdf

# Remote HTML file. Because of security you cannot use external resources, you have to embed them all (CSS, images, ...) in your HTML file
curl -F template="http://to-html-file" $SERVER

# Local HTML file
curl -F template=@docs/invoice-template.html $SERVER -o result.pdf

# HTML file using query string <html><body>Some_HTML</body></html>
curl "$SERVER?template=%3Chtml%3E%3Cbody%3ESome_HTML%3C%2Fbody%3E%3C%2Fhtml%3E" -o result.pdf

# Template file & JSON data
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