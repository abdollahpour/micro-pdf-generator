You can generate PDFs fast and ralaible on fly!

    curl -F template="http://to-your-template?v1" -F data=@sample.json https://pdf-generator-address

And you have your PDF ready!

You can either pass template and JSON data as string, URL address or file in your POST request.

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