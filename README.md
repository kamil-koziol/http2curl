# http2curl

A minimal, lightweight tool to convert `.http` files into `curl` commands.

# Usage

```sh
Usage of http2curl:
  -args string
    	Args that will be passed to CURL
  -i string
    	Input file, use '-' for stdin (default "-")
  -o string
    	Output file, use '-' for stdout (default "-")
```

# Example

Consider the following HTTP request saved in `request.http`:

```http
GET https://api.example.com/data

Accept: application/json

Authorization: Bearer token123
```

Running the command:

```
http2curl -i request.http -args "-v"
```

produces the equivalent curl command:

```bash
/usr/bin/curl -v -X GET -H 'Accept: application/json' -d 'Authorization: Bearer token123' https://api.example.com/data
```

---

Let me know if you want it more formal, casual, or technical!


# Installation

`go install github.com/kamil-koziol/http2curl@latest`
