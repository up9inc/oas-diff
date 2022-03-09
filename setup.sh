#!/bin/bash

echo "Cloning diff project to lib folder..."
mkdir lib || true && git clone --single-branch --branch feature/public_diff_functions https://github.com/up9inc/diff.git lib
