#!/bin/bash
#
curl -X POST -H "Content-Type: application/json" -d '{"type": "ping", "Payload": "windows11"}' http://192.168.1.40:8080/execute
