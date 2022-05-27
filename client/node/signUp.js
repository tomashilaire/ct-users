const {SignUpRequest} = require("./users_pb.js");

/**
 * @param {!proto.pb.Authentication.prototype} client The
 *     rpc client
 * @param body Payload with sign up information
 * @param {!string} body.name User name
 * @param {!string} body.email User email
 * @param {!string} body.password User password
 * @param {!string} body.confirmPassword User password confirmation
 * @param {!string} body.type User type
 * @param callback
 *     call metadata
 *     callback The callback function(response)
 */
module.exports = function signUp(client, body, callback) {
    const {
        name,
        email,
        password,
        confirmPassword,
        type
    } = body

    let request = new SignUpRequest();
    request.setName(name);
    request.setEmail(email);
    request.setPassword(password);
    request.setConfirmpassword(confirmPassword);
    request.setType(type);

    client.signUp(request, {}, function(err, response) {
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
                details: "Sign Up successfully",
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