const tap = require("tap");

function euclid(w, h) {
  const short = w > h ? h : w;
  const long = w > h ? w : h;

  // if it divides equally, we know the answer
  if (long % short === 0) {
    return short;
  }

  // calculate new dimensions of remaining space and loop
  const fit = Math.floor(long / short);
  return euclid(long - short * fit, short);
}

tap.equal(euclid(1680, 640), 80);

function sum(arr) {
  if (arr.length === 1) {
    return arr[0];
  }

  return arr.pop() + sum(arr);
}

tap.equal(sum([2, 4, 6]), 12);

function count(arr) {
  if (arr.length === 0) {
    return 0;
  }

  return 1 + count(arr.slice(1));
}

tap.equal(count([1, 1, 1]), 3);

function max(arr) {
  if (arr.length === 1) {
    return arr[0];
  }

  let m = arr[0];
  let n = max(arr.slice(1));

  return m > n ? m : n;
}

tap.equal(max([2, 1, 3]), 3);
tap.equal(max([1, 5, 10, 25, 16, 1]), 25);

// recursive binary search
function binarySearch(list, item) {
  const mid = Math.floor((list.length - 1) / 2);
  const guess = list[mid];

  if (guess === item) {
    return mid;
  }

  if (guess > item) {
    return binarySearch(list.slice(0, mid), item);
  }

  return mid + 1 + binarySearch(list.slice(mid + 1), item);
}

tap.equal(binarySearch([1, 3, 5, 7, 9], 3), 1);
tap.equal(binarySearch([0, 2, 4, 6, 8, 10], 10), 5);

function qsort(arr) {
  if (arr.length < 2) {
    return arr;
  }

  const pivot = arr.splice(Math.floor((arr.length - 1) / 2), 1);

  const less = arr.filter(v => v <= pivot[0]);
  const greater = arr.filter(v => v > pivot[0]);

  return qsort(less).concat(pivot, qsort(greater));
}

tap.deepEqual(qsort([5, 2, 4, 3, 1]), [1, 2, 3, 4, 5]);
