#!/bin/sh

UNAME="$(command -v uname)" || exit 1
CHMOD="$(command -v chmod)" || exit 1
CP="$(command -v cp)" || exit 1

case $(${UNAME}) in
    [Ff][Rr][Ee][Ee][Bb][Ss][Dd])
        ${CP} -v bin/go-statsd-zabbix-freebsd-amd64 /usr/local/sbin/gostatsd
        ${CHMOD} +x /usr/local/sbin/gostatsd
        ${CP} -v daemon/freebsd/gostatsd /usr/local/etc/rc.d/gostatsd
        ${CHMOD} +x /usr/local/etc/rc.d/gostatsd
        ;;
    *)
        echo 'OS not implemented by this script!'
        exit 1
        ;;
esac
