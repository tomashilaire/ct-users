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

function connect(host, port) {
    return new user_proto.Authentication(host+":"+port,
        grpc.credentials.createInsecure());
}

function signUp(client, signUpForm, callback) {
    const {name, email, password, confirmPassword, type} = signUpForm
    client.SignUp({
        name: name,
        email: email,
        password: password,
        confirmPassword: confirmPassword,
        type: type
    }, function(err, response) {
        if (err) {
            console.log(err.message);
            return callback("error");
        }
        return callback(response.user.id);
    });
}

module.exports = {
    signUp,
    connect
}