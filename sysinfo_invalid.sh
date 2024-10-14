#!/bin/bash
#
curl -X POST -H "Content-Type: application/json" -d '{"type": "invalid", "Payload": "windows11"}' http://localhost:8080/execute
