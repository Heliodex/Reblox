import { getGraphQLRateLimiter } from "npm:graphql-rate-limit"
import db from "./db.js"

rateLimiter = getGraphQLRateLimiter identifyContext: (ctx) -> ctx.id
 
log = (x) ->
	console.log await x
	return x

limit = (max, window, callback) -> 
	(parent, args, context, info) ->
		errorMessage = await rateLimiter { parent, args, context, info }, { max, window }
		throw new Error errorMessage if errorMessage 
		callback()

resolvers =
	Query:
		users: limit 1, "1s", -> 
			log db.select "user"

	Mutation:
		create: limit 1, "10s", -> 
			log db.create "user",
				username: "Heliodex"
				password: "1234"

		delete: limit 1, "10s", -> 
			log db.delete "user"


export default resolvers
