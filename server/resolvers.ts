const dinosaurs = [
	{
		name: "Aardonyx",
		description: "An early stage in the evolution of sauropods.",
	},
	{
		name: "Abelisaurus",
		description: '"Abel\'s lizard" has been reconstructed from a single skull.',
	},
]

export const resolvers = {
	Query: {
		dinosaurs: () => {
			console.log("dinosaurs")
			return dinosaurs
		},
		dinosaur: (_: any, args: any) => {
			console.log("dinosaur")
			return dinosaurs.find(dinosaur => dinosaur.name === args.name)
		},
	},
}
