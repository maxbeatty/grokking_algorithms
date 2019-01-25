const tap = require("tap");

function greedy(items, limit) {
  let curWeight = 0;
  const selected = [];

  items
    .sort((a, b) => a.value > b.value)
    .forEach(i => {
      if (curWeight + i.weight <= limit) {
        selected.push(i.name);
        curWeight += i.weight;
      }
    });

  return selected;
}

tap.deepEqual(
  greedy(
    [
      { name: "laptop", value: 2000, weight: 3 },
      { name: "stereo", value: 3000, weight: 4 },
      { name: "guitar", value: 1500, weight: 1 },
      { name: "mp3 player", value: 1000, weight: 1 },
      { name: "iphone", value: 2000, weight: 1 }
    ],
    4
  ),
  ["mp3 player", "guitar", "iphone"]
);

tap.deepEqual(
  greedy(
    [
      { name: "water", value: 10, weight: 3 },
      { name: "book", value: 3, weight: 1 },
      { name: "food", value: 9, weight: 2 },
      { name: "jacket", value: 5, weight: 2 },
      { name: "camera", value: 6, weight: 1 }
    ],
    6
  ),
  ["book", "jacket", "camera", "food"]
);

function knapsack(items, limit) {
  const matrix = new Array(items.length + 1);
  for (let i = 0; i < matrix.length; i++) {
    matrix[i] = new Array(limit + 1).fill(0);
  }

  for (let i = 1; i < items.length; i++) {
    for (let j = 1; j <= limit; j++) {
      /*
      If the item at the index matching the current row fits within the weight capacity represented by the current column, take the maximum of either:
      The total worth of the items already in the bag or,
      The total worth of all the items in the bag except the item at the previous row index, plus the new item’s worth
    */
      const curItem = items[i - 1];

      if (curItem.weight <= j) {
        const prev = matrix[i - 1][j]; // worth of existing items
        // value of current item + value of the remaining space
        const maybe = curItem.value + matrix[i - 1][j - curItem.weight];
        matrix[i][j] = Math.max(prev, maybe);
      } else {
        matrix[i][j] = matrix[i - 1][j];
      }
    }
  }

  let i = items.length;
  let w = limit; // weight limit
  const selected = [];

  // trace back the knapsack to see what items we're going to add
  while (i > 0 && w > 0) {
    const pick = matrix[i][w];

    // we know bottom rows have biggest value
    // if current row (pick) isn't the same value as prev row...
    if (pick !== matrix[i - 1][w]) {
      // we want to take the previous item
      selected.push(items[i - 1].name);
      // remaining weight is now...
      w -= items[i - 1].weight;
    }
    // keep looking up the table
    i--;
  }

  return selected;
}

tap.deepEqual(
  knapsack(
    [
      { name: "stereo", value: 3000, weight: 4 },
      { name: "laptop", value: 2000, weight: 3 },
      { name: "guitar", value: 1500, weight: 1 },
      { name: "mp3 player", value: 1000, weight: 1 },
      { name: "iphone", value: 2000, weight: 1 }
    ],
    4
  ),
  ["iphone", "mp3 player", "guitar"]
);

tap.deepEqual(
  knapsack(
    [
      { name: "water", value: 10, weight: 3 },
      { name: "book", value: 3, weight: 1 },
      { name: "food", value: 9, weight: 2 },
      { name: "jacket", value: 5, weight: 2 },
      { name: "camera", value: 6, weight: 1 }
    ],
    6
  ),
  ["camera", "food", "water"]
);
