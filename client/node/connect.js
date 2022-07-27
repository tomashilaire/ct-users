const {AuthenticationClient} = require('./users_grpc_web_pb.js');

module.exports = function connect(url) {
    return new AuthenticationClient(url);
}

