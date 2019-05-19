# Contrast Project Node.js Buildpack

Your project is to add a after-complile Hook to this buildpack. The hook can do anything you want, print a message, show the current weather, output CPU statistics.

Please write tests to validate your hook.

### Building the Buildpack

To build this buildpack, you'll need [Go](https://www.golang.org) installed. 

Once Go is installed, run the following commands from the buildpack's directory:

1. Source the .envrc file in the buildpack directory.

   ```bash
   source .envrc
   ```

1. Install buildpack-packager

    ```bash
     go install github.com/cloudfoundry/libbuildpack/packager/buildpack-packager
    ```

1. Build the buildpack

    ```bash
    buildpack-packager build -stack cflinux3
    ```

### Testing

To test this buildpack, run the following command from the buildpack's directory:

1. Source the .envrc file in the buildpack directory.

   ```bash
   source .envrc
   ```
   To simplify the process in the future, install [direnv](https://direnv.net/) which will automatically source .envrc when you change directories.

1. Run unit tests

    ```bash
    ./scripts/unit.sh


### Acknowledgements

This project is a fork of the Cloud Foundry [Node buildpack](http://docs.cloudfoundry.org/buildpacks/node/index.html).
