#!/bin/bash

export MONGO_URI=mongodb://admin:admin_password@localhost:27017
export SIGNED_KEY=jsadl4vcxbei8923fv8v2jiucv1iujhu

go run ./cmd/http/main.go