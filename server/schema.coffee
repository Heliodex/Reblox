import { gql } from "https://deno.land/x/graphql_tag/mod.ts"

export typeDefs = gql"""
	type Query {
		hello: String
	}
"""
