#!/bin/bash
dbName="i_want_to_believe";
collectionName="sightings_with_coords";
dataPath="data.json";
myUserName="rust-scraper";
myPwd="iwanttobelieve";
myIp="127.0.0.1";
port="27017"
mongoimportCommand="mongoimport --authenticationDatabase=admin --uri mongodb://$myUserName:$myPwd@$myIp:$port/$dbName --collection $collectionName --type json --jsonArray --file $dataPath"
eval $mongoimportCommand
