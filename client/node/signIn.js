const {SignInRequest} = require("./users_pb.js");

/**
 * @param {!proto.pb.Authentication.prototype} client The
 *     rpc client
 * @param body Payload with sign in credentials
 * @param {!string} body.email User email
 * @param {!string} body.password User password
 * @param callback
 *     call metadata
 *     callback The callback function(response)
 */
module.exports = function signIn(client, body, callback) {
    const {
        email,
        password
    } = body

    let request = new SignInRequest();
    request.setEmail(email);
    request.setPassword(password);

    client.signIn(request, {}, function(err, response) {
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
        let user = response.getUser();
        return callback({
            status: {
                code: 0,
                details: "Sign in successfully",
                metadata: response.metadata
            },
            data: {
                User: {
                    Id: user.getId(),
                    Name: user.getName(),
                    Email: user.getEmail(),
                    Created: user.getCreated(),
                    Updated: user.getUpdated()
                },
                Token: response.getToken()
            }
        });
    });
}