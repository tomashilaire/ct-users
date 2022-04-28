let PROTO_PATH = __dirname + '/../../internal/handlers/usersprotohdl/users.proto';

let grpc = require('@grpc/grpc-js');
let protoLoader = require('@grpc/proto-loader');
let packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    });
let user_proto = grpc.loadPackageDefinition(packageDefinition).pb;

let interceptor = function(options, nextCall) {
    return new grpc.InterceptingCall(nextCall(options), {
        sendMessage: function(message, next) {
            console.log("intercepted");
            console.log(message);
            console.log(options.method_descriptor.name)
            next(message);
        }
    });
};

module.exports = function connect(host, port) {
    return new user_proto.Authentication(host+":"+port,
        grpc.credentials.createInsecure(),
        {
            "grpc.keepalive_time_ms": 10000,
            "grpc.keepalive_permit_without_calls": 1,
            interceptors: [interceptor],
        });
}

