const adminDB = db.getSiblingDB('admin');
// Create a user with roles in the admin database
adminDB.createUser({
    user: 'rust-scraper',
    pwd: 'iwanttobelieve',
    roles: [
        { role: 'userAdminAnyDatabase', db: 'admin' },
        { role: 'readWriteAnyDatabase', db: 'admin' },
        // Add additional roles as needed
    ],
});
// Specify the database and collection names
const dbName = "i_want_to_believe";
const collectionName = "sightings_with_coords";
const targetDB = db.getSiblingDB(dbName);
if (!targetDB.getSiblingDB(dbName).getCollectionNames().length) {
    // Create the database if it doesn't exist
    targetDB.getSiblingDB(dbName).createCollection(collectionName);
}
targetDB.sightings_with_coords.createIndex({"city": "text", "description": "text", "state": "text", "country": "text", "shape": "text", "duration": "text"}, { collation: { locale: "simple", strength: 2}});
