const tap = require("tap");

function dijkstra(graph) {
  const costs = { fin: Infinity };
  const parents = { fin: null };

  for (const g in graph["start"]) {
    costs[g] = graph["start"][g];
    parents[g] = "start";
  }

  const processed = {};

  function findLowestCostNode() {
    let lowestCost = Infinity;
    let lowestCodeNode;

    for (const node in costs) {
      const cost = costs[node];

      if (cost < lowestCost && !processed[node]) {
        lowestCode = cost;
        lowestCodeNode = node;
      }
    }

    return lowestCodeNode;
  }

  let node = findLowestCostNode();

  while (node != null) {
    const cost = costs[node];
    const neighbors = graph[node];

    for (const n in neighbors) {
      const newCost = cost + neighbors[n];

      if (costs[n] > newCost) {
        costs[n] = newCost;
        parents[n] = node;
      }
    }

    processed[node] = true;

    node = findLowestCostNode();
  }

  return costs["fin"];
}

tap.equal(
  dijkstra({
    start: {
      a: 6,
      b: 2
    },
    a: {
      fin: 1
    },
    b: {
      a: 3,
      fin: 5
    },
    fin: {}
  }),
  6
);
