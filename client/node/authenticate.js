let grpc = require('@grpc/grpc-js');

module.exports = function authenticate(client, token, callback) {
    let meta = new grpc.Metadata();
    meta.add("authorization", token)
    client.Authenticate({}, function(err, response) {
        if (err) {
            console.log(err.message);
            return callback(err.message);
        }
        return callback(response.authId);
    });
}