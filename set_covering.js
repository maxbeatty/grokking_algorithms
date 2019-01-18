const tap = require("tap");

function setCovering(stations, statesNeeded) {
  const finalStations = [];

  while (statesNeeded.length) {
    let bestStation;
    let statesCovered = [];

    for (let s in stations) {
      const stationStates = stations[s];

      const covered = statesNeeded.filter(
        st => stationStates.indexOf(st) !== -1
      );

      if (covered.length > statesCovered.length) {
        bestStation = s;
        statesCovered = covered;
      }
    }

    statesNeeded = statesNeeded.filter(s => statesCovered.indexOf(s) === -1);
    finalStations.push(bestStation);
  }

  return finalStations;
}

tap.deepEqual(
  setCovering(
    {
      kone: ["id", "nv", "ut"],
      ktwo: ["wa", "id", "mt"],
      kthree: ["or", "nv", "ca"],
      kfour: ["nv", "ut"],
      kfive: ["ca", "az"]
    },
    ["mt", "wa", "or", "id", "nv", "ut", "ca", "az"]
  ),
  ["kone", "ktwo", "kthree", "kfive"]
);
