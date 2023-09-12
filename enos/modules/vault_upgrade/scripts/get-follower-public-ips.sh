#!/bin/bash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: BUSL-1.1


set -e

binpath=${VAULT_INSTALL_DIR}/vault
export VAULT_ADDR="http://localhost:8200"

instances=${VAULT_INSTANCES}

# Find the leader
leader_address=$($binpath status -format json | jq '.leader_address | rtrimstr(":8200") | ltrimstr("http://")')

# Get the public ip addresses of the followers
follower_ips=$(jq ".[] | select(.private_ip!=$leader_address) | .public_ip" <<< "$instances")

echo "$follower_ips" | sed 's/\"//g' | tr '\n' ' '
