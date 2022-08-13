# tinyurl

Simple url shortener in go with redis as datastore. Uses first 8 chars after Base58 encoding to get short url.

Following env variables need to be set -
  - REDIS_ADDRESS : Should be in format \<host\>:\<port\>
  - REDIS_PASSWORD : Password of redis db

Steps to run -
  - go run .
