#!/usr/bin/env bash

set -e

# make sure that we don't miss a version when releasing
# usage: ./version.bash VERSION

if [[ $# -ne 1 ]]; then
  echo >&2 "usage: $(basename "$0") VERSION"
  exit 1
fi

releasedate="$(date '+%Y/%m/%d')"
label=org.opencontainers.image.version

sed -i "s|## unreleased|## v$1 ($releasedate)|" CHANGELOG.md
sed -i "s|const version = \".*\"|const version = \"$1\"|" main.go
sed -i "s|LABEL $label=.*|LABEL $label=v$1|" Dockerfile

exit 0
