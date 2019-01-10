const tap = require("tap");

function findLowestIdx(list) {
  let low = list[0];
  let lowIdx = 0;

  list.forEach((val, idx) => {
    if (val < low) {
      low = val;
      lowIdx = idx;
    }
  });

  return lowIdx;
}

// pretend Array.sort() doesn't exist
function selectionSort(list) {
  const newArr = [];

  while (list.length) {
    // find index of lowest number
    const idx = findLowestIdx(list);

    newArr.push(list.splice(idx, 1)[0]);
  }

  return newArr;
}

tap.deepEqual(selectionSort([5, 3, 6, 2, 10]), [2, 3, 5, 6, 10]);
