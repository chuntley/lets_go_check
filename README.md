# Lets Go Check

A small Go server used for checking LetsEncrypt certificates. 
This was originally created to be used alongside HaProxy.

## Running

`./lets_go_check`

## Options

- `port`: set the port of the server (default: 9009)

- `checkPath`: the path to check for the LetsEncrypt challenge (default: /var/www/letsencrypt)

## Example

`./lets_go_check -port=8888 -checkPath=/path/to/check`