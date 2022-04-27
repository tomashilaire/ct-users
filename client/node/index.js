var PROTO_PATH = __dirname + '/../../internal/handlers/usersprotohdl/users.proto';
console.log(PROTO_PATH);
var parseArgs = require('minimist');
var grpc = require('@grpc/grpc-js');
var protoLoader = require('@grpc/proto-loader');
var packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    });
var user_proto = grpc.loadPackageDefinition(packageDefinition).pb;

function main() {
    var argv = parseArgs(process.argv.slice(2), {
        string: 'target'
    });
    var target;
    if (argv.target) {
        target = argv.target;
    } else {
        target = 'localhost:9092';
    }
    var client = new user_proto.Authentication(target,
        grpc.credentials.createInsecure());

    client.SignUp({
        name: "thilaire",
        email: "tomh6@hotmail.com",
        password: "7410",
        confirmPassword: "7410",
        type: "partner"
    }, function(err, response) {
        if (err) {
            console.log(err.message);
            return
        }
        console.log('response:', response);
        console.log('User ID:', response.user.id);
    });
}

function connect(host, port) {
    return new user_proto.Authentication(host+":"+port,
        grpc.credentials.createInsecure());
}

function signUp(client) {
    client.SignUp({
        name: "thilaire",
        email: "tomh6@hotmail.com",
        password: "7410",
        confirmPassword: "7410",
        type: "partner"
    }, function(err, response) {
        if (err) {
            console.log(err.message);
            return
        }
        console.log('response:', response);
        console.log('User ID:', response.user.id);
        return response.user.id;
    });
}

module.exports = connect
module.exports = signUp