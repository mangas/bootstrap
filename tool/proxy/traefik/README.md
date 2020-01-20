# traefik proxy

As a Dev and for the production system you need a Proxy to make it easier to configure things

# traefik
https://www.infoq.com/news/2019/11/traefik-routing-release/
https://containo.us/blog/traefik-2-0-6531ec5196c2/
- Code: https://github.com/containous/traefik

https://containo.us/maesh/
- Is just a Fancy Hosted version but is a good way to read the docs
- code: https://github.com/containous/maesh

- Replaces nginx
- love reload
- SNI ( Server name idenification) so secure
- Middleware so our code is not doing metrics
- Canaries, Feature flags
- Web Dashboard
- Works with K8 (CRD)

Golang Interpreter power this and is actually a YAML to GOLANG mapper.

https://github.com/containous/yaegi

We can use this for runtime things such as:

- Mage alternative
- Runtime logic embedded in the app itself like a JavaScript VM, but a golang VM.

## git tooling

https://github.com/containous/traefik/blob/master/docs/content/contributing/maintainers.md

https://github.com/containous/lobicornis/

- makes it much easier to manage PR's for maintainers and contribs


https://github.com/containous/bibikoffi/

- closes stale issues

https://github.com/containous/aloba

- manages lablels
- reports via slack 

