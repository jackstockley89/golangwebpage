#!/bin/bash

LINTER="quay.io/coreos/po-tooling"

lint_files() {
  if [ -x "$(command -v po-lint)" ]; then
    echo "Linting '${2}' files in directory '${1}'..."
    had_errors=0
    for file in $(find "${1}" -name "${2}"); do
      echo "${file}"
      po-lint "${file}"
      retval=$?
      if [ $retval -ne 0 ]; then
        had_errors=1
      fi
    done
    exit ${had_errors}
  elif [ -x "$(command -v docker)" ]; then
    echo "Using Dockerized linter."
    docker run --rm --volume "$PWD:/data:ro" --workdir /data ${LINTER} \
    /bin/bash -c "/go/bin/po-lint $1/$2"
  else
    echo "Linter executable not found."
    exit 1
  fi
}

lint_files "kubectl_deploy" "*.yaml"