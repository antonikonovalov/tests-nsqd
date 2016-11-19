#!/bin/bash
for i in {1..1000}
do
        curl -d "hello world $i" 'http://127.0.0.1:4151/put?topic=test'
        curl -d "hello world $((i+1000))" 'http://127.0.0.1:4153/put?topic=test'
done