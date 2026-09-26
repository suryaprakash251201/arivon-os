#!/usr/bin/env bats

setup() {
    ARIVON_BIN="${ARIVON_BIN:-/usr/bin/arivon}"
}

@test "arivon binary exists" {
    run test -x "$ARIVON_BIN"
    [ "$status" -eq 0 ]
}

@test "arivon info returns version" {
    run "$ARIVON_BIN" info
    [ "$status" -eq 0 ]
    [[ "$output" == *"Version"* ]]
}

@test "arivon info returns kernel" {
    run "$ARIVON_BIN" info
    [ "$status" -eq 0 ]
    [[ "$output" == *"Kernel"* ]]
}

@test "arivon status returns healthy or degraded" {
    run "$ARIVON_BIN" status
    [ "$status" -eq 0 ]
    [[ "$output" == *"Overall"* ]]
}

@test "arivon status --json is valid JSON" {
    run "$ARIVON_BIN" status --json
    [ "$status" -eq 0 ]
    [[ "$output" == *"\"overall\""* ]]
}

@test "arivon status shows system section" {
    run "$ARIVON_BIN" status
    [ "$status" -eq 0 ]
    [[ "$output" == *"System"* ]]
}

@test "arivon status shows resources section" {
    run "$ARIVON_BIN" status
    [ "$status" -eq 0 ]
    [[ "$output" == *"Resources"* ]]
}

@test "arivon status shows services section" {
    run "$ARIVON_BIN" status
    [ "$status" -eq 0 ]
    [[ "$output" == *"Services"* ]]
}
