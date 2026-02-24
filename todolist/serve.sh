#!/bin/bash

caddy file-server --browse --listen :8080 --proxy /api/* localhost:3002/api