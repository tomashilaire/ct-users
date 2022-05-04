const {SignUpRequest} = require("./users_pb.js");

module.exports = function signUp(client, signUpForm, callback) {
    const {name, email, password, confirmPassword, type} = signUpForm
    let request = new SignUpRequest();
    request.setName(name);
    request.setEmail(email);
    request.setPassword(password);
    request.setConfirmpassword(confirmPassword);
    request.setType(type);
    client.SignUp(request, {}, function(err, response) {
        if (err) {
            console.log(err.message);
            return callback(err.message, null);
        }
        return callback(null, response.user.id);
    });
}