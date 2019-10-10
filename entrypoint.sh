#!/bin/sh
set -e

cp /usr/share/zoneinfo/${TIMEZONE} /etc/localtime
echo ${TIMEZONE} > /etc/timezone

/bin/sh -c ./cerise >> log/cerise.log 2>&1