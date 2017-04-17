Simulate HTTP errors. There are services to do this (e.g. httpstat.us), but
if you can't change the URL the system is calling, or need customization,
etc., this is one way to do it.

e.g. set your hosts file to handle some domain, to re-route traffic to this,
to simulate an error you're getting from a 3rd party API.

Run it like this (for it's default port 8000):

```
go run http_error_simulator.go
```

Or for example, to run it on a specific port, and in this case, port 80, which
requires sudo:

```
sudo PORT=80 go run http_error_simulator.go
```

This also, partially works with SSL, however it is a locally signed cert, and
thus will typically fail cert verification, unless you skip that. You can use
the -k flag with curl to ignore it, e.g.:

```
sudo PORT=443 go run http_error_simulator.go
curl -i -k https://localhost
```

See this article for more on how to generate the local certs and other
options: https://gist.github.com/denji/12b3a568f092ab951456

