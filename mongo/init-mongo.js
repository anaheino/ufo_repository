const adminDB = db.getSiblingDB('admin');
// Create a user with roles in the admin database
adminDB.createUser({
    user: 'rust-scraper',
    pwd: 'iwanttobelieve',
    roles: [
        { role: 'userAdminAnyDatabase', db: 'admin' },
        { role: 'readWriteAnyDatabase', db: 'admin' },
    ],
});
const dbName = "i_want_to_believe";
const collectionName = "sightings_with_coords";
const targetDB = db.getSiblingDB(dbName);
if (!targetDB.getSiblingDB(dbName).getCollectionNames().length) {
    targetDB.getSiblingDB(dbName).createCollection(collectionName);
}
targetDB.sightings_with_coords.createIndex({"city": "text", "description": "text", "state": "text", "country": "text", "shape": "text", "duration": "text"}, { collation: { locale: "simple", strength: 2}});
