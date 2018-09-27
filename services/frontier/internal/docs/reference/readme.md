---
title: Overview
---

Frontier is an API server for the DigitalBits ecosystem.  It acts as the interface between [digitalbits-core](https://github.com/digitalbitsorg/digitalbits-core) and applications that want to access the DigitalBits network. It allows you to submit transactions to the network, check the status of accounts, subscribe to event streams, etc. See [an overview of the DigitalBits ecosystem](https://developer.digitalbits.io/guides/) for details of where Frontier fits in. 



Frontier provides a RESTful API to allow client applications to interact with the DigitalBits network. You can communicate with Frontier using cURL or just your web browser. However, if you're building a client application, you'll likely want to use a DigitalBits SDK in the language of your client.
SDF provides a [JavaScript SDK](https://developer.digitalbits.io/js-digitalbits-sdk/reference/index.html) for clients to use to interact with Frontier.

SDF runs a instance of Frontier that is connected to the test net: [https://frontier.testnet.digitalbits.io/](https://frontier.testnet.digitalbits.io/) and one that is connected to the public DigitalBits network:
[https://frontier.livenet.digitalbits.io/](https://frontier.livenet.digitalbits.io/).

## Libraries

SDF maintained libraries:<br />
- [JavaScript](https://github.com/digitalbitsorg/js-digitalbits-sdk)
- [Java](https://github.com/digitalbitsorg/java-digitalbits-sdk)
- [Go](https://github.com/digitalbitsorg/go)

Community maintained libraries (in various states of completeness) for interacting with Frontier in other languages:<br>
- [Ruby](https://github.com/digitalbitsorg/ruby-digitalbits-sdk)

Community maintained Stellar libraries that can be forked for use on the DigitalBits network (in various states of completeness) for interacting with Frontier in other languages:<br>
- [Python](https://github.com/StellarCN/py-stellar-base)
- [C#](https://github.com/QuantozTechnology/csharp-stellar-base)

## Additional Resources


You can also watch a [talk on Horizon](https://www.youtube.com/watch?v=AtJ-f6Ih4A4) by Stellar.org developer Scott Fleckenstein, which is similar to Frontier:

[![Horizon: API webserver](https://img.youtube.com/vi/AtJ-f6Ih4A4/sddefault.jpg "Horizon: API webserver")](https://www.youtube.com/watch?v=AtJ-f6Ih4A4)
