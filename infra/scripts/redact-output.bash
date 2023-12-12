#!/bin/bash

sed -e 's/AWS_SECRET_ACCESS_KEY".*/<REDACTED>/g' \
    -e 's/AWS_ACCESS_KEY_ID".*/<REDACTED>/g' \
    -e 's/$AWS_SECRET_ACCESS_KEY".*/<REDACTED>/g' \
    -e 's/$AWS_ACCESS_KEY_ID".*/<REDACTED>/g' \
    -e 's/\[id=.*\]/\[id=<REDACTED>\]/g'
