const {AuthenticateRequest} = require("./users_pb.js");

module.exports = function authenticate(client, token, callback) {
    let request = new AuthenticateRequest();
    client.authenticate(request, {"authorization": token}, function(err, response) {
        if (err) {
            return callback({
                status: {
                    code: err.code,
                    details: err.message,
                    metadata: err.metadata
                },
                data: null
            });
        }
        return callback({
            status: {
                code: 0,
                details: "Authenticated successfully",
                metadata: response.metadata
            },
            data: {
                AuthId: response.getAuthid()
            }
        });
    });
}