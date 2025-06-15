#!/usr/bin/env bash
# wait-for-it.sh: Wait for a host and port to become available

host="$1"
port="$2"
shift 2

until nc -z "$host" "$port"; do
  echo "Waiting for $host:$port to be available..."
  sleep 1
done

exec "$@"

