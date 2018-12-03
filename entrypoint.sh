#!/usr/bin/env bash

set -e
set -u
set -o pipefail

# Se utilizan las mismas variables del administrativa_crud_api ya que este es una extensión del mismo
if [ -n "${PARAMETER_STORE:-}" ]; then
  export RESOLUCIONES_CRUD_DB_USER="$(aws ssm get-parameter --name /${PARAMETER_STORE}/administrativa_crud_api/db/username --output text --query Parameter.Value)"
  export RESOLUCIONES_CRUD_DB_PASS="$(aws ssm get-parameter --with-decryption --name /${PARAMETER_STORE}/administrativa_crud_api/db/password --output text --query Parameter.Value)"
fi

exec ./main "$@"
