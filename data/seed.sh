#!/bin/bash

for file in *.json; do
    collection_name=$(basename "$file" .json)
    mongoimport --drop --uri=$MONGODB_URI --authenticationDatabase admin -c $collection_name --file $file --type json
    # mongoimport --drop --host mongo --port 27017 --username root --password root --authenticationDatabase admin -d cinema -c $collection_name --file $file --type json
    echo "Collection '$collection_name' imported successfully."
done