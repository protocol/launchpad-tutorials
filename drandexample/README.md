# drandexample
Example on how to use Drand as a client

# What, Why, How?
Created a small script to show & describe how to get random numbers from the drand network.

Example: Imagine you are a full time L5 software engineer and have more important things to think about than what to get for lunch.
    You decide to leave it up to randomness to choose your next meal. But you still have preferences.
    You assign weights to your preferences such that items you would like to eat most often have heavier weights (chances of being chosen)
    And things you don't want to eat as often, have smaller probability of being chosen.

Things to note:
    Drand mainnet releases a random number every 30 seconds. The problem that arises is if you want to test if the biased randomness works or not, it would take a really long time to test.
    There is work to shorten this time frame to 3 seconds, which is better, but its not as convenient as instant access of psuedo-random numbers like math.random() or crypto.getRandomValues().

This script shortens down a `randomness` value from the network and uses it as a probability for selecting an item.

### Requirements 
Node 12 and up

# Run
1. `npm install`

2. `node bias.js`
