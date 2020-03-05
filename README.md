# go-statsd-zabbix

An implementation of statsd server to parse the metrics to the Zabbix Server.

### For build:
- Required: [Go compile >= 1.11](https://golang.org/doc/install)
```sh
# auto build
git clone https://github.com/willyrgf/go-statsd-zabbix.git
cd go-statsd-zabbix
./build.sh
```
Or
```sh
# manual build
git clone https://github.com/willyrgf/go-statsd-zabbix.git
cd go-statsd-zabbix
go build .
```

### For install:
```sh
git clone https://github.com/willyrgf/go-statsd-zabbix.git
cd go-statsd-zabbix
./install.sh
```

### Configure like a daemon in FreeBSD:
```sh
cat <<EOF >> /etc/rc.conf
# gostatsd
gostatsd_enable="YES"
gostatsd_storage="Zabbix"
gostatsd_storage_url="zabbix://192.168.10.187:10051"
EOF
```

### Enable the statsd module on Asterisk (version >= 13)
```
[general]
enabled = yes                   ; When set to yes, statsd support is enabled
server = 127.0.0.1              ; server[:port] of statsd server to use.
                                ; If not specified, the port is 8125
add_newline = no                ; Append a newline to every event. This is
                                ; useful if you want to run a fake statsd
                                ; server using netcat (nc -lu 8125)
```
