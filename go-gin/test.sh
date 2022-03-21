#!/bin/bash

post(){
    curl http://localhost:8080/albums \
        --include \
        --header "Content-Type: application/json" \
        --request "POST" \
        --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist":"Better Carter","price":49.99}'
}

get(){
    curl -v --connect-timeout 1 http://localhost:8080/albums --request "GET" --header "Content-Type: application/json"
}

get-id() {
    curl -v --location --connect-timeout 1 http://localhost:8080/albums/$2 --request "GET" --header "Content-Type: application/json"
}

$1
