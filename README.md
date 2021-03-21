[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
![Docker Image Size (latest by date)](https://img.shields.io/docker/image-size/abdollahpour/micro-pdf-generator)
[![Coverage Status](https://coveralls.io/repos/github/abdollahpour/micro-pdf-generator/badge.svg?branch=master)](https://coveralls.io/github/abdollahpour/micro-pdf-generator?branch=master)
[![Build Status](https://secure.travis-ci.org/abdollahpour/micro-pdf-generator.svg?branch=master)](http://travis-ci.org/abdollahpour/micro-pdf-generator)

# micro-pdf-generator

Fast HTTP [microservice](http://microservices.io/patterns/microservices.html) written in Go for PDF generating. micro-pdf-generator can be used as a private or public HTTP service for massive HTML to pdf conversion. You can use query param, string, and URL as an input and go template engine to update input data as well. Ex

```sh
curl -F template="http://to-html-file" https://pdf-generator-address/pdf/sample.pdf
curl -F template=@local_html_file https://pdf-generator-address/pdf/sample.pdf
curl -F https://pdf-generator-address/pdf/sample.pdf?template=<html><body>Some_HTML</body></html>
curl -F template="http://to-your-template" -F data=@sample.json -F download=true -F waitFor=body https://pdf-generator-address/pdf/sample.pdf
```

Parameters
---

For any given parameter you can use form field, form file or query string.

* **template** (required): template content in HTML format. If you want to use template engine you can use [Golang template format](https://golang.org/pkg/text/template/).
* **data** (optional default empty): JSON data if you use golang template in your HTML.
download (optional default false): force browser to download the PDF (not open).
* **waitFor** (optional default body): Query string that engine uses to wait for HTML to be ready.

Configurations
---

micro-pdf-generator without any configurations but if you need more customization you can set some environment variables:

* **timeout** (default 15): Default timeout to fetch remote template (using URL)
* **port** (default 8080)
* **host** (default 0.0.0.0)
* **temp_dir** (default OS temp dir)
* **max_size** (default 6): Maximum template size in MB

To build
===

    make compile

The binaries will be ready in `bin` directory.

To test
===

    make test

To print the test converage

    make coverage

To build docker image
===

    make name=<IMGE_NAME> image

TODO
===
* Add more test
* Complete docs
* Complete templates
* Add system template for HTML errors
* Add JSON schema support