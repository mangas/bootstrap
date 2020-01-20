## Proxy

## Proxy Transports

We need C2S and P2P actually. 

With Client2Server based its simpler and we can hide the Server behind Tor as an onion.

With P2P, we can also use Tor

All of them ensure no data persists on the device IF we want, Or a cache be be on the devices.


## Application 

We are building the equivalent of Gsuite to provide an alternative, so that users are not sharing their data with Google
Tor want it: https://support.torproject.org/tormessenger/

The data is stored on:

1. Device or
2. On ours Servers or
3. The Users home Rasp PI Server. Which is an exact copy of our Servers and multi-tenanted so that a User can share another friends Raspberry PI

## Auth

Keystore allows you to store your keys wither:

- OS TPM or
- Hardware TPM for those that want it ( 30 euro)

Hardware TPM

NOTE: Yubikey is cracked.
https://www.yubico.com/blog/infineon-rsa-key-generation-issue/

Support many of them: https://twofactorauth.org/

Nitrekey berlin

- Keypass integrated
- API: https://github.com/snapcore/snapd/blob/master/interfaces/builtin/u2f_devices.go
- Product: https://shop.nitrokey.com/shop/product/nitrokey-fido2-55
	- 30 euro each

USB Armory
- Provides Key Storage but cal also run our OS on the USB, and so you can move to any Computer and its 100% secure
- Web: http://inversepath.com/usbarmory
- The guy behind it is highly respects.
	- https://github.com/abarisani
		- he uses age: https://github.com/FiloSottile/age/issues/63
- They use golang :)
	- https://github.com/f-secure-foundry/interlock
		- Uses Signal: https://github.com/aebruno/textsecure
- Flutter GUI 
	- We can run all the Apps 




## FIDO2 and webauthn
WebAuthn Level 1
https://github.com/duo-labs/webauthn
Flutter Web users can use this.
Relay
https://github.com/e3b0c442/warp

FIDO2
https://github.com/GeertJohan/yubig
a pure-Go Yubico OTP Validation Client and is following the Yubico Validation Protocol Version 2.0.
SO from a Flutter Client

## Key Sign

OS based :
https://github.com/FiloSottile/age
doc: https://docs.google.com/document/d/11yHom20CrsuX8KQJXBBw04s80Unjv8zCg_A7sPAX_9Y/preview#
- ssh-rsa and ssh-ed25519 SSH public keys
- multiple recipients
- no need for traditional OS ssh local signing tools !!
- Fork that is able to be used as a lib: https://github.com/schollz/age/

Hardware based:
https://github.com/f-secure-foundry/tamago