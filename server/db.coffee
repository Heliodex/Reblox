import Surreal from "https://deno.land/x/surrealdb/mod.ts"

db = new Surreal "http://localhost:8000/rpc"

console.log "Database connecting..."

try
	# Signin as a namespace, database, or root user
	await db.signin
		user: "root",
		pass: "root",

	# Select a specific namespace / database
	await db.use "test", "test" 
catch e
	console.error "ERROR", e

console.log "Database connected"

export default db
