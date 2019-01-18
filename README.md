# README
Simple kafka consumer which is a kind of wrapper of [sarama](https://github.com/Shopify/sarama).

## Getting started
Copy `kafka-cli.yaml.templ` into `$HOME/.config/kafka-cli.yaml`.

Replace each value for `topic`, `hosts`, `data-cert` and `ca-cert` with yours.

```
$ kafka-cli consume
```
