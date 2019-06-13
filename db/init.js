db.createUser({
  user: "jl",
  pwd: "pwd",
  roles: [{ role: "readWrite", db: "pairot" }],
  mechanisms: ["SCRAM-SHA-256"]
});
