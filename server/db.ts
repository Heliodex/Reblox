
const db = new Surreal("http://localhost:8000/rpc")

try {
	// Signin as a namespace, database, or root user
	await db.signin({
		user: "root",
		pass: "root",
	})

	// Select a specific namespace / database
	await db.use("test", "test")

	// Create a new person with a random id
	const created = await db.create("person", {
		title: "Founder & CEO",
		name: {
			first: "Tobie",
			last: "Morgan Hitchcock",
		},
		marketing: true,
		identifier: Math.random().toString(36).substr(2, 10),
	})

	// Update a person record with a specific id
	const updated = await db.change("person:jaime", {
		marketing: true,
	})

	// Select all people records
	const people = await db.select("person")

	// Perform a custom advanced query
	const groups = await db.query("SELECT marketing, count() FROM type::table($tb) GROUP BY marketing", {
		tb: "person",
	})

	console.log("created\n", created)
	console.log("updated\n", updated)
	console.log("people\n", people)
	console.log("groups\n", groups)
} catch (e) {
	console.error("ERROR", e)
}
