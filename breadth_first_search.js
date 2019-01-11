const tap = require("tap");

function search(graph, name, check) {
  const q = [name];
  const searched = {};
  let person;

  while (q.length) {
    person = q.shift();

    if (!searched[person]) {
      for (let friend of graph[person]) {
        if (check(friend)) {
          return true;
        }

        searched[person] = true;
        q.push(...graph[person]);
      }
    }
  }

  return false;
}

const graph = {};
graph.you = ["alice", "bob", "claire"];
graph.bob = ["anuj", "peggy"];
graph.alice = ["peggy"];
graph.claire = ["thom", "jonny"];
graph.anuj = [];
graph.peggy = [];
graph.thom = [];
graph.jonny = [];

function check(str) {
  return str[str.length - 1] === "m";
}

tap.ok(search(graph, "you", check));
