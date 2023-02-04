export const typeDefs = `
  type Dinosaur {
    name: String
    description: String
  }

  type Query {
    dinosaurs: [Dinosaur]
        dinosaur(name: String): Dinosaur
  }
`
