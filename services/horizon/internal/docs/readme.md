---
title: Horizon
---

Horizon is the server for the client facing API for the DigitalBits ecosystem.  It acts as the interface between [digitalbits-core](https://developer.digitalbits.io/learn/digitalbits-core) and applications that want to access the DigitalBits network. It allows you to submit transactions to the network, check the status of accounts, subscribe to event streams, etc. See [an overview of the DigitalBits ecosystem](https://developer.digitalbits.io/guides/) for more details.

You can interact directly with horizon via curl or a web browser but SDF provides a [JavaScript SDK](https://developer.digitalbits.io/js-digitalbits-sdk/learn/) for clients to use to interact with Horizon.

SDF runs a instance of Horizon that is connected to the test net [https://horizon.testnet.digitalbits.io/](https://horizon.testnet.digitalbits.io/).

## Libraries

SDF maintained libraries:<br />
- [JavaScript](https://github.com/digitalbitsorg/js-digitalbits-sdk)
- [Java](https://github.com/digitalbitsorg/java-digitalbits-sdk)
- [Go](https://github.com/digitalbitsorg/go)

Community maintained libraries (in various states of completeness) for interacting with Horizon in other languages:<br>
- [Ruby](https://github.com/digitalbitsorg/ruby-digitalbits-sdk)


Community maintained libraries of Stellar that can be forked to be used on the DigitalBits network (in various states of completeness) for interacting with Horizon in other languages:<br>
- [Python](https://github.com/StellarCN/py-stellar-base)
- [C# .NET 2.0](https://github.com/QuantozTechnology/csharp-stellar-base)
- [C# .NET Core 2.x](https://github.com/elucidsoft/dotnetcore-stellar-sdk)
- [C++](https://bitbucket.org/bnogal/stellarqore/wiki/Home)
