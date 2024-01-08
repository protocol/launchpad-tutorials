import Client, { HTTP } from 'drand-client';
import fetch from 'node-fetch';
import AbortController from 'abort-controller';

global.fetch = fetch;
global.AbortController = AbortController;

const chainHash = '8990e7a9aaed2ffed73dbd7092123d6f289930540d7651336225dc172e51b2ce'; // Mainnet chain hash (hex encoded)
const urls = [
  'https://api.drand.sh',
  'https://drand.cloudflare.com'
]; // various endpoint options to access Drand network. 

const HEX = 16;
const FoodOptions = { "pho": 0.3, "croquets": 0.29, "pizza": 0.28, "pasta": 0.07, "molé_verde": 0.03, "shrimp": .03 };

// This function takes in a list of items and the probablilty of them being selected.
//    returns the number that is randomly selected 
async function weightedRandom(prob) {
  if (validateWeights(prob) != 1) {
    return "Weights not equal to 1"
  }
  const options = { chainHash };

  const client = await Client.wrap(HTTP.forURLs(urls, chainHash), options);

  // e.g. use the client to get the latest randomness round:
  const res = await client.get();

  // Get a number between 0 and 100 from Drand.
  const rand = randomPercentFrom(res.randomness);

  // This for loop selects which key:pair to return
  let sum = 0, r = rand;
  for (let [key, value] of Object.entries(prob)) {
    sum += value * 100;
    if (r <= sum) {
      return key;
    }
  }
}

//runs code above
weightedRandom(FoodOptions).then((lunch) => (console.log(lunch)))
// Press Ctrl + C to stop example!

// This function returns a number between 0 & 99.
function randomPercentFrom(randomness) {
  var i = 0;
  var randomDecimal = -1;
  var rand;

  while (randomDecimal <= 0 || randomDecimal > 100) {
    // Grab only 2 digits from randomness, in a "sliding window" fashion
    rand = randomness.slice(0 + i, 2 + i);
    // Convert hexadecimal randomness value to decimal (base16 -> base10)
    randomDecimal = parseInt(rand, HEX);
    i++;
  }
  return randomDecimal;
}

// This function checks the user input to make sure all probabilities add up to 1
function validateWeights(probabilities) {
  var sum = 0;
  for (let [_, value] of Object.entries(probabilities)) {
    sum += value;
  }
  return sum;
}