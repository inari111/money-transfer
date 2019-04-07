#!/bin/bash
set -eu

export $(cat .env | grep -v ^# | xargs)
