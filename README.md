# Contrast Project Node.js Buildpack

Thank you for your interest in joining Contrast Security.

Your project is to:

* Fork or clone this repository 
* Add an after-complile Hook to this buildpack. The hook can do anything you want, print a message, show the current weather, output CPU statistics.
* Write tests to validate your hook.
* Send us a link to your repository


### Building the Buildpack

To build this buildpack, you'll need [Go](https://www.golang.org) installed. I recommend 1.12 or later. The buildpack will not build on 1.10 or earlier. 

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
1. Run unit tests

    ```bash
    ./scripts/unit.sh


### Acknowledgements

This project is a fork of the Cloud Foundry [Node buildpack](http://docs.cloudfoundry.org/buildpacks/node/index.html).
