import Surreal from "https://deno.land/x/surrealdb/mod.ts"

db = new Surreal "http://localhost:8000/rpc"

try
	# Signin as a namespace, database, or root user
	await db.signin
		user: "root",
		pass: "root",

	# Select a specific namespace / database
	await db.use "test", "test" 

	# Create a new person with a random id
	created = await db.create "person",
		title: "Founder & CEO"
		name:
			first: "Tobie"
			last: "Morgan Hitchcock"
		marketing: true
		identifier: Math.random().toString(36).substring 2, 10

	# Update a person record with a specific id
	updated = await db.change "person:jaime",
		marketing: true,

	people = await db.select "person"

	groups = await db.query "SELECT marketing, count() FROM type::table($tb) GROUP BY marketing",
		tb: "person",

	console.log "created\n", created
	console.log "updated\n", updated
	console.log "people\n", people
	console.log "groups\n", groups
catch e
	console.error "ERROR", e

