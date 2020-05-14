# HMAC Demo

A simple demo to show Go <> Python interop demo for creating signed JSON payloads

## Getting Started

### Python

```bash
$ pipenv --three
$ pipenv update
...
(python) bash-3.2$ python demo.py
Encoding: b'{"foo": "bar"}'
Matched!
```

### Go

```bash
go run demo.go
2020/05/14 16:35:28 Secret: CLIENT_SECRET Data: {"foo": "bar"}
2020/05/14 16:35:28 String Literal Bytes: [123 34 102 111 111 34 58 32 34 98 97 114 34 125]
2020/05/14 16:35:28 String Literal Result: 025e27785260c68eea22af4aca0b14ff6d41b80986beb26a10b384ca5e14408b
2020/05/14 16:35:28 Secret: CLIENT_SECRET Data: {"foo":"bar"}
2020/05/14 16:35:28 Marshalled Bytes: [123 34 102 111 111 34 58 34 98 97 114 34 125]
2020/05/14 16:35:28 Marshalled Result: 6ee84d2105efe2f41f1e403295679505577c123401204021ac290e2a2a55f542
```

## Caveat

Encoding whitespace in the JSON payload will cause issues. Go uses compact format `{"foo":"bar"}`
whereas Python will inject a space prior to the value e.g. `{"foo": "bar"}`. This difference
is enough to throw out the signature so the workaround I chose was to modify the Python
json.dump seperators as follows:

```python
json.dumps(data,separators=(',', ':'))
```
