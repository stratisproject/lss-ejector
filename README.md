# Liquid Solo Staking ejector

`LSS Ejector` service plays an important role in Liquid Solo Staking. Every validator should run an ejector service to properly handle the validator exiting process.

## Building

### Requirements

- [Go 1.21.4+](https://go.dev/doc/install)
- Make - `sudo apt-get update && sudo apt-get install build-essential`

### Build or install

Build binary locally
```
make build-linux
```

or install systemwide
```
make install
```

### Running ejector service

1. Install supervisor
```
sudo apt-get install supervisor
```

1. Copy config content below to `/etc/supervisor/conf.d/ejector.conf`

```
[program:lss-ejector]
command=lss-ejector start
  --execution_endpoint=<execution_node_rpc>
  --consensus_endpoint=<beacon_node_rpc>
  --keys_dir=<path_to_keystore_folder>
  --staking_address=<lss_staking_contract_address>
environment=KEYSTORE_PASSWORD="<keystore_password>"
redirect_stderr=true
stdout_logfile=/var/log/supervisor-lss-ejector.log
stdout_logfile_maxbytes=500MB
stdout_logfile_backups=10
autostart=true
startretries=10
stopwaitsecs=60
```

1. Update supervisorctl config

```
sudo supervisorctl update
```

#### Useful commands

Start ejector service
```
sudo supervisorctl start lss-ejector
```

Stop ejector service
```
sudo supervisorctl stop lss-ejector
```

Restart ejector service
```
sudo supervisorctl restart lss-ejector
```

Watch ejector logs
```
tail -f /var/log/supervisor-lss-ejector.log
```

