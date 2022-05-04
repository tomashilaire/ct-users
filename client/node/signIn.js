const {SignInRequest} = require("./users_pb.js");
module.exports = function signIn(client, signInForm, callback) {
    const {email, password} = signInForm

    let request = new SignInRequest();
    request.setEmail(email);
    request.setPassword(password);

    client.signIn(request, {}, function(err, response) {
        if (err) {
            console.log(err.message);
            return callback(err.message, null, null);
        }
        return callback(null, response.getUser(), response.getToken());
    });
}