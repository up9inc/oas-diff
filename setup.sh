#!/bin/bash

echo "Cloning diff project to lib folder..."
mkdir lib || true && git clone --single-branch --branch feature/custom_differ_interceptor https://github.com/up9inc/diff.git lib

echo "Building oas-diff.."
make build
