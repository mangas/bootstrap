# tor proxy

Allows a multi hop way to connect and hide your IP via the many hops you jump through an exit to the public internet

Runs locally on 127.0.0.1:9150 as a socks proxy

UDP is NOT supported.

## Domain Fronting
https://trac.torproject.org/projects/tor/wiki/doc/Snowflake

Domain fronting only used for brief Signalling / NAT-piercing to setup P2P webrtc which handles the actual traffic.

Which is EXACTLY what we run now for WebRTC TURN and STUN.

How it works:

1. Volunteers visit websites which host the "snowflake" proxy. 
2. Tor clients automatically find available browser proxies over the
domain fronted signaling channel.
3. Tor client and Proxy proxy establish a WebRTC peer connection.
4. Proxy connects to some relay.
5. Tor occurs.

## Bridge

RISEUP ( ie XR run a sponsorship and i bet they use it)

https://bridges.torproject.org/
91.194.60.100:9001 8D31A74671EACA6B3865DDBD6A5DD70D7F6A24F6
46.98.74.100:9001 8FA326EC142A0156E5171780550189388A5B33B4
QR CODE: SEE Drive

## Browser

https://www.torproject.org/download/
For all Os’s but not ios.
Grab one.
Test it to see its REALLY hiding you
https://www.whatsmyip.org/
https://www.whatsmyip.org/traceroute/


## Sec Web Site

https://2019.www.torproject.org/getinvolved/volunteer

All the things here WE should have for GetCourageNow !

We are telling everyone we run Tor and also asking for coders. This is a good combo because they know we are serious.
