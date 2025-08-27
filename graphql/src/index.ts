import { ApolloServer } from '@apollo/server';
import { startStandaloneServer } from '@apollo/server/standalone';

// A schema is a collection of type definitions (hence "typeDefs")
// that together define the "shape" of queries that are executed against
// your data.
const typeDefs = `#graphql
  # Comments in GraphQL strings (such as this one) start with the hash (#) symbol.

  type Character {
    name: String
    level: Int
    current: Boolean
  }
  
  type Item {
    name: String
    amount: Int
    selectable: Boolean
    description: String
  }

  type KeyItem {
    name: String
    description: String
  }

  type Magic {
    name: String
    description: String
    mp: Int
    selectable: Boolean
  }

  type EnemySkill {
    name: String
    description: String
    mp: Int
    obtained: Boolean
  }

  # The "Query" type is special: it lists all of the available queries that
  # clients can execute, along with the return type for each. In this
  # case, the "books" query returns an array of zero or more Books (defined above).
  type Query {
    books: [Book]
    characters: [Character]
    currentCharacters: [Character]
    items: [Item]
    keyItems: [KeyItem]
    magic: [Magic]
    enemySkill: [EnemySkill]
  }
`;

const characters = [
  {
    name: "Cloud",
    level: 48,
    current: true,
    currentHP: 3000,
    maxHP: 3000,
    currentMP: 264,
    maxMP: 473
  }, 
  {
    name: "Tifa",
    level: 44,
    currentHP: 3200,
    maxHP: 3200,
    currentMP: 118,
    maxMP: 393
  },
  {
    name: "Barret",
    level: 45,
    currentHP: 3334,
    maxHP: 3334,
    currentMP: 212,
    maxMP: 361    
  },
  { name: "Red XIII"},
  { name: "Yuffie" },
  { name: "Cid" },
  { name: "Vincent" },
  { name: "Cait Sith"}
]

const items = [
  { name: "Potion": amount: 0, selectable: true, description: ""}
  { name: "Soft", amount: 6, selectable: false, description: "Cures [Petrify]" },
  { name: "Turbo Ether", amount: 18, selectable: true, description: "Restores MP" },
  { name: "Ether", amount: 10, selectable: true },
  { name: "Phoenix Down", amount: 29, selectable: true },
  { name: "X-Potion", amount: 9, selectable: true },
  { name: "Remedy", amount: 1, selectable: false },
  { name: "Reagan Greens", amount: 3, selectable: false },
  { name: "M-Tentacles", amount: 5, selectable: false },
  { name: "Megalixir", amount: 6, selectable: true },
  { name: "Elixir", amount: 12, selectable: true }
]

const keyItems = [{ name: "Silk Dress", description: "Dress made of silk" },
{ name: "Blonde Wig", description: "" },
{ name: "Glass Tiara", description: "" },
{ name: "Cologne", description: "" },
{ name: "Key to Ancients", description: "" },
{ name: "Lunar Harp", description: "" },
{ name: "Basement Key", description: "" },
{ name: "PHS", description: "" },
{ name: "Gold Ticket", description: "" },
{ name: "Keystone", description: "" },
{ name: "Leviathan Scales", description: "" },
{ name: "Glacier Map", description: "" },
{ name: "Snowboard", description: "" }]

const magic = [
  { name: "Cure", description: "Restores HP", mp: 5, selectable: true },
  { name: "Cure2", description: "Restores HP", mp: 12, selectable: true },
  { name: "Cure3", description: "Restores HP", mp: 32, selectable: true },

  { name: "Regen", description: "Restores HP over time", mp: 37, selectable: false },
  { name: "Ultima", description: "Extreme Magic Attack", mp: 87, selectable: false },

  { name: "Fire", description: "", mp: 4, selectable: false },
  { name: "Fire2", description: "", mp: 12, selectable: false },
  { name: "Fire3", description: "", mp: 30, selectable: false },

  { name: "Ice", description: "", mp: 4, selectable: false },
  { name: "Ice2", description: "", mp: 12, selectable: false },
  { name: "Ice3", description: "", mp: 30, selectable: false },
  
  { name: "Bolt", description: "", mp: 4, selectable: false },
  { name: "Bolt2", description: "", mp: 12, selectable: false },
  { name: "Bolt3", description: "", mp: 30, selectable: false },
]

const enemySkill = [
  { name: "Frog Song", description: "Causes [Sleepel/Frog] on one opponent", mp: 5, obtained: true },
  { name: "L4 Suicide", description: "", mp: 0, obtained: true },
  { name: "Magic Hammer", description: "", mp: 0, obtained: true },
  { name: "White Wind", description: "", mp: 0, obtained: true },
  { name: "Big Guard", description: "", mp: 0, obtained: true },
  { name: "Death Force", description: "", mp: 0, obtained: true },
  { name: "Flame Thrower", description: "", mp: 0, obtained: true },
  { name: "Laser", description: "", mp: 0, obtained: true },
  { name: "Matra Magic", description: "", mp: 0, obtained: true },
  { name: "Bad Breath", description: "", mp: 0, obtained: true },
  { name: "Beta", description: "", mp: 0, obtained: true },
  { name: "Aqualung", description: "", mp: 0, obtained: true },
]


// Resolvers define how to fetch the types defined in your schema.
// This resolver retrieves books from the "books" array above.
const resolvers = {
  Query: {
    characters: () => characters,
    currentCharacters: () => characters.filter(x => x.current),
    items: () => items,
    keyItems: () => keyItems,
    magic: () => magic,
    enemySkil: () => enemySkill
  },
};

// The ApolloServer constructor requires two parameters: your schema
// definition and your set of resolvers.
const server = new ApolloServer({
  typeDefs,
  resolvers,
});

// Passing an ApolloServer instance to the `startStandaloneServer` function:
//  1. creates an Express app
//  2. installs your ApolloServer instance as middleware
//  3. prepares your app to handle incoming requests
const { url } = await startStandaloneServer(server, {
  listen: { port: 4000 },
});

console.log(`🚀  Server ready at: ${url}`);