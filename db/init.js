db.createUser({
  user: "jl",
  pwd: "vwdilab",
  roles: [{ role: "readWrite", db: "pairot" }],
  mechanisms: ["SCRAM-SHA-256"]
});
