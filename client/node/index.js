global.XMLHttpRequest = require('xhr2');
const authenticate = require("./authenticate")
const signIn = require("./signIn.js")
const signUp = require("./signUp.js")
const connect = require("./connect.js")

module.exports = {
    authenticate,
    signIn,
    signUp,
    connect
}

