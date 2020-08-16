#!/bin/bash
MNEMONIC=$1
ADDRESS=$2

ganache-cli      \
  -h 0.0.0.0     \
  -i 420         \
  -m "$MNEMONIC" \
  -u $ADDRESS