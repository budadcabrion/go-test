#!/bin/bash


curl --request GET "http://localhost:1234/setscore?name=alanis&score=1"
curl --request GET "http://localhost:1234/setscore?name=benny&score=10"
curl --request GET "http://localhost:1234/setscore?name=corey&score=11"
curl --request GET "http://localhost:1234/setscore?name=doug&score=23"
curl --request GET "http://localhost:1234/setscore?name=ethan&score=4"
curl --request GET "http://localhost:1234/setscore?name=francis&score=1"


curl --request GET "http://localhost:1234/getscores?start=0&count=10"
echo ""
curl --request GET "http://localhost:1234/getscores?start=1&count=2"
echo ""
curl --request GET "http://localhost:1234/getscores?start=10&count=10"
echo ""
curl --request GET "http://localhost:1234/getscores?start=3&count=3"
echo ""
