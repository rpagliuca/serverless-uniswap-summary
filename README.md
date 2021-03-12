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
* Visit the `endpoint` on your browser. Remember to replace the end of the URL with your own Ethereum wallet address.
    * The endpoint should look something like: [https://random\_identifier.execute-api.us-east-1.amazonaws.com/dev/{walletAddress}](#)

## About the environment variables
* AWS\_PROFILE: used by `serverless` CLI locally to authenticate with `aws-cli`.
* ETHERSCAN\_API\_KEY: secret published to the AWS Lambda as env var. Register for a free api key at www.etherscan.io.
