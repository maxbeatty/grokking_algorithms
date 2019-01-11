const tap = require("tap");

// countdown using recursion
function countdown(i) {
  console.log(i);

  return i <= 0 ? null : countdown(i - 1);
}

// stack function calls
const greet2 = name => console.log(`how are you, ${name}?`);

const bye = () => console.log("ok bye!");

const greet = name => {
  console.log(`hello, ${name}!`);
  greet2(name);
  console.log("getting ready to say bye...");
  bye();
};

// factorial using recursion
function fact(i) {
  return i === 0 || i === 1 ? 1 : i * fact(i - 1);
}

// TODO: add mocking lib to count console.log invocations
// tap.doesNotThrow(() => countdown(5));
// tap.doesNotThrow(() => greet("adit"));

tap.equal(fact(0), 1);
tap.equal(fact(5), 120);
tap.equal(fact(1), 1);
tap.equal(fact(3), 6);
