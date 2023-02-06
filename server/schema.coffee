import { gql } from "https://deno.land/x/graphql_tag/mod.ts"

typeDefs = gql'
	type Query {
		users: [User]
	}

	type Mutation {
		create: User
		delete: String
	}

	type User {
		id: String
		username: String
		password: String
	}
'

export default typeDefs
