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
    buildpack-packager build -any-stack
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

### Testing with cf local

You can test the build pack using cf local and docker

1. Download and install Docker Desktop

   [Docker Hub](https://hub.docker.com/search/?type=edition&offering=community)
   I used [Docker Desktop for Mac](https://hub.docker.com/editions/community/docker-ce-desktop-mac)

2. Install the Cloud Foundry CLI

   https://docs.cloudfoundry.org/cf-cli/install-go-cli.html

3. Install cf local

   ```bash
   cf install-plugin cflocal

4. Make a simple nodejs app

   You can use [mine](https://github.com/bryanateoan/nodejsHelloWorld)

5. Start docker

6. Stage with this buildpack, make sure you are in the test-cf-nodejs-app directory

   ```bash
   cf local stage helloWorld -b https://github.com/bryanateoan/project-nodejs-buildpack


   


### Acknowledgements

This project is a fork of the Cloud Foundry [Node buildpack](http://docs.cloudfoundry.org/buildpacks/node/index.html).
