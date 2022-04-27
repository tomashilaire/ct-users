module.exports = function signUp(client, signUpForm, callback) {
    const {name, email, password, confirmPassword, type} = signUpForm
    client.SignUp({
        name,
        email,
        password,
        confirmPassword,
        type
    }, function(err, response) {
        if (err) {
            console.log(err.message);
            return callback("error");
        }
        return callback(response.user.id);
    });
}