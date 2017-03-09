var https = require("https");

module.exports = function(context, callback) {
    var req = https.request({
        host: 'hooks.slack.com',
        port: '443',
        path: '/services/T4C8JHY1F/B4FLEQ76G/LL7I2QoG8OytoBnLnP8Y6qZF',
        method: 'POST',
        headers: { 'Content-Type': 'application/json' }
    }, function(res) {
        console.log("done");
    });

    req.write('{"text": "Hi from QCon"}');
    req.end();
    callback(200, "ok\n");
}
