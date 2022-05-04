const {AuthenticateRequest} = require("./users_pb.js");

module.exports = function authenticate(client, token, callback) {
    let request = new AuthenticateRequest();
    client.authenticate(request, {"authorization": token}, function(err, response) {
        if (err) {
            console.log(err.message);
            return callback(err.message, null);
        }
        return callback(null, response.authId);
    });
}