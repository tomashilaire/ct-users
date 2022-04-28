module.exports = function authenticate(client, callback) {
    client.Authenticate({}, function(err, response) {
        if (err) {
            console.log(err.message);
            return callback(err.message);
        }
        return callback(response.authId);
    });
}