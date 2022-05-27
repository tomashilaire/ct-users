const {AuthenticateRequest} = require("./users_pb.js");

/**
 * @param {!proto.pb.Authentication.prototype} client The
 *     rpc client
 * @param body Payload authentication token
 * @param {!string} body.token Authentication token
 * @param callback
 *     call metadata
 *     callback The callback function(response)
 */
module.exports = function authenticate(client, body, callback) {
    const {
        token
    } = body;
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