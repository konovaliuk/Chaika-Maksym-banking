#!/bin/bash
set -e

cockroach sql --insecure --execute "CREATE DATABASE IF NOT EXISTS banking;"
