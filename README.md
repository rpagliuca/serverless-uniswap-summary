# About

* This is a Golang serverless app that makes use the module `github.com/rpagliuca/github.com/rpagliuca/go-uniswap-summary` to calculate some useful information for Uniswap liquidity provider.

* Due to API limits from Etherscan, this is not deployed publicly for the end user. You should deploy your own.

# Usage instructions
* Create an account at AWS
* Install `aws-cli` on your computer and set it up with your credentials
* Clone this repo
* Install the `serverless` framework (https://www.serverless.com/)
* Copy the file `env.json.dist` to `env.json` and enter your custom configuration
* Run the command `make deploy` (or perhaps `make build` + `serverless deploy`)
* The last command will show you the deployed lambda function `endpoint`
* Visit the `endpoint` on your browser
* It is a good idea to keep the `endpoint` secret, because it contains all your Uniswap LP positions openly

## About the environment variables

### LIQUIDITY_PROVIDER_POOL_POSITIONS
* This environment variable is a string tha represents:
    * Multiple positions separated by the pipe `|` character
        * 5 values for each position separated by comma `,`
            * Value 1: The token that represents the liquidity pool of choice. Should be DAI_WETH_LP or DAI_USDC_LP.
            * Value 2: The first token of the liquidity pool. Should be DAI, WETH or USDC.
            * Value 3: The initial quantity of your position of the first token.
            * Value 4: The second token of the liquidity pool. Should be DAI, WETH or USDC.
            * Value 5: The initial quantity of your position of the second token.
    * If you want to use tokens not provided by default (only DAI_WETH and DAI_USDC pools are provided):
        * Create custom `unisummary.Token`, one for each token that you wish (base tokens and pool tokens)
        * Customize `parse_config.go` to allow your custom types 

### Other variables
* The other environment variables are:
    * Your Etherscan API key (even the free account works)
    * Your Ethereum wallet address (which should contain some of the liquidity provider pool tokens used in the configuration). 
