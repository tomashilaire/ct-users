module.exports = function signIn(client, signInForm, callback) {
    const {email, password} = signInForm
    client.SignIn({
        email,
        password
    }, function(err, response) {
        if (err) {
            console.log(err.message);
            return callback(err.message);
        }
        return callback(response.user, response.token);
    });
}