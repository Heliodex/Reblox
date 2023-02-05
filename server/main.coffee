import { Server } from "https://deno.land/std/http/server.ts"
import { GraphQLHTTP } from "https://deno.land/x/gql/mod.ts"
import { makeExecutableSchema } from "https://esm.sh/@graphql-tools/schema"
import { typeDefs } from "./schema.js"
import { resolvers } from "./resolvers.js"

schema = makeExecutableSchema { resolvers, typeDefs }

server = new Server
	handler: (req) -> await GraphQLHTTP({schema, graphiql: true})(req)
	port: 3000,

console.log "Server running on port 3000"

server.listenAndServe()
