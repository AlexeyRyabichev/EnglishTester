const login = require('./jsPrivate/login.mjs');

module.exports = {
    checkAccess: login.checkAccess(user, password)
};