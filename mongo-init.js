const collectionName = "users"
db.createCollection(collectionName)

usersCollection = db.getCollection(collectionName)
usersCollection.insertOne({
    name: "Initial User",
    email: "user@initial.com",
    password: "$2a$10$ZCcxLV4NCU94ezBxy.neEu66GXCpfg.QoQRLN6whyB6BJVaNYmwPe",
    birthDate: "1980-06-21",
    address: {
        street: "Rua Inicial, 0 - Centro",
        city: "Begin",
        state: "São Inicio",
        country: "Brasil"
    }
})