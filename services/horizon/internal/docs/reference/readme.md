---
title: Overview
---

Horizon is an API server for the DigitalBits ecosystem.  It acts as the interface between [digitalbits-core](https://github.com/digitalbitsorg/digitalbits-core) and applications that want to access the DigitalBits network. It allows you to submit transactions to the network, check the status of accounts, subscribe to event streams, etc. See [an overview of the DigitalBits ecosystem](https://developer.digitalbits.io/guides/) for details of where Horizon fits in. You can also watch a [talk on Horizon](https://www.youtube.com/watch?v=AtJ-f6Ih4A4) by Stellar.org developer Scott Fleckenstein:

[![Horizon: API webserver for the DigitalBits network](https://img.youtube.com/vi/AtJ-f6Ih4A4/sddefault.jpg "Horizon: API webserver for the DigitalBits network")](https://www.youtube.com/watch?v=AtJ-f6Ih4A4)

Horizon provides a RESTful API to allow client applications to interact with the DigitalBits network. You can communicate with Horizon using cURL or just your web browser. However, if you're building a client application, you'll likely want to use a DigitalBits SDK in the language of your client.
SDF provides a [JavaScript SDK](https://developer.digitalbits.io/js-digitalbits-sdk/learn/index.html) for clients to use to interact with Horizon.

SDF runs a instance of Horizon that is connected to the test net: [https://horizon.testnet.digitalbits.io/](https://horizon.testnet.digitalbits.io/) and one that is connected to the public DigitalBits network:
[https://horizon.livenet.digitalbits.io/](https://horizon.livenet.digitalbits.io/).

## Libraries

SDF maintained libraries:<br />
- [JavaScript](https://github.com/digitalbitsorg/js-digitalbits-sdk)
- [Java](https://github.com/digitalbitsorg/java-digitalbits-sdk)
- [Go](https://github.com/digitalbitsorg/go)

Community maintained libraries (in various states of completeness) for interacting with Horizon in other languages:<br>
- [Ruby](https://github.com/digitalbitsorg/ruby-digitalbits-sdk)

Community maintained Stellar libraries that can be forked for use on the DigitalBits network (in various states of completeness) for interacting with Horizon in other languages:<br>
- [Python](https://github.com/StellarCN/py-stellar-base)
- [C#](https://github.com/QuantozTechnology/csharp-stellar-base)
