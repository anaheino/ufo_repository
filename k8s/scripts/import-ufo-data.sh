#!/bin/bash

# Config
NAMESPACE="ufo-app"
POD=$(kubectl get pods -n $NAMESPACE -l app=ufodb -o jsonpath="{.items[0].metadata.name}")
LOCAL_PORT=27017
REMOTE_PORT=27017
dbName="i_want_to_believe";
collectionName="sightings_with_coords";
myIp="127.0.0.1";
dataPath="./../../mongo/data.json";
myUserName="rust-scraper";
myPwd="iwanttobelieve";

echo "⏳ Port-forwarding MongoDB ($POD)..."
kubectl port-forward -n $NAMESPACE pod/$POD $LOCAL_PORT:$REMOTE_PORT &
PF_PID=$!

# Wait a moment for port-forward to establish
sleep 3

echo "📦 Importing data into MongoDB..."
mongoimportCommand="mongoimport --authenticationDatabase=admin --uri mongodb://$myUserName:$myPwd@$myIp:$LOCAL_PORT/$dbName --collection $collectionName --type json --jsonArray --file $dataPath"
eval $mongoimportCommand


# Kill the port-forward
echo "🧹 Cleaning up..."
kill $PF_PID

echo "✅ Import complete and connection closed."