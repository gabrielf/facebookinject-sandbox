# facebookinject-sandbox

Proof of concept web app using [facebook-inject](https://github.com/facebookarchive/inject) so that:

* The entire app can be created and run from a test
* Mocks are by default used for dependencies that usually talk to external systems
* Dependencies can be easily overridden from tests
* Dependencies can be added with little boilerplate

Example of how a new dependency is added can be found in [PR#1](https://github.com/gabrielf/facebookinject-sandbox/pull/1/files?diff=split&w=1).
