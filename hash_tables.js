const tap = require("tap");

const voted = {};

const checkVoter = name => {
  if (voted[name]) {
    return "kick them out!";
  } else {
    voted[name] = true;
    return "let them vote!";
  }
};

tap.equal(checkVoter("mike"), "let them vote!");
tap.equal(checkVoter("mike"), "kick them out!");
