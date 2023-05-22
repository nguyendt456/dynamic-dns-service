#!/bin/bash

sed -i "s/username:/username: ${DDNS_USERNAME}/g" ./config_test.yaml
sed -i "s/password:/password: ${DDNS_PASSWORD}/g" ./config_test.yaml
