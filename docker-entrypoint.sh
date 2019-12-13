#!/bin/sh
set -e

# allow the container to be started with `--user`
if [[ "$1" = '/app/main' ]]; then
	chown -R app:app /app
	exec gosu app "$@"
fi

exec "$@"
