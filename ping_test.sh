#!/bin/bash
#
curl -X POST -H "Content-Type: application/json" -d '{"type": "ping", "Payload": "windows11"}' http://localhost:8080/execute
