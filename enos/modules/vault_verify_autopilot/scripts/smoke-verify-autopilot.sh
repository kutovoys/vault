#!/bin/bash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: BUSL-1.1


autopilot_version="${VAULT_AUTOPILOT_UPGRADE_VERSION}"
autopilot_status="${VAULT_AUTOPILOT_UPGRADE_STATUS}"

export VAULT_ADDR="http://localhost:8200"
[[ -z "$VAULT_TOKEN" ]] && fail "VAULT_TOKEN env variable has not been set"

function fail() {
  echo "$1" 1>&2
  exit 1
}

count=0
retries=7
while :; do
    state=$("${VAULT_INSTALL_DIR}"/vault read -format=json sys/storage/raft/autopilot/state)
    status="$(jq -r '.data.upgrade_info.status' <<< "$state")"
    target_version="$(jq -r '.data.upgrade_info.target_version' <<< "$state")"

    if [ "$status" = "$autopilot_status" ] && [ "$target_version" = "$autopilot_version" ]; then
        exit 0
    fi

    wait=$((2 ** count))
    count=$((count + 1))
    if [ "$count" -lt "$retries" ]; then
        echo "$state"
        sleep "$wait"
    else
        fail "Autopilot did not get into the correct status"
    fi
done
