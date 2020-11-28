db.createUser(
    {
        user: "boryspil",
        pwd: "boryspil",
        roles: [
            {
                role: "readWrite",
                db: "boryspil_db"
            }
        ]
    }
)