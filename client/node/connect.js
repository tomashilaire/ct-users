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
    let requester = {
        start: function(metadata, listener, next) {
            let newListener = {
                onReceiveMetadata: function(metadata, next) {
                    console.log("start onReceiveMetadata")
                    next(metadata);
                },
                onReceiveMessage: function(message, next) {
                    console.log("start onReceiveMessage")
                    next(message);
                },
                onReceiveStatus: function(status, next) {
                    console.log("start onReceiveStatus")
                    next(status);
                }
            };
            next(metadata, newListener);
        },
        sendMessage: function(message, next) {
            console.log("sendMessage")
            console.log(message)
            next(message);
        },
        halfClose: function(next) {
            console.log("halfClose")
            next();
        },
        cancel: function(message, next) {
            console.log("cancel")
            next();
        }
    };
    return new grpc.InterceptingCall(nextCall(options), requester);
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

