#!/bin/bash
curl -s https://api.thecatapi.com/v1/images/search?mime_types=jpg | jq -r '.[0].url' | xargs curl -s -o kot.jpg
img2txt kot.jpg
curl -s https://api.chucknorris.io/jokes/random | jq -r '.value'
