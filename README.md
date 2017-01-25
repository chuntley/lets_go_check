# Let's Go Check

A small Go server used for Let's Encrypt challenges on HaProxy load balanced servers. 

## Running

Install supervisord and use the sample config from below to run the Go server.

Using the sample post-hook below, create a shell script. Then in your cron, add the following:

`10 */12 * * * certbot renew --post-hook "/path/to/posthook.sh"`

## Options

- `port`: set the port of the server (default: 9009)

- `checkPath`: the path to check for the Let's Encrypt challenge (default: /var/www/letsencrypt)

## Sample Let's Encrypt post-hook

```
#!/usr/bin/env bash

cat /etc/letsencrypt/live/your-website.com/fullchain.pem \
    /etc/letsencrypt/live/your-website.com/privkey.pem \
    > /etc/letsencrypt/live/your-website.com/haproxy.pem
    
systemctl restart haproxy
```

## Sample HaProxy

```
frontend http_front
   bind *:80
   acl url_letsencrypt path_beg /.well-known
   use_backend letsencrypt_back if url_letsencrypt
   default_backend www_back
   
frontend https_front
   bind *:443 ssl crt /etc/letsencrypt/live/your-website.com/haproxy.pem
   default_backend www_back
   
backend letsencrypt_back
   server letsencrypt_check 127.0.0.1:9009 check
```

## Sample supervisord

```
[program:lets_go_check]
command=/root/lets_go_check
directory=/root
autostart=true
autorestart=true
startretries=10
redirect_stderr=true
stdout_logfile=/var/log/supervisor/lets_go_check.log
stdout_logfile_maxbytes=50MB
stdout_logfile_backups=10
```