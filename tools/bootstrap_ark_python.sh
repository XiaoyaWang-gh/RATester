#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
TARGET_DIR="${ROOT}/.pydeps_ark"

rm -rf "${TARGET_DIR}"
mkdir -p "${TARGET_DIR}"

python3 -m pip install --quiet --target "${TARGET_DIR}" 'cryptography==43.0.3'

echo "Ark Python override installed to ${TARGET_DIR}"
echo "Use it via:"
echo "  export PYTHONPATH=${TARGET_DIR}:\$PYTHONPATH"
