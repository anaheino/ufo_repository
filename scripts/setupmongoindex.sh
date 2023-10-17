mongosh
use i_want_to_believe
db.createCollection("sightings_with_coords", { collation: { locale: 'en', strength: 2 }});
exit
mongoimport --authenticationDatabase=admin --uri mongodb://{myUserName}:{myPwd}@{myIp}:27017/i_want_to_believe --collection sightings_with_coords --type json --jsonArray --file ufo_data/sightings.json
db.sightings_with_coords.createIndex({"city": "text", "description": "text", "state": "text", "country": "text", "shape": "text", "duration": "text"}, { collation: { locale: "simple", strength: 2}});
