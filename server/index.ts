import { serve } from "https://deno.land/std/http/server.ts"
import { gql } from "https://deno.land/x/graphql_tag/mod.ts"
import { GraphQLHTTP } from "https://deno.land/x/gql/mod.ts"
import { makeExecutableSchema } from "https://esm.sh/@graphql-tools/schema"

const typeDefs = gql`
	type Query {
		hello: String
		goodbye: String
	}
`

const resolvers = {
	Query: {
		hello: () => `Hello World!`,
		goodbye: () => `Goodbye World!`,
	},
}

const schema = makeExecutableSchema({ typeDefs, resolvers })

await serve(GraphQLHTTP({ schema, graphiql: true }), { port: 3000 })
