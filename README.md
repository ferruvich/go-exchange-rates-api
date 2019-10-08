# Exchange Rates API

## Specification

Write a golang/PHP application that is *ready to be deployed*.

The application should:

- Have an endpoint that calls [this api](https://exchangeratesapi.io/) to get the latest exchange rates for the base currencies of GBP and USD;
- It should return the value of 1 GBP or 1 USD in euros;
- It should check that value against the historic rate for the last week and make a naive recommendation as to whether this is a good time to exchange money or not.
