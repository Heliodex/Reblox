export resolvers =
	Query:
		hello: ->
			console.log("hello")
			return "hello world!"
