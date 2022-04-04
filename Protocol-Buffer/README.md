# Protocol Buffer

* Protocol buffers are a method of serializing data that can be transmitted between multiple microservices in a platform-neutral way.

* The Protobuf is optimized to be faster than JSON and XML by removing many responsibilities usually done by data formats and making it focus only on the ability to serialize and deserialize data as fast as possible. Another important optimization is regarding how much network bandwidth is being utilized by making the transmitted data as small as possible.

* The Protobuf is a binary transfer format, meaning the data is transmitted as a binary. This improves the speed of transmission more than the raw string because it takes less space and bandwidth. Since the data is compressed, the CPU usage will also be less.

* The only disadvantage is the Protobuf files or data isn’t as human-readable as JSON or XML


## Installation

    $ brew install protobuf

## Generate Code using Protobuf

    $ mkdir js/
    $ protoc --js_out=js *.proto
