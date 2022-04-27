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