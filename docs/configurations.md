Parameters
---

For any given parameter you can use either form field, form file or query strings.

* **template** (required): template content/file in HTML format. If you want to use template engine you can use [Golang template format](https://golang.org/pkg/text/template/) in your HTML.
* **data** (optional default empty): JSON data if you use template rather than plain HTML.
* **waitFor** (optional default body): Query string that engines uses to check your HTML is ready. It's useful when you generate your HTML using Javascript.
* **download** (optional default false): Force browser to download the PDF (not open).

Configurations
---

micro-pdf-generator works with zero configurations but if you need more customization you can set some environment variables:

* **MPG_TIMEOUT** (default 15): Default timeout to fetch remote template (using URL)
* **MPG_PORT** (default 8080)
* **MPG_HOST** (default 0.0.0.0)
* **MPG_TEMP_DIR** (default OS temp dir)
* **MPG_PORT_MAX_SIZE** (default 6): Maximum template size in MB