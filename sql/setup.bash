#!/bin/bash

psql -U demouser -d cyclingblog -a -f ../sql/create_table.sql
psql -U demouser -d cyclingblog -a -f ../sql/table_insert.sql

exit
