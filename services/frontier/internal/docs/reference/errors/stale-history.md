---
title: Stale History
---

A frontier server may be configured to reject historical requests when the history is known to be further out of date than the configured threshold.  In such cases, this error is returned.  To resolve this error (provided you are the frontier instance's operator) please ensure that the ingestion system is running correctly and importing new ledgers.

## Attributes

As with all errors Frontier returns, `stale_history` follows the [Problem Details for HTTP APIs](https://tools.ietf.org/html/draft-ietf-appsawg-http-problem-00) draft specification guide and thus has the following attributes:

| Attribute | Type   | Description                                                                                                                     |
| --------- | ----   | ------------------------------------------------------------------------------------------------------------------------------- |
| Type      | URL    | The identifier for the error.  This is a URL that can be visited in the browser.                                                |
| Title     | String | A short title describing the error.                                                                                             |
| Status    | Number | An HTTP status code that maps to the error.                                                                                     |
| Detail    | String | A more detailed description of the error.                                                                                       |
| Instance  | String | A token that uniquely identifies this request. Allows server administrators to correlate a client report with server log files  |

## Example

```shell
$ curl -X GET "https://frontier.testnet.digitalbits.io/transactions?cursor=1&order=desc"
{
  "type": "stale_history",
  "title": "Historical DB Is Too Stale",
  "status": 503,
  "detail": "This frontier instance is configured to reject client requests when it can determine that the history database is lagging too far behind the connected instance of digitalbits-core.  If you operate this server, please ensure that the ingestion system is properly running.",
  "instance": "node0.frontier.testnet.digitalbits.io/ngUFNhn76T-078058"
}
```
