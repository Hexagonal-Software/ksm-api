## ksm-api

Wrapper for Keeper Secrets Manager to offer a REST connection to the Keeper Vault.
I wrote this app in order to more easily allow access to other applications where I cannot install the Keeper Secrets Manager app.

### Usage

Just clone the repository, download the depencies and run `make build` to create an executable. After building it, create a KSM_CONFIG env variable with your config and run the application.

There is also a `make release-docker` to create a Docker image that you can use.

### License

Please check the license file in the repository.

## Credits

All rights for the logos and the SDK go to KeeperSecurity. 

This project uses Go, Fiber, Cobra and Viper, so all credits go to their respective owners. Along with a big thank you from my part :)
