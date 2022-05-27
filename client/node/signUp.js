const {SignUpRequest} = require("./users_pb.js");

module.exports = function signUp(client, signUpForm, callback) {
    const {name, email, password, confirmPassword, type} = signUpForm
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