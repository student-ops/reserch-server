print("Start #################################################################")
db = db.getSiblingDB("test")
db.createUser({
    user: "test",
    pwd: "password", // or cleartext password,
    roles: [{ role: "readWrite", db: "test" }],
})

db.createCollection("surroundings")
db.surroundings.insertMany([
    {
        temp: "13",
        press: "1000",
    },
    {
        temp: "35",
        press: "500",
    },
])

print("END #################################################################")

project / database
