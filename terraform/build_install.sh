#!/usr/bin/env bash

cd ../provider && go build && mv dashdog-provider ~/.terraform.d/plugins/terraform-provider-dashdog_v0.0.1 && cd ../terraform
