var PROTO_PATH = __dirname + '/../../internal/handlers/usersprotohdl/users.proto';

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

function connect(host, port) {
    return new user_proto.Authentication(host+":"+port,
        grpc.credentials.createInsecure());
}

function signUp(client, signUpForm, callback) {
    client.SignUp({
        name: signUpForm.name,
        email: signUpForm.email,
        password: signUpForm.password,
        confirmPassword: signUpForm.confirmPassword,
        type: signUpForm.type
    }, function(err, response) {
        if (err) {
            console.log(err.message);
            return callback("error");
        }
        return callback(response.user.id);
    });
}

exports.signUp = signUp;
exports.connect = connect;